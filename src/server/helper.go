package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func filterHasEventId(f *pb.Query_Filter) bool {
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

func (s *EventRepositoryServer) readEventByEventId(ctx context.Context, eventId string) (*pb.Event, error) {
	dbEvent, err := s.q.ReadEventByEventId(ctx, eventId)
	if err != nil {
		return nil, err
	}
	pbEvent, err := mapping.DB2PB(&dbEvent)
	if err != nil {
		return nil, err
	}
	return pbEvent, nil
}

func (s *EventRepositoryServer) updateEventStatus(ctx context.Context, eventId string, status db.DeliveryStatus) *pb.EventIdWrapper {
	err := s.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: eventId, Status: status})
	if err != nil {
		return &pb.EventIdWrapper{EventId: eventId, Success: false, Error: err.Error()}
	}
	return &pb.EventIdWrapper{EventId: eventId, Success: true}
}

func (s *EventRepositoryServer) createEvent(ctx context.Context, event *pb.Event) *pb.EventIdWrapper {
	//TODO determine node id
	node, err := snowflake.NewNode(1)
	if err != nil {
		return &pb.EventIdWrapper{Success: false, Error: err.Error()}
	}
	event.EventId = node.Generate().String()
	dbEvent, err := mapping.PB2DB(event)
	if err != nil {
		return &pb.EventIdWrapper{Success: false, Error: err.Error()}
	}
	err = s.q.CreateEvent(ctx, db.CreateEventParams(*dbEvent))
	if err != nil {
		return &pb.EventIdWrapper{Success: false, Error: err.Error()}
	}
	return &pb.EventIdWrapper{EventId: event.EventId, Success: true}
}
