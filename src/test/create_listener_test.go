package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/listener"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/listener/mock"
)

// TestPollMessages tests the pollMessages function
func TestPollMessages(t *testing.T) {
	// Create test consumer info
	testConsumerInfo := &listener.KafkaConsumerInfo{
		GroupId:  "test-topic",
		Name:     "test-consumer",
		MaxBytes: 1024,
		BaseURI:  "http://test-uri",
		PodName:  "test-pod",
	}

	// Test successful polling
	t.Run("Successful polling", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		// Create mock client
		mockKafkaClient := mock.NewMockKafkaBridgeClient(ctrl)

		// Setup mock response
		expectedMessages := []map[string]any{
			{
				"topic": "test-topic",
				"key":   "test-key",
				"value": "test-value",
			},
		}

		// Configure mock to return expected messages
		mockKafkaClient.EXPECT().
			Poll(gomock.Any(), testConsumerInfo.GroupId, testConsumerInfo.Name, testConsumerInfo.MaxBytes).
			Return(expectedMessages, nil)

		// Call the function
		messages, err := listener.PollMessages(mockKafkaClient, testConsumerInfo)

		// Assertions
		assert.NoError(t, err)
		assert.Len(t, messages, 1)
		assert.Equal(t, "test-topic", messages[0]["topic"])
		assert.Equal(t, "test-key", messages[0]["key"])
		assert.Equal(t, "test-value", messages[0]["value"])
	})

	// Test error case
	t.Run("Poll error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		// Create mock client
		mockKafkaClient := mock.NewMockKafkaBridgeClient(ctrl)

		// Configure mock to return an error
		mockKafkaClient.EXPECT().
			Poll(gomock.Any(), testConsumerInfo.GroupId, testConsumerInfo.Name, testConsumerInfo.MaxBytes).
			Return(nil, errors.New("connection error"))

		// Call the function
		messages, err := listener.PollMessages(mockKafkaClient, testConsumerInfo)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, messages)
		assert.Contains(t, err.Error(), "connection error")
	})
}

// Skip TestCreateListener for now as it requires more mocking setup
// We'll focus on the PollMessages test which is more straightforward
