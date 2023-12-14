package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, ttl time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
}

type RedisStore struct {
	client RedisClient
	ttl    time.Duration
}

func (s *RedisStore) Get(ctx context.Context, key string) ([]byte, error) {
	item, err := s.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return item, err
}

func (s *RedisStore) Set(ctx context.Context, key string, value []byte) error {
	return s.client.Set(ctx, key, value, s.ttl).Err()
}

func (s *RedisStore) Delete(ctx context.Context, key string) error {
	return s.client.Del(ctx, key).Err()
}

func (s *RedisStore) Clear(ctx context.Context) error {
	return s.client.FlushAll(ctx).Err()
}
