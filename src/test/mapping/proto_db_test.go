package mapping_test

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	. "github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
)

func createTestPBEvent() *pb.Event {
	return &pb.Event{
		EventId:       "test-event-123",
		EventTopic:    "test.topic",
		EventType:     "TestEvent",
		Source:        "test-source",
		CreatedAt:     time.Now().Unix(),
		ExpiresAt:     time.Now().Add(24 * time.Hour).Unix(),
		Status:        pb.Event_PENDING,
		RetryCount:    0,
		DedupKey:      "dedup-key-123",
		Metadata:      map[string]string{"key1": "value1", "key2": "value2"},
		Payload:       `{"test": "data"}`,
		TargetService: "test-service",
		CorrelationId: "correlation-123",
	}
}

func createTestDBEvent() *db.Event {
	metadata, _ := Map2Bytes(map[string]string{"key1": "value1", "key2": "value2"})
	return &db.Event{
		EventID:       "test-event-123",
		EventTopic:    "test.topic",
		EventType:     "TestEvent",
		Source:        "test-source",
		CreatedAt:     time.Now().Unix(),
		ExpiresAt:     pgtype.Int8{Int64: time.Now().Add(24 * time.Hour).Unix(), Valid: true},
		Status:        db.DeliveryStatusPENDING,
		RetryCount:    pgtype.Int4{Int32: 0, Valid: true},
		DedupKey:      pgtype.Text{String: "dedup-key-123", Valid: true},
		Metadata:      metadata,
		Payload:       pgtype.Text{String: `{"test": "data"}`, Valid: true},
		TargetService: pgtype.Text{String: "test-service", Valid: true},
		CorrelationID: pgtype.Text{String: "correlation-123", Valid: true},
	}
}

func TestDB2PB(t *testing.T) {
	// Setup
	dbEvent := createTestDBEvent()

	// Execute
	pbEvent, err := DB2PB(dbEvent)

	// Verify
	require.NoError(t, err, "DB2PB should not return an error")
	assert.Equal(t, dbEvent.EventID, pbEvent.EventId)
	assert.Equal(t, dbEvent.EventTopic, pbEvent.EventTopic)
	assert.Equal(t, dbEvent.EventType, pbEvent.EventType)
	assert.Equal(t, dbEvent.Source, pbEvent.Source)
	assert.Equal(t, dbEvent.CreatedAt, pbEvent.CreatedAt)
	assert.Equal(t, dbEvent.ExpiresAt.Int64, pbEvent.ExpiresAt)
	assert.Equal(t, string(dbEvent.Status), pbEvent.Status.String())
	assert.Equal(t, dbEvent.RetryCount.Int32, pbEvent.RetryCount)
	assert.Equal(t, dbEvent.DedupKey.String, pbEvent.DedupKey)
	assert.Equal(t, dbEvent.TargetService.String, pbEvent.TargetService)
	assert.Equal(t, dbEvent.CorrelationID.String, pbEvent.CorrelationId)

	// Test metadata
	metadata, err := Bytes2Map(dbEvent.Metadata)
	require.NoError(t, err, "Bytes2Map should not return an error")
	assert.Equal(t, metadata, pbEvent.Metadata)
}

func TestPB2DB(t *testing.T) {
	// Setup
	pbEvent := createTestPBEvent()

	// Execute
	dbEvent, err := PB2DB(pbEvent)

	// Verify
	require.NoError(t, err, "PB2DB should not return an error")
	assert.Equal(t, pbEvent.EventId, dbEvent.EventID)
	assert.Equal(t, pbEvent.EventTopic, dbEvent.EventTopic)
	assert.Equal(t, pbEvent.EventType, dbEvent.EventType)
	assert.Equal(t, pbEvent.Source, dbEvent.Source)
	assert.Equal(t, pbEvent.CreatedAt, dbEvent.CreatedAt)
	assert.Equal(t, pbEvent.ExpiresAt, dbEvent.ExpiresAt.Int64)
	assert.Equal(t, pbEvent.Status.String(), string(dbEvent.Status))
	assert.Equal(t, pbEvent.RetryCount, dbEvent.RetryCount.Int32)
	assert.Equal(t, pbEvent.DedupKey, dbEvent.DedupKey.String)
	assert.Equal(t, pbEvent.TargetService, dbEvent.TargetService.String)
	assert.Equal(t, pbEvent.CorrelationId, dbEvent.CorrelationID.String)

	// Test metadata
	metadata, err := Map2Bytes(pbEvent.Metadata)
	require.NoError(t, err, "Map2Bytes should not return an error")
	assert.Equal(t, metadata, dbEvent.Metadata)
}

func TestDB2PB_NilEvent(t *testing.T) {
	// Test with nil event
	pbEvent, err := DB2PB(nil)
	assert.Error(t, err, "DB2PB should return an error for nil event")
	assert.Nil(t, pbEvent, "PB event should be nil when error occurs")
}

func TestPB2DB_NilEvent(t *testing.T) {
	// Test with nil event
	dbEvent, err := PB2DB(nil)
	assert.Error(t, err, "PB2DB should return an error for nil event")
	assert.Nil(t, dbEvent, "DB event should be nil when error occurs")
}

func TestDB2PB_InvalidMetadata(t *testing.T) {
	// Create a DB event with invalid metadata
	dbEvent := createTestDBEvent()
	dbEvent.Metadata = []byte("invalid json")

	// Execute and verify
	pbEvent, err := DB2PB(dbEvent)
	assert.Error(t, err, "DB2PB should return an error for invalid metadata")
	t.Logf("Expected error: %v", err)
	assert.Nil(t, pbEvent, "PB event should be nil when error occurs")
}
