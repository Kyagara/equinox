// Cache package to provide an interface to interact with cache stores.
package cache

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/redis/go-redis/v9"
)

var (
	ErrCacheIsDisabled = errors.New("cache is disabled")
	ErrRedisOptionsNil = errors.New("redis options is nil")
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
	// Returns an item from the cache. If no item is found, returns nil for the item and error.
	Get(ctx context.Context, key string) ([]byte, error)

	// Saves an item under the key provided.
	Set(ctx context.Context, key string, value []byte) error

	// Deletes an item from the cache.
	Delete(ctx context.Context, key string) error

	// Clears the entire cache.
	//
	// For Redis, the entire 'cache' namespace under 'equinox' is deleted.
	Clear(ctx context.Context) error
}

// Creates a new Cache using BigCache.
//
// Requires a BigCache config that can be created with bigcache.DefaultConfig(n*time.Minute).
func NewBigCache(ctx context.Context, config bigcache.Config) (*Cache, error) {
	bigcache, err := bigcache.New(ctx, config)
	if err != nil {
		return nil, err
	}
	cache := &Cache{
		store:     BigCacheStore{client: bigcache},
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
		store: RedisStore{
			client:    redis,
			ttl:       ttl,
			namespace: "equinox:cache",
		},
		TTL:       ttl,
		StoreType: RedisCache,
	}
	return cache, nil
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	if c.TTL == 0 {
		return nil, ErrCacheIsDisabled
	}
	return c.store.Get(ctx, key)
}

func (c *Cache) Set(ctx context.Context, key string, item []byte) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}
	return c.store.Set(ctx, key, item)
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}
	return c.store.Delete(ctx, key)
}

func (c *Cache) Clear(ctx context.Context) error {
	if c.TTL == 0 {
		return ErrCacheIsDisabled
	}
	return c.store.Clear(ctx)
}

// Helper that returns a cache key from the request URL and optional Authorization header.
//
// `url` is the full request URL (including scheme, host, path, and query string).
// Any query parameters that affect the response should be left in the URL.
//
// `authHeader` is the value of the HTTP `Authorization` header, if present.
//
// Returns the full request url with a '-<hash>' appended at the end if an authHeader was provided.
func GetCacheKey(url string, authHeader string) (string, bool) {
	if authHeader == "" {
		return url, false
	}

	hash := sha256.New()
	_, _ = hash.Write([]byte(authHeader))
	hashedAuth := hash.Sum(nil)

	return fmt.Sprintf("%s-%x", url, hashedAuth), true
}
