package server

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/rs/zerolog/log"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

func filterHasEventId(f *pb.Query_Filter) bool {
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

func (s *EventRepositoryServer) readEventByEventId(ctx context.Context, eventId string) (*pb.Event, error) {
	dbEvent, err := s.q.ReadEventByEventId(ctx, eventId)
	if err != nil {
		log.Error().Err(err).Msg("failed to read event by event id")
		return nil, err
	}
	pbEvent, err := mapping.DB2PB(&dbEvent)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert db event to pb event")
		return nil, err
	}
	return pbEvent, nil
}

func (s *EventRepositoryServer) updateEventStatus(ctx context.Context, eventId string, status db.DeliveryStatus) *pb.EventIdWrapper {
	err := s.q.UpdateEventStatus(ctx, db.UpdateEventStatusParams{EventID: eventId, Status: status})
	if err != nil {
		log.Error().Err(err).Msg("failed to update event status")
		return &pb.EventIdWrapper{EventId: eventId, Success: false, Error: err.Error()}
	}
	return &pb.EventIdWrapper{EventId: eventId, Success: true}
}

func (s *EventRepositoryServer) createEvent(ctx context.Context, event *pb.Event) *pb.EventIdWrapper {
	//TODO determine node id
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Error().Err(err).Msg("failed to create node")
		return &pb.EventIdWrapper{Success: false, Error: err.Error()}
	}
	event.EventId = node.Generate().String()
	dbEvent, err := mapping.PB2DB(event)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert pb to db")
		return &pb.EventIdWrapper{Success: false, Error: err.Error()}
	}
	err = s.q.CreateEvent(ctx, db.CreateEventParams(*dbEvent))
	if err != nil {
		log.Error().Err(err).Msg("failed to create event")
		return &pb.EventIdWrapper{Success: false, Error: err.Error()}
	}
	return &pb.EventIdWrapper{EventId: event.EventId, Success: true}
}
