package listener

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"

	constant_kafka "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
)

// Global consumer manager instance
var consumerManager ConsumerManagerInterface

// SetConsumerManagerForTesting sets the consumer manager for testing
func SetConsumerManagerForTesting(cm ConsumerManagerInterface) {
	consumerManager = cm
}

// ResetConsumerManagerForTesting resets the consumer manager after testing
func ResetConsumerManagerForTesting() {
	consumerManager = nil
}

// InitConsumerManager initializes the consumer manager with the Kafka Bridge client
func InitConsumerManager(ctx context.Context, rdb *redis.Client) {
	// Only initialize once
	if consumerManager != nil {
		return
	}

	// Create Kafka Bridge API client
	config := kafka_bridge.NewConfiguration()
	config.Servers = []kafka_bridge.ServerConfiguration{
		{URL: constant_kafka.KAFKA_BRIDGE_HOST},
	}
	client := kafka_bridge.NewAPIClient(config)

	// Create consumer manager
	consumerManager = NewConsumerManager(client, rdb)
	log.Println("Initialized Kafka consumer manager")
}

// ShutdownConsumerManager cleans up all consumers
func ShutdownConsumerManager(ctx context.Context) {
	if consumerManager != nil {
		consumerManager.Shutdown(ctx)
		log.Println("Shutdown Kafka consumer manager")
		consumerManager = nil
	}
}

// CreateListener fetches and processes messages from Kafka topics
func CreateListener(ctx context.Context, q dao.Query, rdb *redis.Client) {
	// Initialize consumer manager if needed
	if consumerManager == nil && rdb != nil {
		InitConsumerManager(ctx, rdb)
	}

	// Get or create consumer for the topic
	consumerInfo, err := consumerManager.InitializeConsumer(ctx, constant_kafka.KAFKA_BRIDGE_CREATE_TOPIC)
	if err != nil {
		log.Printf("Failed to initialize consumer: %v", err)
		return
	}

	// Get the Kafka Bridge client from the consumer manager
	client := consumerManager.GetKafkaBridgeClient()

	// Poll for messages
	messages, err := PollMessages(client, consumerInfo)
	if err != nil {
		log.Printf("Error polling messages: %v", err)
		return
	}

	// Process messages asynchronously
	for _, record := range messages {
		go SaveRecord(ctx, q, record)
	}
}

// PollMessages fetches messages from Kafka Bridge API using the KafkaBridgeClient interface
func PollMessages(client KafkaBridgeClient, consumerInfo *KafkaConsumerInfo) ([]map[string]any, error) {
	// Use the Poll method from the KafkaBridgeClient interface
	ctx := context.Background()
	messages, err := client.Poll(ctx, consumerInfo.GroupId, consumerInfo.Name, consumerInfo.MaxBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to poll for messages: %w", err)
	}

	return messages, nil
}

// SaveRecord processes and saves a Kafka record
func SaveRecord(ctx context.Context, q dao.Query, record map[string]any) {
	if record["topic"] != constant_kafka.KAFKA_BRIDGE_CREATE_TOPIC {
		return
	}
	payload := record["value"].([]byte)
	pbEvent, err := mapping.Bytes2PB(payload)
	if err != nil {
		log.Println(err)
		return
	}
	dbEvent, err := mapping.PB2DB(pbEvent)
	if err != nil {
		log.Println(err)
		return
	}
	err = q.CreateEvent(ctx, db.CreateEventParams(*dbEvent))
	if err != nil {
		log.Println(err)
	}
}
