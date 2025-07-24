package redis

import "fmt"

const (
	consumerInfoPrefix = "consumer_info:"
	groupIdKey         = "group_id:%s"
	nameKey            = "name:%s"
	maxBytesKey        = "max_bytes:%s"
)

func GetConsumerInfoKey(topic string) (string, string, string) {
	return fmt.Sprintf(consumerInfoPrefix+groupIdKey, topic), fmt.Sprintf(consumerInfoPrefix+nameKey, topic), fmt.Sprintf(consumerInfoPrefix+maxBytesKey, topic)
}
