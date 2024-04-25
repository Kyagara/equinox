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
	t.Parallel()
	ctx := context.Background()
	config := bigcache.DefaultConfig(5 * time.Minute)
	c, err := cache.NewBigCache(ctx, config)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	invalidConfig := bigcache.Config{LifeWindow: -time.Minute}
	c, err = cache.NewBigCache(ctx, invalidConfig)
	require.NotEmpty(t, err)
	require.Nil(t, c)
}

func TestBigCacheMethods(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))
	require.NoError(t, err)
	require.NotEmpty(t, cache)

	// Data
	key := "https://euw1.api.riotgames.com"
	response := []byte("{data: 123}")

	err = cache.Set(ctx, key, response)
	require.NoError(t, err)

	retrievedData, err := cache.Get(ctx, key)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedData)

	err = cache.Delete(ctx, key)
	require.NoError(t, err)

	// Get on deleted key
	retrievedData, err = cache.Get(ctx, key)
	require.NoError(t, err)
	require.Empty(t, retrievedData)

	err = cache.Set(ctx, key, response)
	require.NoError(t, err)

	err = cache.Clear(ctx)
	require.NoError(t, err)

	// Get on cleared cache
	retrievedData, err = cache.Get(ctx, "test")
	require.NoError(t, err)
	require.Empty(t, retrievedData)
}
