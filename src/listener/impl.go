package listener

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

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
func (k *kafkaBridgeClientImpl) ListSubscriptions(ctx context.Context, groupID string, consumerName string) (any, *http.Response, error) {
	return k.client.ConsumersAPI.ListSubscriptions(ctx, groupID, consumerName).Execute()
}

// DeleteConsumer deletes a consumer
func (k *kafkaBridgeClientImpl) DeleteConsumer(ctx context.Context, groupID string, consumerName string) (*http.Response, error) {
	return k.client.ConsumersAPI.DeleteConsumer(ctx, groupID, consumerName).Execute()
}

// Poll polls for messages
func (k *kafkaBridgeClientImpl) Poll(ctx context.Context, groupID string, consumerName string, maxBytes int) ([]map[string]any, error) {
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
func (r *redisClientImpl) Set(ctx context.Context, key string, value string, expiration any) error {
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

// INTERFACE
var _ KafkaBridgeClient = (*kafkaBridgeClientImpl)(nil)
var _ RedisClient = (*redisClientImpl)(nil)
