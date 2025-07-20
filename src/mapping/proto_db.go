package mapping

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func DB2PB(e *db.Event) *pb.Event {
	metadata, err := Bytes2Map(e.Metadata)
	if err != nil {
		return nil
	}
	statusIdx := pb.Event_DeliveryStatus_value[string(e.Status)]
	status := pb.Event_DeliveryStatus(statusIdx)
	return &pb.Event{
		EventId:       e.EventID,
		EventTopic:    e.EventTopic,
		EventType:     e.EventType,
		Source:        e.Source,
		CreatedAt:     e.CreatedAt,
		ExpiresAt:     e.ExpiresAt.Int64,
		Status:        status,
		RetryCount:    e.RetryCount.Int32,
		DedupKey:      e.DedupKey.String,
		Metadata:      metadata,
		Payload:       e.Payload,
		TargetService: e.TargetService.String,
		CorrelationId: e.CorrelationID.String,
	}
}
