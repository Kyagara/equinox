package cache

import (
	"context"
	"errors"

	"github.com/allegro/bigcache/v3"
)

type BigCacheClient interface {
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
	Delete(key string) error
	Reset() error
}

type BigCacheStore struct {
	client BigCacheClient
}

func (s *BigCacheStore) Get(_ctx context.Context, key string) ([]byte, error) {
	item, err := s.client.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
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
