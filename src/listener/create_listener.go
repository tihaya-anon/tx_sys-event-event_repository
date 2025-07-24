package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"

	constant_kafka "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/kafka"
	constant_redis "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/dao"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/mapping"
	wrapper "github.com/tihaya-anon/tx_sys-event-event_repository/src/redis"
)

type ConsumerInfo struct {
	GroupId  string
	Name     string
	MaxBytes int
}

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

// TODO unimplemented
func defaultFetch(ctx context.Context) wrapper.FetchResp {
	return wrapper.FetchResp{}
}

func getConsumerInfoByTopic(ctx context.Context, rdb *redis.Client, topic string) (*ConsumerInfo, error) {
	groupIdKey, nameKey, maxBytesKey := constant_redis.GetConsumerInfoKey(topic)
	wg := sync.WaitGroup{}
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()
	wg.Add(3)
	errChan := make(chan error, 1)
	var (
		groupId     string
		name        string
		maxBytesInt int
	)
	go func() {
		defer wg.Done()
		select {
		case <-ctxCancel.Done():
			return
		default:
		}
		groupId_, err := wrapper.Get(ctx, rdb, groupIdKey, defaultFetch)
		if err != nil {
			select {
			case errChan <- fmt.Errorf("get group id failed for topic %s", topic):
				cancel()
			default:
			}
			return
		}
		groupId = groupId_
	}()
	go func() {
		defer wg.Done()
		select {
		case <-ctxCancel.Done():
			return
		default:
		}
		name_, err := wrapper.Get(ctx, rdb, nameKey, defaultFetch)
		if err != nil {
			select {
			case errChan <- fmt.Errorf("get name failed for topic %s", topic):
				cancel()
			default:
			}
			return
		}
		name = name_
	}()
	go func() {
		defer wg.Done()
		select {
		case <-ctxCancel.Done():
			return
		default:
		}
		maxBytes, err := wrapper.Get(ctx, rdb, maxBytesKey, defaultFetch)
		if err != nil {
			select {
			case errChan <- fmt.Errorf("get max bytes failed for topic %s", topic):
				cancel()
			default:
			}
			return
		}
		maxBytesInt, err = strconv.Atoi(maxBytes)
		if err != nil {
			select {
			case errChan <- fmt.Errorf("convert max bytes failed for topic %s", topic):
				cancel()
			default:
			}
			return
		}
	}()
	wg.Wait()
	select {
	case err := <-errChan:
		return nil, err
	default:
		return &ConsumerInfo{
			GroupId:  groupId,
			Name:     name,
			MaxBytes: maxBytesInt,
		}, nil
	}
}
