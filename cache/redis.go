package cache

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client    *redis.Client
	namespace string
	ttl       time.Duration
}

func (s RedisStore) Get(ctx context.Context, key string) ([]byte, error) {
	keys := []string{s.namespace, key}
	newKey := strings.Join(keys, ":")
	item, err := s.client.Get(ctx, newKey).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return item, err
}

func (s RedisStore) Set(ctx context.Context, key string, value []byte) error {
	keys := []string{s.namespace, key}
	newKey := strings.Join(keys, ":")
	return s.client.Set(ctx, newKey, value, s.ttl).Err()
}

func (s RedisStore) Delete(ctx context.Context, key string) error {
	keys := []string{s.namespace, key}
	newKey := strings.Join(keys, ":")
	return s.client.Del(ctx, newKey).Err()
}

func (s RedisStore) Clear(ctx context.Context) error {
	keys := []string{s.namespace, "*"}
	cache := strings.Join(keys, ":")
	return s.client.Del(ctx, cache).Err()
}
