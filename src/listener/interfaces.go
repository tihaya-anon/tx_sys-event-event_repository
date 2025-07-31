package listener

import (
	"context"
	"net/http"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

// KafkaBridgeClient defines the interface for Kafka Bridge API operations
//
//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock
type KafkaBridgeClient interface {
	CreateConsumer(ctx context.Context, groupID string, consumer kafka_bridge.Consumer) (*kafka_bridge.CreatedConsumer, *http.Response, error)
	Subscribe(ctx context.Context, groupID string, consumerName string, topics kafka_bridge.Topics) (*http.Response, error)
	ListSubscriptions(ctx context.Context, groupID string, consumerName string) (interface{}, *http.Response, error)
	DeleteConsumer(ctx context.Context, groupID string, consumerName string) (*http.Response, error)
	Poll(ctx context.Context, groupID string, consumerName string, maxBytes int) ([]map[string]interface{}, error)
}

// RedisClient defines the interface for Redis operations
//
//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock
type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration interface{}) error
	Ping(ctx context.Context) error
	Close() error
}

// ConsumerManagerInterface defines the interface for consumer management operations
//
//go:generate mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock
type ConsumerManagerInterface interface {
	InitializeConsumer(ctx context.Context, topic string) (*KafkaConsumerInfo, error)
	CleanupConsumers(ctx context.Context)
	GetConsumerInfo(topic string) (*KafkaConsumerInfo, bool)
	Shutdown(ctx context.Context)
}
