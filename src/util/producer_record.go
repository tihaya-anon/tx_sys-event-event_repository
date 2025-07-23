package util

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func BuildProducerRecord(e *pb.Event) map[string]any {
	if e == nil {
		return nil
	}
	payload := e.Payload
	timestamp := e.CreatedAt
	headers := mapping.Metadata2Headers(e.Metadata)
	return map[string]any{
		"timestamp": timestamp,
		"key":       e.EventId,
		"value":     payload,
		"headers":   headers,
	}
}
