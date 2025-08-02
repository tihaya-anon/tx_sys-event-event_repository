package util_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

// mockReader is a mock implementation of the db.Queries interface
type mockReader struct {
	readEventByDedupKeyFunc func(ctx context.Context, dedupKey pgtype.Text) (db.Event, error)
}

// ReadEventByDedupKey implements the ReadEventByDedupKey method of db.Queries
func (m *mockReader) ReadEventByDedupKey(ctx context.Context, dedupKey pgtype.Text) (db.Event, error) {
	if m.readEventByDedupKeyFunc != nil {
		return m.readEventByDedupKeyFunc(ctx, dedupKey)
	}
	return db.Event{}, nil
}

func TestIsDup(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func() *mockReader
		dedupKey       string
		expectedEvent  *db.Event
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name: "success - not a duplicate",
			setupMock: func() *mockReader {
				return &mockReader{
					readEventByDedupKeyFunc: func(ctx context.Context, dedupKey pgtype.Text) (db.Event, error) {
						return db.Event{}, pgx.ErrNoRows
					},
				}
			},
			dedupKey:      "unique-key-123",
			expectedEvent: nil,
			expectedError: false,
		},
		{
			name: "success - duplicate found",
			setupMock: func() *mockReader {
				return &mockReader{
					readEventByDedupKeyFunc: func(ctx context.Context, dedupKey pgtype.Text) (db.Event, error) {
						return db.Event{
							EventID:   "550e8400-e29b-41d4-a716-446655440000",
							DedupKey:  dedupKey,
							EventType: "test-event",
						}, nil
					},
				}
			},
			dedupKey: "duplicate-key-123",
			expectedEvent: &db.Event{
				EventID:   "550e8400-e29b-41d4-a716-446655440000",
				DedupKey:  pgtype.Text{String: "duplicate-key-123", Valid: true},
				EventType: "test-event",
			},
			expectedError: false,
		},
		{
			name: "error - database error",
			setupMock: func() *mockReader {
				return &mockReader{
					readEventByDedupKeyFunc: func(ctx context.Context, dedupKey pgtype.Text) (db.Event, error) {
						return db.Event{}, errors.New("database connection error")
					},
				}
			},
			dedupKey:       "error-key",
			expectedEvent:  nil,
			expectedError:  true,
			expectedErrMsg: "database connection error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q *mockReader
			if tt.setupMock != nil {
				q = tt.setupMock()
			}

			event, _, err := util.IsDup(context.Background(), q, tt.dedupKey)

			if tt.expectedError {
				require.Error(t, err)
				if tt.expectedErrMsg != "" {
					assert.Contains(t, err.Error(), tt.expectedErrMsg)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedEvent, event)
			}
		})
	}
}
