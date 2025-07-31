package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/listener"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/listener/mock"
)

func TestConsumerManager_InitializeConsumer_ExistingConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockKafkaClient := mock.NewMockKafkaBridgeClient(ctrl)
	mockRedisClient := mock.NewMockRedisClient(ctrl)

	// Create test data
	testTopic := "test-topic"
	testConsumerInfo := &listener.KafkaConsumerInfo{
		GroupId:  testTopic,
		Name:     "test-consumer",
		MaxBytes: 1024,
		BaseURI:  "http://test-uri",
		PodName:  "test-pod",
	}

	// Create consumer manager with mocks
	cm := listener.NewConsumerManagerWithInterfaces(mockKafkaClient, mockRedisClient)

	// Set up expectations for existing consumer case
	// First, the manager will check its local cache (which is empty)
	// Then it will try to get from Redis
	consumerInfoJSON, _ := json.Marshal(testConsumerInfo)
	mockRedisClient.EXPECT().Get(gomock.Any(), fmt.Sprintf("consumer:%s", testTopic)).Return(string(consumerInfoJSON), nil)

	// The consumer manager will store the consumer info back in Redis
	mockRedisClient.EXPECT().Set(gomock.Any(), fmt.Sprintf("consumer:%s", testTopic), gomock.Any(), gomock.Any()).Return(nil)

	// Mock the ListSubscriptions call to verify the consumer exists
	mockResp := &http.Response{StatusCode: 200}
	mockKafkaClient.EXPECT().ListSubscriptions(gomock.Any(), testConsumerInfo.GroupId, testConsumerInfo.Name).Return(nil, mockResp, nil)

	// Test the InitializeConsumer method
	result, err := cm.InitializeConsumer(context.Background(), testTopic)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testConsumerInfo.Name, result.Name)
	assert.Equal(t, testConsumerInfo.BaseURI, result.BaseURI)
}

func TestConsumerManager_InitializeConsumer_CreateNewConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockKafkaClient := mock.NewMockKafkaBridgeClient(ctrl)
	mockRedisClient := mock.NewMockRedisClient(ctrl)

	// Create test data
	testTopic := "test-topic"

	// Create consumer manager with mocks
	cm := listener.NewConsumerManagerWithInterfaces(mockKafkaClient, mockRedisClient)

	// Set up expectations for creating a new consumer
	// First, the manager will check its local cache (which is empty)
	// Then it will try to get from Redis, but fail
	mockRedisClient.EXPECT().Get(gomock.Any(), gomock.Any()).Return("", errors.New("not found")).AnyTimes()

	// Mock the CreateConsumer call
	testInstanceId := "test-consumer"
	testBaseUri := "http://test-uri"
	createdConsumer := &kafka_bridge.CreatedConsumer{
		InstanceId: &testInstanceId,
		BaseUri:    &testBaseUri,
	}
	mockResp := &http.Response{StatusCode: 200}
	mockKafkaClient.EXPECT().CreateConsumer(gomock.Any(), gomock.Any(), gomock.Any()).Return(createdConsumer, mockResp, nil)

	// Mock the Subscribe call
	mockSubscribeResp := &http.Response{StatusCode: 204}
	mockKafkaClient.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockSubscribeResp, nil)

	// Mock Redis set calls for storing consumer info
	mockRedisClient.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	// Test the InitializeConsumer method
	result, err := cm.InitializeConsumer(context.Background(), testTopic)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, *(createdConsumer.BaseUri), result.BaseURI)
}

func TestConsumerManager_Shutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockKafkaClient := mock.NewMockKafkaBridgeClient(ctrl)
	mockRedisClient := mock.NewMockRedisClient(ctrl)

	// Create test data
	testTopic := "test-topic"
	testConsumerInfo := &listener.KafkaConsumerInfo{
		GroupId:  testTopic,
		Name:     "test-consumer",
		MaxBytes: 1024,
		BaseURI:  "http://test-uri",
		PodName:  "test-pod",
	}

	// Create consumer manager with mocks
	cm := listener.NewConsumerManagerWithInterfaces(mockKafkaClient, mockRedisClient)

	// Add a consumer to the manager's cache
	cm.AddConsumerForTesting(testTopic, testConsumerInfo)

	// Mock the DeleteConsumer call
	mockResp := &http.Response{StatusCode: 204}
	mockKafkaClient.EXPECT().DeleteConsumer(gomock.Any(), testConsumerInfo.GroupId, testConsumerInfo.Name).Return(mockResp, nil)

	// Test the Shutdown method
	cm.Shutdown(context.Background())
}
