package listener

// KafkaConsumerInfo represents the information needed to interact with a Kafka consumer
type KafkaConsumerInfo struct {
	GroupId  string
	Name     string
	MaxBytes int
	BaseURI  string
	PodName  string
}
