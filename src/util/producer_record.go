package util

import (
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func BuildProducerRecord(e *pb.Event, k *string) *kafka_bridge.ProducerRecord {
	if e == nil {
		return nil
	}
	payload := &e.Payload
	timestamp := e.CreatedAt
	kafkaHeaders := make([]kafka_bridge.KafkaHeader, len(e.Metadata))
	i := 0
	for k, v := range e.Metadata {
		kafkaHeaders[i] = kafka_bridge.KafkaHeader{
			Key:   k,
			Value: v,
		}
		i++
	}
	return &kafka_bridge.ProducerRecord{
		Timestamp: &timestamp,
		Key:       &kafka_bridge.RecordKey{String: k},
		Value:     *kafka_bridge.NewNullableRecordValue(&kafka_bridge.RecordValue{String: payload}),
		Headers:   kafkaHeaders,
	}
}
