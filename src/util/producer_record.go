package util

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func BuildProducerRecord(e *pb.Event) map[string]any {
	payload := e.Payload
	headers := make([]map[string]string, len(e.Metadata))
	timestamp := e.CreatedAt
	for k, v := range e.Metadata {
		headers = append(headers, map[string]string{
			"key":   k,
			"value": v,
		})
	}
	return map[string]any{
		"timestamp": timestamp,
		"key":       e.EventId,
		"value":     payload,
		"headers":   headers,
	}
}
