package cache_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRedis(t *testing.T) {
	s := miniredis.RunT(t)

	assert.NotNil(t, s, "expecting non-nil miniredis instance")

	ctx := context.Background()

	config := &redis.Options{
		Network: "tcp",
		Addr:    s.Addr(),
	}

	_, err := cache.NewRedis(ctx, nil, 4*time.Minute)

	redisErr := fmt.Errorf("redis options is empty")

	require.Equal(t, redisErr, err, fmt.Sprintf("want err %v, got %v", redisErr, err))

	cache, err := cache.NewRedis(ctx, config, 4*time.Minute)

	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	assert.NotNil(t, cache, "expecting non-nil Redis")
}

func TestRedisMethods(t *testing.T) {
	s := miniredis.RunT(t)

	assert.NotNil(t, s, "expecting non-nil miniredis instance")

	ctx := context.Background()

	config := &redis.Options{
		Network: "tcp",
		Addr:    s.Addr(),
	}

	cache, err := cache.NewRedis(ctx, config, 4*time.Minute)

	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	require.NotNil(t, cache, "expecting non-nil Redis")

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

	assert.Empty(t, data)
}
