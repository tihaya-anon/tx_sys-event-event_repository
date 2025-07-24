package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type FetchResp struct {
	Val string
	Err error
	TTL time.Duration
}

func Set(ctx context.Context, rdb *redis.Client, key string, value any, ttl time.Duration) error {
	var err error
	err = rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}
	err = rdb.Set(ctx, key+":default", value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, rdb *redis.Client, key string, fetch func(context.Context) FetchResp) (string, error) {
	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return getDefaultAndFetch(ctx, rdb, key, fetch)
	}
	return value, nil
}

func getDefaultAndFetch(ctx context.Context, rdb *redis.Client, key string, fetch func(context.Context) FetchResp) (string, error) {
	value, err := rdb.Get(ctx, key+":default").Result()
	wg := sync.WaitGroup{}
	wg.Add(1)

	var resp FetchResp
	go func() {
		resp = fetch(ctx)
		Set(ctx, rdb, key, resp.Val, resp.TTL)
		wg.Done()
	}()

	if err == redis.Nil {
		wg.Wait()
		return resp.Val, fmt.Errorf("no default value found for key %s", key)
	}
	if err != nil {
		return "", fmt.Errorf("error %s for key %s", err.Error(), key)
	}
	return value, nil
}
