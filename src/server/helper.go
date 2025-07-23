package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func filterWithEventId(f *pb.Query_Filter) bool {
	return f.Field == "event_id" && f.Op == pb.Query_Filter_EQ && len(f.Values) == 1
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
			pbEvent, err := mapping.DB2PB(&dbEvent)
			if err != nil {
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
