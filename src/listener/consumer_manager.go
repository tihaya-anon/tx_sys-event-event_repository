package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	constant_kafka "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

// ConsumerManager handles the lifecycle of Kafka consumers
type ConsumerManager struct {
	kafkaClient KafkaBridgeClient
	redisClient RedisClient
	consumers   map[string]*KafkaConsumerInfo
	mu          sync.RWMutex
}

// NewConsumerManager creates a new consumer manager
func NewConsumerManager(client *kafka_bridge.APIClient, rdb *redis.Client) *ConsumerManager {
	return &ConsumerManager{
		kafkaClient: &kafkaBridgeClientImpl{client: client},
		redisClient: &redisClientImpl{rdb: rdb},
		consumers:   make(map[string]*KafkaConsumerInfo),
	}
}

// NewConsumerManagerWithInterfaces creates a new consumer manager with custom interfaces
func NewConsumerManagerWithInterfaces(kafkaClient KafkaBridgeClient, redisClient RedisClient) *ConsumerManager {
	return &ConsumerManager{
		kafkaClient: kafkaClient,
		redisClient: redisClient,
		consumers:   make(map[string]*KafkaConsumerInfo),
	}
}

// InitializeConsumer creates a consumer for a topic if it doesn't exist
func (cm *ConsumerManager) InitializeConsumer(ctx context.Context, topic string) (*KafkaConsumerInfo, error) {
	// Check if we already have this consumer in memory
	cm.mu.RLock()
	if consumer, exists := cm.consumers[topic]; exists {
		cm.mu.RUnlock()
		return consumer, nil
	}
	cm.mu.RUnlock()

	// Try to get from Redis and validate it still exists
	// Use the interface methods instead of type assertion
	consumerInfoJSON, err := cm.redisClient.Get(ctx, fmt.Sprintf("consumer:%s", topic))
	var consumerInfo *KafkaConsumerInfo
	if err == nil && consumerInfoJSON != "" {
		consumerInfo = &KafkaConsumerInfo{}
		if err := json.Unmarshal([]byte(consumerInfoJSON), consumerInfo); err != nil {
			return nil, fmt.Errorf("failed to unmarshal consumer info: %w", err)
		}
	}
	if err == nil && consumerInfo != nil {
		// Verify the consumer still exists by checking subscriptions
		_, resp, err := cm.kafkaClient.ListSubscriptions(ctx, consumerInfo.GroupId, consumerInfo.Name)
		if err == nil && resp.StatusCode == 200 {
			// Store in local cache and Redis
			cm.mu.Lock()
			defer cm.mu.Unlock()
			cm.consumers[topic] = consumerInfo
			cm.storeConsumerInfo(ctx, topic, consumerInfo)
			return consumerInfo, nil
		}
	}

	// Create a new consumer
	return cm.CreateConsumer(ctx, topic)
}

// CreateConsumer creates a new consumer for a topic
func (cm *ConsumerManager) CreateConsumer(ctx context.Context, topic string) (*KafkaConsumerInfo, error) {
	// Validate topic name
	if topic == "" {
		topic = constant_kafka.KAFKA_BRIDGE_CREATE_TOPIC // Default topic if none provided
	}

	// Use topic as consumer group ID for simplicity
	groupID := topic

	// Generate a unique consumer name with pod name and timestamp
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "local"
	}
	timestamp := time.Now().Unix()
	consumerName := fmt.Sprintf("%s-%s-%d", topic, podName, timestamp)

	// Setup consumer config
	consumer := kafka_bridge.NewConsumer()
	consumer.SetName(consumerName)
	consumer.SetFormat("json")
	consumer.SetAutoOffsetReset("earliest")
	consumer.SetFetchMinBytes(1)
	consumer.SetConsumerRequestTimeoutMs(30000)
	consumer.SetEnableAutoCommit(true)

	// Create consumer via API
	createdConsumer, resp, err := cm.kafkaClient.CreateConsumer(ctx, groupID, *consumer)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Failed to create consumer: %v, status: %d", err, resp.StatusCode)
		return nil, fmt.Errorf("failed to create consumer: %v", err)
	}

	// Subscribe to topic
	topics := kafka_bridge.NewTopics()
	topics.SetTopics([]string{topic})

	resp, err = cm.kafkaClient.Subscribe(ctx, groupID, consumerName, *topics)
	if err != nil || resp.StatusCode != 204 {
		log.Printf("Failed to subscribe to topic: %v, status: %d", err, resp.StatusCode)
		// Clean up the consumer since subscription failed
		cm.kafkaClient.DeleteConsumer(ctx, groupID, consumerName)
		return nil, fmt.Errorf("failed to subscribe to topic: %v", err)
	}

	// Create consumer info
	consumerInfo := &KafkaConsumerInfo{
		GroupId:  groupID,
		Name:     consumerName,
		MaxBytes: 1048576, // 1MB
		BaseURI:  createdConsumer.GetBaseUri(),
		PodName:  podName,
	}

	// Store in Redis and local cache
	cm.storeConsumerInfo(ctx, topic, consumerInfo)

	cm.mu.Lock()
	cm.consumers[topic] = consumerInfo
	cm.mu.Unlock()

	log.Printf("Successfully created consumer %s for topic %s", consumerName, topic)
	return consumerInfo, nil
}

// storeConsumerInfo stores consumer info in Redis
func (cm *ConsumerManager) storeConsumerInfo(ctx context.Context, topic string, consumerInfo *KafkaConsumerInfo) error {
	// Convert consumer info to JSON
	consumerInfoJSON, err := json.Marshal(consumerInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal consumer info: %w", err)
	}

	// Store in Redis using the interface
	return cm.redisClient.Set(ctx, fmt.Sprintf("consumer:%s", topic), string(consumerInfoJSON), 24*time.Hour)
}

// CleanupConsumers deletes all consumers created by this manager
func (cm *ConsumerManager) CleanupConsumers(ctx context.Context) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for topic, info := range cm.consumers {
		_, err := cm.kafkaClient.DeleteConsumer(ctx, info.GroupId, info.Name)
		if err != nil {
			log.Printf("Failed to delete consumer for topic %s: %v", topic, err)
		} else {
			log.Printf("Successfully deleted consumer for topic %s", topic)
		}
	}

	cm.consumers = make(map[string]*KafkaConsumerInfo)
}

// GetConsumerInfo returns consumer info for a topic
func (cm *ConsumerManager) GetConsumerInfo(topic string) (*KafkaConsumerInfo, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	info, exists := cm.consumers[topic]
	return info, exists
}

// Shutdown gracefully shuts down the consumer manager
func (cm *ConsumerManager) Shutdown(ctx context.Context) {
	cm.CleanupConsumers(ctx)
}

// GetKafkaBridgeClient returns the Kafka Bridge client used by this consumer manager
func (cm *ConsumerManager) GetKafkaBridgeClient() KafkaBridgeClient {
	return cm.kafkaClient
}

// AddConsumerForTesting adds a consumer to the manager's cache for testing purposes
func (cm *ConsumerManager) AddConsumerForTesting(topic string, info *KafkaConsumerInfo) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.consumers[topic] = info
}
