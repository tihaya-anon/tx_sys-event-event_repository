package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"

	constant_kafka "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
)


func CreateListener(ctx context.Context, q dao.Query, rdb *redis.Client) {
	consumerInfo, err := getConsumerInfoByTopic(ctx, rdb, constant_kafka.KAFKA_BRIDGE_CREATE_TOPIC)
	if err != nil {
		return
	}
	consumerURL := fmt.Sprintf(
		"%s/consumers/%s/instances/%s/records?max_bytes=%d",
		constant_kafka.KAFKA_BRIDGE_HOST, consumerInfo.GroupId, consumerInfo.Name, consumerInfo.MaxBytes,
	)
	resp, err := http.Get(consumerURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	var body []map[string]any
	_ = json.Unmarshal(bytes, &body)
	for _, record := range body {
		go saveRecord(ctx, q, record)
	}
}

func saveRecord(ctx context.Context, q dao.Query, record map[string]any) {
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
