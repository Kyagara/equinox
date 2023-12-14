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
