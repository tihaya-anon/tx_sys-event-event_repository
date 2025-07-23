package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
)

func CreateListener(ctx context.Context, q *db.Queries) {
	groupId, name, maxBytes := getComsumerInfoByTopic(constant.KAFKA_BRIDGE_CREATE_TOPIC)
	consumerURL := fmt.Sprintf(
		"%s/consumers/%s/instances/%s/records?max_bytes=%d",
		constant.KAFKA_BRIDGE_HOST, groupId, name, maxBytes,
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

func saveRecord(ctx context.Context, q *db.Queries, record map[string]any) {
	if record["topic"] != constant.KAFKA_BRIDGE_CREATE_TOPIC {
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

//TODO get suggestion from cache
func getComsumerInfoByTopic(topic string) (string, string, int) {
	groupId := fmt.Sprintf("<test-group-id-%s>", topic)
	name := fmt.Sprintf("<test-name-%s>", topic)
	maxBytes := 1024
	return groupId, name, maxBytes
}
