package handler

import (
	"context"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
)

type GrpcHandler struct {
	q *db.Queries
	kafka.UnimplementedEventRepositoryServer
}

// CreateEvent implements kafka.EventRepositoryServer.
func (g *GrpcHandler) CreateEvent(ctx context.Context, req *kafka.CreateEventReq) (*kafka.CreateEventResp, error) {
	panic("unimplemented")
}

// DeadEvent implements kafka.EventRepositoryServer.
func (g *GrpcHandler) DeadEvent(ctx context.Context, req *kafka.DeadEventReq) (*kafka.DeadEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusDEAD})
	if err != nil {
		return nil, err
	}
	return &kafka.DeadEventResp{}, nil
}

// DeliveredEvent implements kafka.EventRepositoryServer.
func (g *GrpcHandler) DeliveredEvent(ctx context.Context, req *kafka.DeliveredEventReq) (*kafka.DeliveredEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusDELIVERED})
	if err != nil {
		return nil, err
	}
	return &kafka.DeliveredEventResp{}, nil
}

// ReadEvent implements kafka.EventRepositoryServer.
func (g *GrpcHandler) ReadEvent(ctx context.Context, req *kafka.ReadEventReq) (*kafka.ReadEventResp, error) {
	panic("unimplemented")
}

// RetryingEvent implements kafka.EventRepositoryServer.
func (g *GrpcHandler) RetryingEvent(ctx context.Context, req *kafka.RetryingEventReq) (*kafka.RetryingEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusRETRYING})
	if err != nil {
		return nil, err
	}
	return &kafka.RetryingEventResp{}, nil
}

func NewGrpcHandler(q *db.Queries) *GrpcHandler {
	return &GrpcHandler{q: q}
}
