package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/redis/go-redis/v9"
)

type CacheStoreType string

const (
	BigCache   CacheStoreType = "BigCache"
	RedisCache CacheStoreType = "Redis"
)

type Cache struct {
	store     Store
	TTL       time.Duration
	StoreType CacheStoreType
}

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
	Clear() error
}

var (
	ErrCacheIsDisabled = errors.New("Cache is disabled")
)

// Creates a new Cache using BigCache.
//
// Requires a BigCache config that can be created with bigcache.DefaultConfig(n*time.Minute).
func NewBigCache(ctx context.Context, config bigcache.Config) (*Cache, error) {
	bigcache, err := bigcache.New(ctx, config)
	if err != nil {
		return nil, err
	}

	cache := &Cache{
		store:     &BigCacheStore{client: bigcache},
		TTL:       config.LifeWindow,
		StoreType: BigCache,
	}

	return cache, nil
}

// Creates a new Cache using go-redis.
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
		TTL:       ttl,
		StoreType: RedisCache,
	}

	return cache, nil
}

// Returns an item from the cache. If no item is found, returns nil for the item and error.
func (c *Cache) Get(key string) ([]byte, error) {
	if c.TTL == 0 {
		return nil, ErrCacheIsDisabled
	}

	return c.store.Get(key)
}

// Saves an item under the key provided.
func (c *Cache) Set(key string, item []byte) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}

	return c.store.Set(key, item, c.TTL)
}

// Deletes an item from the cache.
func (c *Cache) Delete(key string) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}

	return c.store.Delete(key)
}

// Clears the entire cache.
func (c *Cache) Clear() error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}

	return c.store.Clear()
}
