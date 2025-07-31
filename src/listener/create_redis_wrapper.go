package listener

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	constant_redis "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
	wrapper "github.com/tihaya-anon/tx_sys-event-event_repository/src/redis"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

// KafkaConsumerInfo is defined in types.go

// TODO unimplemented
func defaultFetch(ctx context.Context) wrapper.FetchResp {
	return wrapper.FetchResp{}
}

func getConsumerInfoByTopic(ctx context.Context, rdb *redis.Client, topic string) (*KafkaConsumerInfo, error) {
	groupIdKey, nameKey, maxBytesKey := constant_redis.GetConsumerInfoKey(topic)
	var (
		groupId     string
		name        string
		maxBytesInt int
	)
	concurrency := util.NewConcurrency(ctx)
	concurrency.Add(func(ctx_ context.Context) error {
		groupId_, err := wrapper.Get(ctx_, rdb, groupIdKey, defaultFetch)
		if err != nil {
			return err
		}
		groupId = groupId_
		return nil
	})
	concurrency.Add(func(ctx_ context.Context) error {
		name_, err := wrapper.Get(ctx_, rdb, nameKey, defaultFetch)
		if err != nil {
			return err
		}
		name = name_
		return nil
	})
	concurrency.Add(func(ctx_ context.Context) error {
		maxBytes, err := wrapper.Get(ctx_, rdb, maxBytesKey, defaultFetch)
		if err != nil {
			return err
		}
		maxBytesInt_, err := strconv.Atoi(maxBytes)
		if err != nil {
			return err
		}
		maxBytesInt = maxBytesInt_
		return nil
	})

	concurrency.Run()
	concurrency.Wait()

	if err := concurrency.Err(); err != nil {
		return nil, err
	}
	// Try to get additional metadata (may not exist for older consumers)
	baseUriKey := fmt.Sprintf("%s:baseuri", topic)
	podNameKey := fmt.Sprintf("%s:podname", topic)
	
	baseURI, _ := wrapper.Get(ctx, rdb, baseUriKey, defaultFetch)
	podName, _ := wrapper.Get(ctx, rdb, podNameKey, defaultFetch)
	
	return &KafkaConsumerInfo{
		GroupId:  groupId,
		Name:     name,
		MaxBytes: maxBytesInt,
		BaseURI:  baseURI,
		PodName:  podName,
	}, nil
}
