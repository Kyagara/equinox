package cache_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox/cache"
	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBigCache(t *testing.T) {
	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))

	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	assert.NotNil(t, cache, "expecting non-nil BigCache")
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

	assert.Empty(t, data)
}
