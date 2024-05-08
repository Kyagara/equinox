package cache_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/test/util"
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

	logger := util.NewTestLogger()
	logger.Debug().Object("cache", cacheStore).Msg("Testing cache marshal")
}

func TestGetCacheKey(t *testing.T) {
	t.Parallel()

	req := &http.Request{
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
		Header: http.Header{},
	}

	equinoxReq := api.EquinoxRequest{Request: req}
	equinoxReq.URL = req.URL.String()

	hash, isRSO := cache.GetCacheKey(equinoxReq.URL, equinoxReq.Request.Header.Get("Authorization"))
	require.Equal(t, "http://example.com/path", hash)
	require.False(t, isRSO)

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghij")
	hash, isRSO = cache.GetCacheKey(equinoxReq.URL, equinoxReq.Request.Header.Get("Authorization"))
	require.Equal(t, "http://example.com/path-ec2cc2a7cbc79c8d8def89cb9b9a1bccf4c2efc56a9c8063f9f4ae806f08c4d7", hash)
	require.True(t, isRSO)
}
