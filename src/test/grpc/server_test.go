package grpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao/mock"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	server "github.com/tihaya-anon/tx_sys-event-event_repository/src/server"
)

func TestEventRepositoryServer_CreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuery := mock.NewMockQuery(ctrl)
	mockReader := mock.NewMockReader(ctrl)
	srv := server.NewGrpcHandler(mockQuery, mockReader)

	tests := []struct {
		name          string
		setupMocks    func()
		req           *kafka.CreateEventReq
		wantEventID   string
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful event creation",
			setupMocks: func() {
				mockQuery.EXPECT().
					ReadEventByDedupKey(gomock.Any(), gomock.Any()).
					Return(db.Event{}, nil)
			},
			req: &kafka.CreateEventReq{
				Event: &pb.Event{
					DedupKey: "test-key",
					Payload:  []byte(`{"test":"data"}`),
				},
			},
			wantEventID: "", // Will be set in test
			wantErr:     false,
		},
		{
			name: "duplicate event",
			setupMocks: func() {
				event := db.Event{
					EventID:   "existing-id",
					DedupKey:  pgtype.Text{String: "duplicate-key", Valid: true},
					Status:    db.DeliveryStatusPENDING,
					CreatedAt: time.Now().Unix(),
				}
				mockQuery.EXPECT().
					ReadEventByDedupKey(gomock.Any(), gomock.Any()).
					Return(event, nil)
			},
			req: &kafka.CreateEventReq{
				Event: &pb.Event{
					DedupKey: "duplicate-key",
					Payload:  []byte(`{"test":"data"}`),
				},
			},
			wantEventID: "existing-id",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			resp, err := srv.CreateEvent(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				require.NoError(t, err)
				if tt.wantEventID != "" {
					assert.Equal(t, tt.wantEventID, resp.EventId)
				} else {
					// Verify we got a valid UUID
					_, err := uuid.Parse(resp.EventId)
					assert.NoError(t, err, "Expected a valid UUID")
				}
			}
		})
	}
}

func TestEventRepositoryServer_DeadEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuery := mock.NewMockQuery(ctrl)
	mockReader := mock.NewMockReader(ctrl)
	srv := server.NewGrpcHandler(mockQuery, mockReader)

	tests := []struct {
		name          string
		setupMocks    func()
		req           *kafka.DeadEventReq
		wantErr       bool
		expectedError error
	}{
		{
			name: "successful dead event",
			setupMocks: func() {
				mockQuery.EXPECT().
					UpdateEventStatus(gomock.Any(), db.UpdateEventStatusParams{
						EventID: "test-id",
						Status:  db.DeliveryStatusDEAD,
					}).
					Return(nil)
			},
			req: &kafka.DeadEventReq{
				EventId: "test-id",
			},
			wantErr: false,
		},
		{
			name: "database error",
			setupMocks: func() {
				mockQuery.EXPECT().
					UpdateEventStatus(gomock.Any(), gomock.Any()).
					Return(assert.AnError)
			},
			req: &kafka.DeadEventReq{
				EventId: "test-id",
			},
			wantErr:       true,
			expectedError: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := srv.DeadEvent(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError.Error() != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEventRepositoryServer_DeliveredEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuery := mock.NewMockQuery(ctrl)
	mockReader := mock.NewMockReader(ctrl)
	srv := server.NewGrpcHandler(mockQuery, mockReader)

	tests := []struct {
		name          string
		setupMocks    func()
		req           *kafka.DeliveredEventReq
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful delivered event",
			setupMocks: func() {
				mockQuery.EXPECT().
					UpdateEventStatus(gomock.Any(), db.UpdateEventStatusParams{
						EventID: "test-id",
						Status:  db.DeliveryStatusDELIVERED,
					}).
					Return(nil)
			},
			req: &kafka.DeliveredEventReq{
				EventId: "test-id",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := srv.DeliveredEvent(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEventRepositoryServer_ReadEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuery := mock.NewMockQuery(ctrl)
	mockReader := mock.NewMockReader(ctrl)
	srv := server.NewGrpcHandler(mockQuery, mockReader)

	tests := []struct {
		name          string
		setupMocks    func()
		req           *kafka.ReadEventReq
		expectedEvent *pb.Event
		wantErr       bool
		expectedError string
	}{
		{
			name: "read event by ID",
			setupMocks: func() {
				dbEvent := db.Event{
					EventID:   "test-id",
					DedupKey:  pgtype.Text{String: "test-key", Valid: true},
					Status:    db.DeliveryStatusPENDING,
					Payload:   []byte(`{"test":"data"}`),
					CreatedAt: time.Now().Unix(),
				}
				mockQuery.EXPECT().
					ReadEventByEventId(gomock.Any(), "test-id").
					Return(dbEvent, nil)
			},
			req: &kafka.ReadEventReq{
				Query: &pb.Query{
					Filters: []*pb.Query_Filter{
						{
							Field: "event_id",
							Op:    pb.Query_Filter_EQ,
							Values: []string{
								"test-id",
							},
						},
					},
				},
			},
			expectedEvent: &pb.Event{
				EventId:  "test-id",
				DedupKey: "test-key",
				Payload:  []byte(`{"test":"data"}`),
			},
			wantErr: false,
		},
		{
			name: "read events with query",
			setupMocks: func() {
				dbEvents := []db.Event{
					{
						EventID:   "test-id-1",
						DedupKey:  pgtype.Text{String: "test-key-1", Valid: true},
						Status:    db.DeliveryStatusPENDING,
						Payload:   []byte(`{"test":"data1"}`),
						CreatedAt: time.Now().Unix(),
					},
				}
				mockReader.EXPECT().
					Select(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(dbEvents, nil)
				mockReader.EXPECT().
					Count(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(int64(1), nil)
			},
			req: &kafka.ReadEventReq{
				Query: &pb.Query{
					Filters: []*pb.Query_Filter{
						{
							Field: "status",
							Op:    pb.Query_Filter_EQ,
							Values: []string{
								"PENDING",
							},
						},
					},
				},
			},
			expectedEvent: &pb.Event{
				EventId:  "test-id-1",
				DedupKey: "test-key-1",
				Payload:  []byte(`{"test":"data1"}`),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			resp, err := srv.ReadEvent(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp.Result)
				if tt.expectedEvent != nil {
					require.Greater(t, len(resp.Result.Events), 0)
					assert.Equal(t, tt.expectedEvent.EventId, resp.Result.Events[0].EventId)
					assert.Equal(t, tt.expectedEvent.DedupKey, resp.Result.Events[0].DedupKey)
					assert.JSONEq(t, string(tt.expectedEvent.Payload), string(resp.Result.Events[0].Payload))
				}
			}
		})
	}
}

func TestEventRepositoryServer_RetryingEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuery := mock.NewMockQuery(ctrl)
	mockReader := mock.NewMockReader(ctrl)
	srv := server.NewGrpcHandler(mockQuery, mockReader)

	tests := []struct {
		name          string
		setupMocks    func()
		req           *kafka.RetryingEventReq
		wantErr       bool
		expectedError string
	}{
		{
			name: "successful retrying event",
			setupMocks: func() {
				mockQuery.EXPECT().
					UpdateEventStatus(gomock.Any(), db.UpdateEventStatusParams{
						EventID: "test-id",
						Status:  db.DeliveryStatusRETRYING,
					}).
					Return(nil)
			},
			req: &kafka.RetryingEventReq{
				EventId: "test-id",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := srv.RetryingEvent(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedError != "" {
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Helper function to create a test event
func testEvent(id, dedupKey string, payload []byte) *pb.Event {
	return &pb.Event{
		EventId:   id,
		DedupKey:  dedupKey,
		Payload:   payload,
		Status:    pb.Event_PENDING,
		CreatedAt: time.Now().Unix(),
	}
}

// Helper function to create a test filter
func testFilter(field string, op pb.Query_Filter_Operator, values ...string) *pb.Query_Filter {
	return &pb.Query_Filter{
		Field:  field,
		Op:     op,
		Values: values,
	}
}
