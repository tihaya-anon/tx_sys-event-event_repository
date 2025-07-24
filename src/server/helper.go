package server

import (
	"context"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

func filterWithEventId(f *pb.Query_Filter) bool {
	return f.Field == "event_id" && f.Op == pb.Query_Filter_EQ && len(f.Values) == 1
}

func paraMapping(dbEvents []db.Event) ([]*pb.Event, error) {
	pbEvents := make([]*pb.Event, len(dbEvents))
	concurrency := util.NewConcurrency(context.Background())
	for i, dbEvent := range dbEvents {
		concurrency.Add(func(ctx context.Context) error {
			pbEvent, err := mapping.DB2PB(&dbEvent)
			if err != nil {
				return err
			}
			pbEvents[i] = pbEvent
			return nil
		})
	}
	concurrency.Run()
	concurrency.Wait()
	if err := concurrency.Err(); err != nil {
		return nil, err
	}
	return pbEvents, nil
}

func (g *EventRepositoryServer) readEventByEventId(ctx context.Context, eventId string) (*pb.Event, error) {
	dbEvent, err := g.q.ReadEventByEventId(ctx, eventId)
	if err != nil {
		return nil, err
	}
	pbEvent, err := mapping.DB2PB(&dbEvent)
	if err != nil {
		return nil, err
	}
	return pbEvent, nil
}
