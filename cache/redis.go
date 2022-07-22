package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
	Del(ctx context.Context, key ...string) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
}

type RedisStore struct {
	client RedisClient
	ctx    context.Context
	ttl    time.Duration
}

func (s *RedisStore) Get(key string) ([]byte, error) {
	item, err := s.client.Get(s.ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	return []byte(item), nil
}

func (s *RedisStore) Set(key string, value []byte, ttl time.Duration) error {
	_, err := s.client.Set(s.ctx, key, value, s.ttl).Result()

	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStore) Delete(key string) error {
	_, err := s.client.Del(s.ctx, key).Result()

	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStore) Clear() error {
	_, err := s.client.FlushAll(s.ctx).Result()

	return err
}
