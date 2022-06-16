package internal

import (
	"sync"
	"time"
)

type Cache struct {
	items map[string]*CacheItem
	mutex sync.Mutex
}

type CacheItem struct {
	response []byte
	access   int64
}

func NewCache(ttl int64) *Cache {
	cache := &Cache{
		items: map[string]*CacheItem{},
		mutex: sync.Mutex{},
	}

	go func() {
		for now := range time.Tick(time.Second) {
			cache.mutex.Lock()

			for url, item := range cache.items {
				if now.Unix()-item.access > ttl {
					delete(cache.items, url)
				}
			}

			cache.mutex.Unlock()
		}
	}()

	return cache
}

// Adds a http.Response in the cache
func (c *Cache) Set(url string, res []byte) error {
	c.mutex.Lock()

	c.items[url] = &CacheItem{
		response: res,
		access:   time.Now().Unix(),
	}

	c.mutex.Unlock()

	return nil
}

// Gets a http.Response from the cache
func (c *Cache) Get(url string) (*CacheItem, error) {
	c.mutex.Lock()

	item, ok := c.items[url]

	c.mutex.Unlock()

	if ok {
		return item, nil
	}

	return nil, nil
}

// Clears the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k := range c.items {
		delete(c.items, k)
	}
}
