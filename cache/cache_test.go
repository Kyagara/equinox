package cache_test

import (
	"context"
	"os"
	"testing"

	"github.com/Kyagara/equinox/cache"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestCacheMethods(t *testing.T) {
	t.Parallel()

	// Cache is disabled
	cacheStore := &cache.Cache{}
	require.NotNil(t, cacheStore)

	ctx := context.Background()

	key := "https://euw1.api.riotgames.com"
	response := []byte("{data: 123}")

	err := cacheStore.Set(ctx, key, response)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	_, err = cacheStore.Get(ctx, key)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Delete(ctx, key)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Clear(ctx)
	require.Equal(t, cache.ErrCacheIsDisabled, err)

	cacheStore.MarshalZerologObject(&zerolog.Event{})

	cacheStore.StoreType = cache.BigCache
	cacheStore.TTL = 1

	var logger zerolog.Logger
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Object("cache", cacheStore).Logger()
	logger.Info().Msg("Testing cache marshal")
}
