package cache_test

import (
	"context"
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
	require.NoError(t, err)
	require.NotEmpty(t, c)
	invalidConfig := bigcache.Config{LifeWindow: -1 * time.Minute}
	c, err = cache.NewBigCache(ctx, invalidConfig)
	require.NotEmpty(t, err)
	require.Nil(t, c)
}

func TestBigCacheMethods(t *testing.T) {
	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))
	require.NoError(t, err)
	require.NotEmpty(t, cache)

	bytes := []byte("data")
	err = cache.Set(ctx, "test", bytes)
	require.NoError(t, err)
	data, err := cache.Get(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, bytes, data)
	err = cache.Delete(ctx, "test")
	require.NoError(t, err)
	err = cache.Set(ctx, "test", bytes)
	require.NoError(t, err)
	err = cache.Clear(ctx)
	require.NoError(t, err)
	data, err = cache.Get(ctx, "test")
	require.NoError(t, err)
	require.Empty(t, data)
}
