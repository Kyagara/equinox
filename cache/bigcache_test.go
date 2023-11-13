package cache_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/require"
)

func TestNewBigCache(t *testing.T) {
	ctx := context.Background()
	config := bigcache.DefaultConfig(5 * time.Minute)
	c, err := cache.NewBigCache(ctx, config)
	require.Nil(t, err)
	require.NotNil(t, c, "expecting non nil cache")
	invalidConfig := bigcache.Config{LifeWindow: -1 * time.Minute}
	c, err = cache.NewBigCache(ctx, invalidConfig)
	require.NotNil(t, err)
	require.Nil(t, c)
}

func TestBigCacheMethods(t *testing.T) {
	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))
	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))
	require.NotNil(t, cache, "expecting non-nil BigCache")

	bytes := []byte("data")
	err = cache.Set("test", bytes)
	require.Nil(t, err, "expecting nil error")
	data, err := cache.Get("test")
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, bytes, data, fmt.Sprintf("want data %v, got %v", bytes, data))
	err = cache.Delete("test")
	require.Nil(t, err, "expecting nil error")
	err = cache.Set("test", bytes)
	require.Nil(t, err, "expecting nil error")
	err = cache.Clear()
	require.Nil(t, err, "expecting nil error")
	data, err = cache.Get("test")
	require.Nil(t, err, "expecting nil error")
	require.Empty(t, data)
}
