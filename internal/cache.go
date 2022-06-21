package internal

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

type Cache struct {
	// Where the cache will be stored.
	store *bigcache.BigCache
	ttl   int
}

func NewCache(ttl int) (*Cache, error) {
	time := time.Duration(ttl) * time.Second

	bCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time))

	if err != nil {
		return nil, err
	}

	cache := &Cache{store: bCache, ttl: ttl}

	return cache, nil
}

// Adds an item in the cache
func (c *Cache) Set(url string, res []byte) error {
	err := c.store.Set(url, res)

	return err
}

// Gets item from the cache
func (c *Cache) Get(url string) []byte {
	// BigCache returns an error if the item is not found.
	res, _ := c.store.Get(url)

	return res
}

// Clears the cache
func (c *Cache) Clear() error {
	err := c.store.Reset()

	return err
}
