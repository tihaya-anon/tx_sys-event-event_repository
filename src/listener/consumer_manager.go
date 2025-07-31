package listener

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
	constant_redis "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
	wrapper "github.com/tihaya-anon/tx_sys-event-event_repository/src/redis"
)

// ConsumerManager handles the lifecycle of Kafka consumers
type ConsumerManager struct {
	client    *kafka_bridge.APIClient
	rdb       *redis.Client
	consumers map[string]*KafkaConsumerInfo
	mu        sync.RWMutex
}

// NewConsumerManager creates a new consumer manager
func NewConsumerManager(client *kafka_bridge.APIClient, rdb *redis.Client) *ConsumerManager {
	return &ConsumerManager{
		client:    client,
		rdb:       rdb,
		consumers: make(map[string]*KafkaConsumerInfo),
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
	consumerInfo, err := getConsumerInfoByTopic(ctx, cm.rdb, topic)
	if err == nil && consumerInfo != nil {
		// Verify the consumer still exists by checking subscriptions
		_, resp, err := cm.client.ConsumersAPI.ListSubscriptions(ctx, consumerInfo.GroupId, consumerInfo.Name).Execute()
		if err == nil && resp.StatusCode == 200 {
			// Store in local cache
			cm.mu.Lock()
			cm.consumers[topic] = consumerInfo
			cm.mu.Unlock()
			return consumerInfo, nil
		}
	}

	// Create a new consumer
	return cm.createNewConsumer(ctx, topic)
}

// createNewConsumer creates a new Kafka consumer and subscribes it to the topic
func (cm *ConsumerManager) createNewConsumer(ctx context.Context, topic string) (*KafkaConsumerInfo, error) {
	// Generate a unique consumer name using pod name if available
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "local-pod"
	}
	
	// Create consumer configuration
	consumerName := fmt.Sprintf("%s-%s-%d", topic, podName, time.Now().UnixNano())
	groupID := fmt.Sprintf("group-%s", topic)
	
	// Setup consumer config
	consumer := kafka_bridge.NewConsumer()
	consumer.SetName(consumerName)
	consumer.SetFormat("json")
	consumer.SetAutoOffsetReset("earliest")
	consumer.SetFetchMinBytes(1)
	consumer.SetConsumerRequestTimeoutMs(30000)
	consumer.SetEnableAutoCommit(true)
	
	// Create consumer via API
	createdConsumer, resp, err := cm.client.ConsumersAPI.CreateConsumer(ctx, groupID).Consumer(*consumer).Execute()
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Failed to create consumer: %v, status: %d", err, resp.StatusCode)
		return nil, fmt.Errorf("failed to create consumer: %v", err)
	}
	
	// Subscribe to topic
	topics := kafka_bridge.NewTopics()
	topics.SetTopics([]string{topic})
	
	resp, err = cm.client.ConsumersAPI.Subscribe(ctx, groupID, consumerName).Topics(*topics).Execute()
	if err != nil || resp.StatusCode != 204 {
		log.Printf("Failed to subscribe to topic: %v, status: %d", err, resp.StatusCode)
		// Clean up the consumer since subscription failed
		cm.client.ConsumersAPI.DeleteConsumer(ctx, groupID, consumerName).Execute()
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
	
	wrapper.Set(ctx, cm.rdb, groupIdKey, info.GroupId, expiration)
	wrapper.Set(ctx, cm.rdb, nameKey, info.Name, expiration)
	wrapper.Set(ctx, cm.rdb, maxBytesKey, fmt.Sprintf("%d", info.MaxBytes), expiration)
	
	// Store additional metadata for observability
	baseUriKey := fmt.Sprintf("%s:baseuri", topic)
	podNameKey := fmt.Sprintf("%s:podname", topic)
	
	wrapper.Set(ctx, cm.rdb, baseUriKey, info.BaseURI, expiration)
	wrapper.Set(ctx, cm.rdb, podNameKey, info.PodName, expiration)
}

// CleanupConsumers deletes all consumers created by this manager
func (cm *ConsumerManager) CleanupConsumers(ctx context.Context) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	for topic, info := range cm.consumers {
		_, err := cm.client.ConsumersAPI.DeleteConsumer(ctx, info.GroupId, info.Name).Execute()
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
