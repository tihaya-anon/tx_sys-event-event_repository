package mapping_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func TestPB2BytesAndBack(t *testing.T) {
	// Create a test event
	testEvent := &pb.Event{
		EventId:       "test-event-123",
		EventTopic:    "test.topic",
		EventType:     "TestEvent",
		Source:        "test-source",
		CreatedAt:     time.Now().Unix(),
		Status:        pb.Event_PENDING,
		Payload:       []byte(`{"test": "data"}`),
		DedupKey:      "dedup-key-123",
		CorrelationId: "correlation-123",
	}

	// Test PB2Bytes
	data, err := mapping.PB2Bytes(testEvent)
	require.NoError(t, err, "PB2Bytes should not return an error")
	assert.NotEmpty(t, data, "PB2Bytes should return non-empty data")

	// Test Bytes2PB
	recoveredEvent, err := mapping.Bytes2PB(data)
	require.NoError(t, err, "Bytes2PB should not return an error")
	assert.Equal(t, testEvent.EventId, recoveredEvent.EventId, "Event ID should match")
	assert.Equal(t, testEvent.EventTopic, recoveredEvent.EventTopic, "Event topic should match")
	assert.Equal(t, testEvent.EventType, recoveredEvent.EventType, "Event type should match")
	assert.Equal(t, testEvent.Source, recoveredEvent.Source, "Source should match")
	assert.Equal(t, testEvent.Status, recoveredEvent.Status, "Status should match")
	assert.Equal(t, testEvent.Payload, recoveredEvent.Payload, "Payload should match")
	assert.Equal(t, testEvent.DedupKey, recoveredEvent.DedupKey, "Dedup key should match")
	assert.Equal(t, testEvent.CorrelationId, recoveredEvent.CorrelationId, "Correlation ID should match")
}

func TestBytes2PB_InvalidData(t *testing.T) {
	// Test with invalid protobuf data
	invalidData := []byte("not a valid protobuf")
	event, err := mapping.Bytes2PB(invalidData)
	assert.Error(t, err, "Bytes2PB should return an error for invalid data")
	assert.Nil(t, event, "Event should be nil when error occurs")
}

func TestPB2Bytes_NilEvent(t *testing.T) {
	// Test with nil event
	data, err := mapping.PB2Bytes(nil)
	assert.Error(t, err, "PB2Bytes should return an error for nil event")
	assert.Nil(t, data, "Data should be nil when error occurs")
}
