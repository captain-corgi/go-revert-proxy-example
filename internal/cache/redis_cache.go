package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client  *redis.Client
	ttl     time.Duration
}

func NewRedisCache(addr string, ttl time.Duration) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{client: client, ttl: ttl}
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, bool) {
	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, false
	}
	return []byte(result), err == nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value []byte) error {
	return r.client.Set(ctx, key, value, r.ttl).Err()
}
