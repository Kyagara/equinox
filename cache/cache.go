package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v9"
)

type Cache struct {
	store Store
	TTL   time.Duration
}

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
	Clear() error
}

// Creates a new Cache using BigCache
// Requires a BigCache config that can be created with bigcache.DefaultConfig(n*time.Minute)
func NewBigCache(config bigcache.Config) (*Cache, error) {
	bigcache, err := bigcache.NewBigCache(config)

	if err != nil {
		return nil, err
	}

	cache := &Cache{
		store: &BigCacheStore{
			client: bigcache,
		},
		TTL: config.LifeWindow,
	}

	return cache, nil
}

// Creates a new Cache using go-redis
func NewRedis(ctx context.Context, options *redis.Options, ttl time.Duration) (*Cache, error) {
	if options == nil {
		return nil, fmt.Errorf("redis options is empty")
	}

	redis := redis.NewClient(options)

	err := redis.Ping(ctx).Err()

	if err != nil {
		return nil, err
	}

	cache := &Cache{
		store: &RedisStore{
			client: redis,
			ttl:    ttl,
			ctx:    ctx,
		},
		TTL: ttl,
	}

	return cache, nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	value, err := c.store.Get(key)

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *Cache) Set(key string, item []byte) error {
	return c.store.Set(key, item, c.TTL)
}

func (c *Cache) Delete(key string) error {
	return c.store.Delete(key)
}

func (c *Cache) Clear() error {
	return c.store.Clear()
}
