// Cache package to provide an interface to interact with cache stores.
package cache

import (
	"context"
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/redis/go-redis/v9"
)

type StoreType string

const (
	BigCache   StoreType = "BigCache"
	RedisCache StoreType = "Redis"
)

type Cache struct {
	store     Store
	StoreType StoreType
	TTL       time.Duration
}

type Store interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error
}

var (
	ErrCacheIsDisabled = errors.New("cache is disabled")
	ErrRedisOptionsNil = errors.New("redis options is nil")
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
		return nil, ErrRedisOptionsNil
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
		},
		TTL:       ttl,
		StoreType: RedisCache,
	}
	return cache, nil
}

// Returns an item from the cache. If no item is found, returns nil for the item and error.
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	if c.TTL == 0 {
		return nil, ErrCacheIsDisabled
	}
	return c.store.Get(ctx, key)
}

// Saves an item under the key provided.
func (c *Cache) Set(ctx context.Context, key string, item []byte) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}
	return c.store.Set(ctx, key, item)
}

// Deletes an item from the cache.
func (c *Cache) Delete(ctx context.Context, key string) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}
	return c.store.Delete(ctx, key)
}

// Clears the entire cache.
func (c *Cache) Clear(ctx context.Context) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}
	return c.store.Clear(ctx)
}
