package listener

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	constant_kafka "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	constant_redis "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

// kafkaBridgeClientImpl implements KafkaBridgeClient using the Kafka Bridge API client
type kafkaBridgeClientImpl struct {
	client *kafka_bridge.APIClient
}

// CreateConsumer creates a new Kafka consumer
func (k *kafkaBridgeClientImpl) CreateConsumer(ctx context.Context, groupID string, consumer kafka_bridge.Consumer) (*kafka_bridge.CreatedConsumer, *http.Response, error) {
	return k.client.ConsumersAPI.CreateConsumer(ctx, groupID).Consumer(consumer).Execute()
}

// Subscribe subscribes a consumer to topics
func (k *kafkaBridgeClientImpl) Subscribe(ctx context.Context, groupID string, consumerName string, topics kafka_bridge.Topics) (*http.Response, error) {
	resp, err := k.client.ConsumersAPI.Subscribe(ctx, groupID, consumerName).Topics(topics).Execute()
	return resp, err
}

// ListSubscriptions lists a consumer's subscriptions
func (k *kafkaBridgeClientImpl) ListSubscriptions(ctx context.Context, groupID string, consumerName string) (interface{}, *http.Response, error) {
	return k.client.ConsumersAPI.ListSubscriptions(ctx, groupID, consumerName).Execute()
}

// DeleteConsumer deletes a consumer
func (k *kafkaBridgeClientImpl) DeleteConsumer(ctx context.Context, groupID string, consumerName string) (*http.Response, error) {
	return k.client.ConsumersAPI.DeleteConsumer(ctx, groupID, consumerName).Execute()
}

// Poll polls for messages
func (k *kafkaBridgeClientImpl) Poll(ctx context.Context, groupID string, consumerName string, maxBytes int) ([]map[string]interface{}, error) {
	// This is implemented in pollMessages function, but we include it here for interface completeness
	return nil, fmt.Errorf("not implemented directly in client wrapper")
}

// redisClientImpl implements RedisClient using go-redis
type redisClientImpl struct {
	rdb *redis.Client
}

// Get gets a value from Redis
func (r *redisClientImpl) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

// Set sets a value in Redis
func (r *redisClientImpl) Set(ctx context.Context, key string, value string, expiration interface{}) error {
	return r.rdb.Set(ctx, key, value, expiration.(time.Duration)).Err()
}

// Ping pings Redis
func (r *redisClientImpl) Ping(ctx context.Context) error {
	return r.rdb.Ping(ctx).Err()
}

// Close closes the Redis connection
func (r *redisClientImpl) Close() error {
	return r.rdb.Close()
}

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
	// Get the Redis client implementation from our interface
	redisImpl, ok := cm.redisClient.(*redisClientImpl)
	if !ok {
		return nil, fmt.Errorf("invalid Redis client implementation")
	}
	consumerInfo, err := getConsumerInfoByTopic(ctx, redisImpl.rdb, topic)
	if err == nil && consumerInfo != nil {
		// Verify the consumer still exists by checking subscriptions
		_, resp, err := cm.kafkaClient.ListSubscriptions(ctx, consumerInfo.GroupId, consumerInfo.Name)
		if err == nil && resp.StatusCode == 200 {
			// Store in local cache
			cm.mu.Lock()
			cm.consumers[topic] = consumerInfo
			cm.mu.Unlock()
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
	cm.storeConsumerInfoInRedis(ctx, topic, consumerInfo)
	
	cm.mu.Lock()
	cm.consumers[topic] = consumerInfo
	cm.mu.Unlock()
	
	log.Printf("Successfully created consumer %s for topic %s", consumerName, topic)
	return consumerInfo, nil
}

// storeConsumerInfoInRedis stores consumer information in Redis
func (cm *ConsumerManager) storeConsumerInfoInRedis(ctx context.Context, topic string, info *KafkaConsumerInfo) {
	// Store required consumer info
	groupIdKey, nameKey, maxBytesKey := constant_redis.GetConsumerInfoKey(topic)
	expiration := 24 * time.Hour
	
	cm.redisClient.Set(ctx, groupIdKey, info.GroupId, expiration)
	cm.redisClient.Set(ctx, nameKey, info.Name, expiration)
	cm.redisClient.Set(ctx, maxBytesKey, fmt.Sprintf("%d", info.MaxBytes), expiration)
	
	// Store additional metadata for observability
	baseUriKey := fmt.Sprintf("%s:baseuri", topic)
	podNameKey := fmt.Sprintf("%s:podname", topic)
	
	cm.redisClient.Set(ctx, baseUriKey, info.BaseURI, expiration)
	cm.redisClient.Set(ctx, podNameKey, info.PodName, expiration)
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
