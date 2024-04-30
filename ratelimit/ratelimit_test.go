package ratelimit_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestRateLimitMethods(t *testing.T) {
	t.Parallel()

	// Rate limit is disabled
	rateStore := &ratelimit.RateLimit{}
	require.NotNil(t, rateStore)

	ctx := context.Background()

	err := rateStore.Reserve(ctx, zerolog.Nop(), "route", "method")
	require.Equal(t, ratelimit.ErrRateLimitIsDisabled, err)

	err = rateStore.Update(ctx, zerolog.Nop(), "route", "method", http.Header{}, time.Duration(0))
	require.Equal(t, ratelimit.ErrRateLimitIsDisabled, err)

	rateStore.MarshalZerologObject(&zerolog.Event{})

	rateStore.StoreType = ratelimit.InternalRateLimit
	rateStore.Enabled = true

	var logger zerolog.Logger
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Object("ratelimit", rateStore).Logger()
	logger.Info().Msg("Testing rate limit marshal")
}

func TestParseHeaders(t *testing.T) {
	t.Parallel()

	emptyLimit := ratelimit.NewLimit(ratelimit.APP_RATE_LIMIT_TYPE)
	require.Empty(t, emptyLimit.Buckets)
	require.Len(t, emptyLimit.Buckets, 0)
	require.Equal(t, emptyLimit.RetryAfter, time.Duration(0))
	require.Equal(t, emptyLimit.Type, ratelimit.APP_RATE_LIMIT_TYPE)

	limit := ratelimit.ParseHeaders(ratelimit.APP_RATE_LIMIT_TYPE, "", "", 0.99, time.Second)
	require.Equal(t, ratelimit.APP_RATE_LIMIT_TYPE, limit.Type)
	require.Empty(t, limit.Buckets)
	require.Equal(t, emptyLimit, limit)
	require.Len(t, limit.Buckets, 0)

	limit = ratelimit.ParseHeaders(ratelimit.METHOD_RATE_LIMIT_TYPE, "10:10,10:20", "1000:10,1000:20", 0.99, time.Second)
	require.NotEmpty(t, limit.Buckets)
	require.Equal(t, ratelimit.METHOD_RATE_LIMIT_TYPE, limit.Type)
}
