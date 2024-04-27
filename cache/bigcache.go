package cache

import (
	"context"

	"github.com/allegro/bigcache/v3"
)

type BigCacheStore struct {
	client *bigcache.BigCache
}

func (s *BigCacheStore) Get(_ctx context.Context, key string) ([]byte, error) {
	item, err := s.client.Get(key)
	if err == bigcache.ErrEntryNotFound {
		return nil, nil
	}
	return item, err
}

func (s *BigCacheStore) Set(_ctx context.Context, key string, value []byte) error {
	return s.client.Set(key, value)
}

func (s *BigCacheStore) Delete(_ctx context.Context, key string) error {
	return s.client.Delete(key)
}

func (s *BigCacheStore) Clear(_ctx context.Context) error {
	return s.client.Reset()
}
