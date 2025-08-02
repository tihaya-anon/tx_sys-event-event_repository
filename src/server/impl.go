package server

import (
	"context"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

type EventRepositoryServer struct {
	r dao.Reader
	q dao.Query
	kafka.UnimplementedEventRepositoryServer
}

// CreateEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) CreateEvent(ctx context.Context, req *kafka.CreateEventReq) (*kafka.CreateEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.Events))
	wg := sync.WaitGroup{}
	for i, event := range req.Events {
		wg.Add(1)
		go func(i int, event *pb.Event) {
			defer wg.Done()
			eventIds[i] = s.createEvent(ctx, event)
		}(i, event)
	}
	wg.Wait()
	return &kafka.CreateEventResp{EventIds: eventIds}, nil
}

// DeadEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) DeadEvent(ctx context.Context, req *kafka.DeadEventReq) (*kafka.DeadEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.EventIds))
	wg := sync.WaitGroup{}
	for i, eventId := range req.EventIds {
		wg.Add(1)
		go func(i int, eventId string) {
			defer wg.Done()
			eventIds[i] = s.updateEventStatus(ctx, eventId, db.DeliveryStatusDEAD)
		}(i, eventId)
	}
	wg.Wait()
	return &kafka.DeadEventResp{EventIds: eventIds}, nil
}

// DeliveredEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) DeliveredEvent(ctx context.Context, req *kafka.DeliveredEventReq) (*kafka.DeliveredEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.EventIds))
	wg := sync.WaitGroup{}
	for i, eventId := range req.EventIds {
		wg.Add(1)
		go func(i int, eventId string) {
			defer wg.Done()
			eventIds[i] = s.updateEventStatus(ctx, eventId, db.DeliveryStatusDELIVERED)
		}(i, eventId)
	}
	wg.Wait()
	return &kafka.DeliveredEventResp{EventIds: eventIds}, nil
}

// ReadEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) ReadEvent(ctx context.Context, req *kafka.ReadEventReq) (*kafka.ReadEventResp, error) {
	for _, f := range req.Query.Filters {
		if filterHasEventId(f) {
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
	dv := &util.DefaultValues{OrderBy: "create_at DESC", Limit: 100}
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
	nextPageToken := util.EncodePageToken(dv, req.Query)
	return &kafka.ReadEventResp{Result: &pb.Result{Events: pbEvents, TotalSize: totalSize, NextPageToken: nextPageToken}}, nil
}

// RetryingEvent implements kafka.EventRepositoryServer.
func (s *EventRepositoryServer) RetryingEvent(ctx context.Context, req *kafka.RetryingEventReq) (*kafka.RetryingEventResp, error) {
	eventIds := make([]*pb.EventIdWrapper, len(req.EventIds))
	wg := sync.WaitGroup{}
	for i, eventId := range req.EventIds {
		wg.Add(1)
		go func(i int, eventId string) {
			defer wg.Done()
			eventIds[i] = s.updateEventStatus(ctx, eventId, db.DeliveryStatusRETRYING)
		}(i, eventId)
	}
	wg.Wait()
	return &kafka.RetryingEventResp{EventIds: eventIds}, nil
}

func NewGrpcHandler(q dao.Query, r dao.Reader) *EventRepositoryServer {
	return &EventRepositoryServer{q: q, r: r}
}
