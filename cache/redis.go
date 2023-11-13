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
	ctx    context.Context
	ttl    time.Duration
}

func (s *RedisStore) Get(key string) ([]byte, error) {
	item := s.client.Get(s.ctx, key)
	if item.Err() == redis.Nil {
		return nil, nil
	}
	return item.Bytes()
}

func (s *RedisStore) Set(key string, value []byte, ttl time.Duration) error {
	return s.client.Set(s.ctx, key, value, s.ttl).Err()
}

func (s *RedisStore) Delete(key string) error {
	return s.client.Del(s.ctx, key).Err()
}

func (s *RedisStore) Clear() error {
	return s.client.FlushAll(s.ctx).Err()
}
