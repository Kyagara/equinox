package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestNewRedis(t *testing.T) {
	t.Parallel()
	s := miniredis.RunT(t)
	require.NotEmpty(t, s)
	ctx := context.Background()
	config := &redis.Options{
		Network: "tcp",
		Addr:    s.Addr(),
	}

	_, err := cache.NewRedis(ctx, nil, 4*time.Minute)
	require.Equal(t, cache.ErrRedisOptionsNil, err)
	cache, err := cache.NewRedis(ctx, config, 4*time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, cache)
}

func TestRedisMethods(t *testing.T) {
	t.Parallel()
	s := miniredis.RunT(t)
	require.NotEmpty(t, s)
	ctx := context.Background()
	config := &redis.Options{
		Network: "tcp",
		Addr:    s.Addr(),
	}

	cache, err := cache.NewRedis(ctx, config, 4*time.Minute)
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
