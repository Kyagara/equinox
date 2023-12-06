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
	s := miniredis.RunT(t)
	require.NotEmpty(t, s, "expecting non-nil miniredis instance")
	ctx := context.Background()
	config := &redis.Options{
		Network: "tcp",
		Addr:    s.Addr(),
	}

	_, err := cache.NewRedis(ctx, nil, 4*time.Minute)
	require.Equal(t, cache.ErrRedisOptionsNil, err)
	cache, err := cache.NewRedis(ctx, config, 4*time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, cache, "expecting non-nil Redis")
}

func TestRedisMethods(t *testing.T) {
	s := miniredis.RunT(t)
	require.NotEmpty(t, s, "expecting non-nil miniredis instance")
	ctx := context.Background()
	config := &redis.Options{
		Network: "tcp",
		Addr:    s.Addr(),
	}

	cache, err := cache.NewRedis(ctx, config, 4*time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, cache, "expecting non-nil Redis")

	bytes := []byte("data")
	err = cache.Set("test", bytes)
	require.NoError(t, err)
	data, err := cache.Get("test")
	require.NoError(t, err)
	require.Equal(t, bytes, data)
	err = cache.Delete("test")
	require.NoError(t, err)
	err = cache.Set("test", bytes)
	require.NoError(t, err)
	err = cache.Clear()
	require.NoError(t, err)
	data, err = cache.Get("test")
	require.NoError(t, err)
	require.Empty(t, data)
}
