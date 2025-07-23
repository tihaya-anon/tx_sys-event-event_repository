package util_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	. "github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

func TestBuildProducerRecord(t *testing.T) {
	// Setup
	timestamp := time.Now().Unix()
	testEvent := &pb.Event{
		EventId:   "test-event-123",
		CreatedAt: timestamp,
		Metadata:  map[string]string{"key1": "value1", "key2": "value2"},
		Payload:   []byte(`{"test": "data"}`),
	}

	// Execute
	record := BuildProducerRecord(testEvent)

	// Verify
	expectedHeaders := []map[string]string{
		{"key": "key1", "value": "value1"},
		{"key": "key2", "value": "value2"},
	}

	assert.Equal(t, testEvent.EventId, record["key"], "Event ID should match key")
	assert.Equal(t, testEvent.Payload, record["value"], "Payload should match value")
	assert.Equal(t, timestamp, record["timestamp"], "Timestamp should match")

	// Verify headers (order doesn't matter)
	recordHeaders := record["headers"].([]map[string]string)
	assert.Len(t, recordHeaders, 2, "Should have 2 headers")
	t.Logf("RecordHeaders: %v", recordHeaders)
	headerMap := make(map[string]string, 0)
	for _, h := range recordHeaders {
		t.Logf("Header: %s=%s", h["key"], h["value"])
		headerMap[h["key"]] = h["value"]
		t.Logf("HeaderMap: %v", headerMap)
	}
	for _, h := range expectedHeaders {
		assert.Equal(t, h["value"], headerMap[h["key"]], "Header %s should match", h["key"])
	}
}

func TestBuildProducerRecord_EmptyMetadata(t *testing.T) {
	// Setup
	timestamp := time.Now().Unix()
	testEvent := &pb.Event{
		EventId:   "test-event-123",
		CreatedAt: timestamp,
		Payload:   []byte(`{"test": "data"}`),
	}

	// Execute
	record := BuildProducerRecord(testEvent)

	// Verify
	assert.Equal(t, testEvent.EventId, record["key"])
	assert.Equal(t, testEvent.Payload, record["value"])
	assert.Equal(t, timestamp, record["timestamp"])
	assert.Empty(t, record["headers"], "Headers should be empty for nil metadata")
}

func TestBuildProducerRecord_NilEvent(t *testing.T) {
	// Execute
	record := BuildProducerRecord(nil)

	// Verify
	nilMap := map[string]any(nil)
	assert.Equal(t, nilMap, record, "Should return nil map for nil event")
}
