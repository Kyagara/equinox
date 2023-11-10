package cache_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/cache"
	"github.com/stretchr/testify/require"
)

func TestCacheMethods(t *testing.T) {
	cacheStore := &cache.Cache{}
	require.NotNil(t, cacheStore, "expecting non-nil BigCache")

	bytes := []byte("data")
	err := cacheStore.Set("test", bytes)
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", nil, err))
	_, err = cacheStore.Get("test")
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", nil, err))
	err = cacheStore.Delete("test")
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", nil, err))
	err = cacheStore.Set("test", bytes)
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", nil, err))
	err = cacheStore.Clear()
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", nil, err))
	_, err = cacheStore.Get("test")
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", nil, err))
}
