package internal

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httputil"
	"time"
)

type CacheItem struct {
	response []byte
	expire   time.Time
}

type Cache struct {
	items map[string]*CacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: map[string]*CacheItem{},
	}
}

func (c *Cache) Get(url string) (*http.Response, error) {
	item, ok := c.items[url]

	if ok {
		if item.expire.Before(time.Now()) {
			delete(c.items, url)

			return nil, nil
		}

		reader := bufio.NewReader(bytes.NewReader(item.response))

		res, err := http.ReadResponse(reader, nil)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, nil
}

func (c *Cache) Set(url string, res *http.Response) error {
	response, err := httputil.DumpResponse(res, true)

	if err != nil {
		return err
	}

	c.items[url] = &CacheItem{response: response, expire: time.Now().Add(2 * time.Minute)}

	return nil
}
