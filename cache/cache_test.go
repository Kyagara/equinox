package cache_test

import (
	"testing"

	"github.com/Kyagara/equinox/cache"
	"github.com/stretchr/testify/require"
)

func TestCacheMethods(t *testing.T) {
	cacheStore := &cache.Cache{}
	require.NotNil(t, cacheStore)

	bytes := []byte("data")
	err := cacheStore.Set("test", bytes)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	_, err = cacheStore.Get("test")
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Delete("test")
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Set("test", bytes)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Clear()
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	_, err = cacheStore.Get("test")
	require.Equal(t, cache.ErrCacheIsDisabled, err)
}
