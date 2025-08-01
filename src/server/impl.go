package server

import (
	"context"
	"log"
	"sync"

	sq "github.com/Masterminds/squirrel"
	kafka_constant "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

type EventRepositoryServer struct {
	r dao.Reader
	q dao.Query
	c *kafka_bridge.APIClient
	kafka.UnimplementedEventRepositoryServer
}

// CreateEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) CreateEvent(ctx context.Context, req *kafka.CreateEventReq) (*kafka.CreateEventResp, error) {
	eventSize := len(req.Events)
	eventIdWrappers := make([]*pb.EventIdWrapper, eventSize)
	newEventIds := make(map[int]string, eventSize)
	wg := sync.WaitGroup{}
	for i, event := range req.Events {
		wg.Add(1)
		go func(i int, ctx_ context.Context, q dao.Query, dedupKey string) {
			defer wg.Done()
			dbEvent, result, err := util.IsDup(ctx_, q, dedupKey)
			switch result {
			case util.DEDUP_RESULT_NEW:
				//TODO generate uuid
				newEventIds[i] = "uuid-gen"
				eventIdWrappers[i] = &pb.EventIdWrapper{EventId: "uuid-gen", Success: true, Error: ""}
			case util.DEDUP_RESULT_DUP:
				eventIdWrappers[i] = &pb.EventIdWrapper{EventId: dbEvent.EventID, Success: true, Error: ""}
			case util.DEDUP_RESULT_ERROR:
				eventIdWrappers[i] = &pb.EventIdWrapper{EventId: "", Success: false, Error: err.Error()}
			}
		}(i, ctx, s.q, event.DedupKey)
	}
	wg.Wait()
	producerRecords := make([]kafka_bridge.ProducerRecord, 0)
	for i, newEventId := range newEventIds {
		wg.Add(1)
		go func(i int, newEventId string) {
			defer wg.Done()
			pr := util.BuildProducerRecord(req.Events[i], &newEventId)
			if pr == nil {
				return
			}
			producerRecords = append(producerRecords, *pr)
		}(i, newEventId)
	}
	wg.Wait()
	producerRecordList := kafka_bridge.NewProducerRecordList()
	producerRecordList.SetRecords(producerRecords)
	_, r, err := s.c.TopicsAPI.Send(ctx, kafka_constant.KAFKA_BRIDGE_CREATE_TOPIC).ProducerRecordList(*producerRecordList).Async(true).Execute()
	if err != nil {
		for i, newEventId := range newEventIds {
			log.Printf("ID: %s, Error: %v, Full Response: %v", newEventId, err, r)
			eventIdWrappers[i] = &pb.EventIdWrapper{EventId: newEventId, Success: false, Error: err.Error()}
		}
	}
	return &kafka.CreateEventResp{EventIds: eventIdWrappers}, nil
}

// DeadEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) DeadEvent(ctx context.Context, req *kafka.DeadEventReq) (*kafka.DeadEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.EventIds))
	wg := sync.WaitGroup{}
	for _, eventId := range req.EventIds {
		wg.Add(1)
		go func(eventId string) {
			defer wg.Done()
			eventIds = append(eventIds, s.updateEventStatus(ctx, eventId, db.DeliveryStatusDEAD))
		}(eventId)
	}
	wg.Wait()
	return &kafka.DeadEventResp{EventIds: eventIds}, nil
}

// DeliveredEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) DeliveredEvent(ctx context.Context, req *kafka.DeliveredEventReq) (*kafka.DeliveredEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.EventIds))
	wg := sync.WaitGroup{}
	for _, eventId := range req.EventIds {
		wg.Add(1)
		go func(eventId string) {
			defer wg.Done()
			eventIds = append(eventIds, s.updateEventStatus(ctx, eventId, db.DeliveryStatusDELIVERED))
		}(eventId)
	}
	wg.Wait()
	return &kafka.DeliveredEventResp{EventIds: eventIds}, nil
}

// ReadEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) ReadEvent(ctx context.Context, req *kafka.ReadEventReq) (*kafka.ReadEventResp, error) {
	for _, f := range req.Query.Filters {
		if filterWithEventId(f) {
			event, err := s.readEventByEventId(ctx, f.Values[0])
			if err != nil {
				return nil, err
			}
			return &kafka.ReadEventResp{Result: &pb.Result{Events: []*pb.Event{event}}}, nil
		}
	}
	// Use PostgreSQL-style placeholders ($1, $2, ...)
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	tableName := "events"
	dv := util.DefaultValues{OrderBy: "create_at DESC", Limit: 100}
	pageQuery, err := util.BuildQueryFromProto(psql, tableName, dv, req.Query)
	if err != nil {
		return nil, err
	}
	dbEvents, err := s.r.Select(ctx, pageQuery.PageSql, pageQuery.PageArgs...)
	if err != nil {
		return nil, err
	}
	totalSize, err := s.r.Count(ctx, pageQuery.TotalSql, pageQuery.TotalArgs...)
	if err != nil {
		return nil, err
	}
	pbEvents, err := paraMapping(dbEvents)
	if err != nil {
		return nil, err
	}
	return &kafka.ReadEventResp{Result: &pb.Result{Events: pbEvents, TotalSize: totalSize}}, nil
}

// RetryingEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) RetryingEvent(ctx context.Context, req *kafka.RetryingEventReq) (*kafka.RetryingEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.EventIds))
	wg := sync.WaitGroup{}
	for _, eventId := range req.EventIds {
		wg.Add(1)
		go func(eventId string) {
			defer wg.Done()
			eventIds = append(eventIds, s.updateEventStatus(ctx, eventId, db.DeliveryStatusRETRYING))
		}(eventId)
	}
	wg.Wait()
	return &kafka.RetryingEventResp{EventIds: eventIds}, nil
}

func (s *EventRepositoryServer) updateEventStatus(ctx context.Context, eventId string, status db.DeliveryStatus) *pb.EventIdWrapper {
	err := s.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: eventId, Status: status})
	if err != nil {
		return &pb.EventIdWrapper{EventId: eventId, Success: false, Error: err.Error()}
	}
	return &pb.EventIdWrapper{EventId: eventId, Success: true, Error: ""}
}

func NewGrpcHandler(q dao.Query, r dao.Reader, c *kafka_bridge.APIClient) *EventRepositoryServer {
	return &EventRepositoryServer{q: q, r: r, c: c}
}
