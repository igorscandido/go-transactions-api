package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/igorscandido/go-transactions-api/pkg/configs"
	"github.com/redis/go-redis/v9"
)

type redisAdapter struct {
	client *redis.Client
}

func NewRedisCache(configs *configs.Configs) (*redisAdapter, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configs.Redis.Addr, configs.Redis.Port),
		Password: configs.Redis.Password,
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}
	return &redisAdapter{client: client}, nil
}

func (r *redisAdapter) Get(ctx context.Context, key string) (interface{}, bool) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, false
	}

	var result interface{}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, false
	}
	return result, true
}

func (r *redisAdapter) Set(ctx context.Context, key string, value interface{}, ttlSeconds int) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return
	}

	err = r.client.Set(ctx, key, jsonValue, time.Duration(ttlSeconds)*time.Second).Err()
	if err != nil {
		return
	}
}

func (r *redisAdapter) Delete(ctx context.Context, key string) {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return
	}
}
