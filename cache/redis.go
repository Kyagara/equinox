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
}

type RedisStore struct {
	client    RedisClient
	namespace string
	ttl       time.Duration
}

func (s *RedisStore) Get(ctx context.Context, key string) ([]byte, error) {
	newKey := s.namespace + key
	item, err := s.client.Get(ctx, newKey).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return item, err
}

func (s *RedisStore) Set(ctx context.Context, key string, value []byte) error {
	newKey := s.namespace + key
	return s.client.Set(ctx, newKey, value, s.ttl).Err()
}

func (s *RedisStore) Delete(ctx context.Context, key string) error {
	newKey := s.namespace + key
	return s.client.Del(ctx, newKey).Err()
}

func (s *RedisStore) Clear(ctx context.Context) error {
	cache := s.namespace + "*"
	return s.client.Del(ctx, cache).Err()
}
