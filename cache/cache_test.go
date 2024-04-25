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
	cacheStore := &cache.Cache{}
	require.NotNil(t, cacheStore)

	ctx := context.Background()

	bytes := []byte("data")
	err := cacheStore.Set(ctx, "test", bytes)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	_, err = cacheStore.Get(ctx, "test")
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Delete(ctx, "test")
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Set(ctx, "test", bytes)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	err = cacheStore.Clear(ctx)
	require.Equal(t, cache.ErrCacheIsDisabled, err)
	_, err = cacheStore.Get(ctx, "test")
	require.Equal(t, cache.ErrCacheIsDisabled, err)

	cacheStore.MarshalZerologObject(&zerolog.Event{})

	var logger zerolog.Logger
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Object("cache", cacheStore).Logger()
	logger.Info().Msg("Testing cache marshal")
}
