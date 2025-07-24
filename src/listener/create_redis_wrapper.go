package listener

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	constant_redis "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/redis"
	wrapper "github.com/tihaya-anon/tx_sys-event-event_repository/src/redis"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/util"
)

type ConsumerInfo struct {
	GroupId  string
	Name     string
	MaxBytes int
}

// TODO unimplemented
func defaultFetch(ctx context.Context) wrapper.FetchResp {
	return wrapper.FetchResp{}
}

func getConsumerInfoByTopic(ctx context.Context, rdb *redis.Client, topic string) (*ConsumerInfo, error) {
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
	return &ConsumerInfo{
		GroupId:  groupId,
		Name:     name,
		MaxBytes: maxBytesInt,
	}, nil
}
