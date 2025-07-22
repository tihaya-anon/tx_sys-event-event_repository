package mapping

import (
	"github.com/jackc/pgx/v5/pgtype"
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

func PB2DB(e *pb.Event) *db.Event {
	metadata, err := Map2Bytes(e.Metadata)
	if err != nil {
		return nil
	}
	var status db.DeliveryStatus
	err = status.Scan(e.Status)
	if err != nil {
		return nil
	}
	return &db.Event{
		EventID:       e.EventId,
		EventTopic:    e.EventTopic,
		EventType:     e.EventType,
		Source:        e.Source,
		CreatedAt:     e.CreatedAt,
		ExpiresAt:     pgtype.Int8{Int64: e.ExpiresAt, Valid: true},
		Status:        status,
		RetryCount:    pgtype.Int4{Int32: e.RetryCount, Valid: true},
		DedupKey:      pgtype.Text{String: e.DedupKey, Valid: true},
		Metadata:      metadata,
		Payload:       e.Payload,
		TargetService: pgtype.Text{String: e.TargetService, Valid: true},
		CorrelationID: pgtype.Text{String: e.CorrelationId, Valid: true},
	}
}
