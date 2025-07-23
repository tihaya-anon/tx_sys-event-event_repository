package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/bwmarrin/snowflake"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

type EventRepositoryServer struct {
	q  *db.Queries
	tx db.DBTX
	kafka.UnimplementedEventRepositoryServer
}

// CreateEvent implements kafka.EventRepositoryServer.
func (g *EventRepositoryServer) CreateEvent(ctx context.Context, req *kafka.CreateEventReq) (*kafka.CreateEventResp, error) {
	event, err := util.IsDup(ctx, g.q, req.Event.DedupKey)
	if err != nil { // error occurred
		return nil, err
	}
	if event != nil { // duplicated
		return &kafka.CreateEventResp{EventId: event.EventID}, nil
	}
	//TODO determine the snowflake node id
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate()
	go func(url string, event *pb.Event) {
		pr := util.BuildProducerRecord(event)
		body, _ := json.Marshal(pr)
		http.Post(url, constant.KAFKA_BRIDGE_JSON, bytes.NewBuffer(body))
	}(fmt.Sprintf("%s/topics/%s", constant.KAFKA_BRIDGE_HOST, constant.KAFKA_BRIDGE_CREATE_TOPIC), req.Event)
	return &kafka.CreateEventResp{EventId: id.String()}, nil
}

// DeadEvent implements kafka.EventRepositoryServer.
func (g *EventRepositoryServer) DeadEvent(ctx context.Context, req *kafka.DeadEventReq) (*kafka.DeadEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusDEAD})
	if err != nil {
		return nil, err
	}
	return &kafka.DeadEventResp{}, nil
}

// DeliveredEvent implements kafka.EventRepositoryServer.
func (g *EventRepositoryServer) DeliveredEvent(ctx context.Context, req *kafka.DeliveredEventReq) (*kafka.DeliveredEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusDELIVERED})
	if err != nil {
		return nil, err
	}
	return &kafka.DeliveredEventResp{}, nil
}

// ReadEvent implements kafka.EventRepositoryServer.
func (g *EventRepositoryServer) ReadEvent(ctx context.Context, req *kafka.ReadEventReq) (*kafka.ReadEventResp, error) {
	for _, f := range req.Query.Filters {
		if filterWithEventId(f) {
			event, err := g.readEventByEventId(ctx, f.Values[0])
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
	dbEvents, err := dao.Read(ctx, g.tx, pageQuery.PageSql, pageQuery.PageArgs...)
	if err != nil {
		return nil, err
	}
	totalSize, err := dao.Count(ctx, g.tx, pageQuery.TotalSql, pageQuery.TotalArgs...)
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
func (g *EventRepositoryServer) RetryingEvent(ctx context.Context, req *kafka.RetryingEventReq) (*kafka.RetryingEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusRETRYING})
	if err != nil {
		return nil, err
	}
	return &kafka.RetryingEventResp{}, nil
}

func newGrpcHandler(q *db.Queries, tx db.DBTX) *EventRepositoryServer {
	return &EventRepositoryServer{q: q, tx: tx}
}
