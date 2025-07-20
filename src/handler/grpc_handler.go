package handler

import (
	"context"
	"fmt"
	"sync"

	sq "github.com/Masterminds/squirrel"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

type GrpcHandler struct {
	q  *db.Queries
	tx db.DBTX
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
	// Use PostgreSQL-style placeholders ($1, $2, ...)
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	dv := util.DefaultValues{OrderBy: "create_at DESC", Limit: 100}
	pageQuery, err := util.BuildQueryFromProto(psql, "events", dv, req.Query)
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
func (g *GrpcHandler) RetryingEvent(ctx context.Context, req *kafka.RetryingEventReq) (*kafka.RetryingEventResp, error) {
	err := g.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: req.EventId, Status: db.DeliveryStatusRETRYING})
	if err != nil {
		return nil, err
	}
	return &kafka.RetryingEventResp{}, nil
}

func paraMapping(dbEvents []db.Event) ([]*pb.Event, error) {
	pbEvents := make([]*pb.Event, len(dbEvents))
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 1)

	for i, dbEvent := range dbEvents {
		wg.Add(1)
		go func(i int, dbEvent db.Event) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
			}
			pbEvent := mapping.DB2PB(&dbEvent)
			if pbEvent == nil {
				select {
				case errChan <- fmt.Errorf("mapping failed for event index %d", i):
					cancel() // First error wins
				default:
				}
				return
			}
			pbEvents[i] = pbEvent
		}(i, dbEvent)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return nil, err
	default:
		return pbEvents, nil
	}
}

func NewGrpcHandler(q *db.Queries, tx db.DBTX) *GrpcHandler {
	return &GrpcHandler{q: q, tx: tx}
}
