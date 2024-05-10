package ratelimit_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/v2/ratelimit"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/stretchr/testify/require"
)

func TestRateLimitMethods(t *testing.T) {
	t.Parallel()

	// Rate limit is disabled
	rateStore := &ratelimit.RateLimit{}
	require.NotNil(t, rateStore)

	ctx := context.Background()

	err := rateStore.Reserve(ctx, util.NewTestLogger(), "route", "method", false)
	require.Equal(t, ratelimit.ErrRateLimitIsDisabled, err)

	err = rateStore.Update(ctx, util.NewTestLogger(), "route", "method", http.Header{}, time.Duration(0))
	require.Equal(t, ratelimit.ErrRateLimitIsDisabled, err)
}

func TestParseHeaders(t *testing.T) {
	t.Parallel()

	limit := ratelimit.ParseHeaders(ratelimit.APP_RATE_LIMIT_TYPE, "", "", 0.99, time.Second)
	require.Equal(t, ratelimit.APP_RATE_LIMIT_TYPE, limit.Type)
	require.Empty(t, limit.Buckets)
	require.Len(t, limit.Buckets, 0)

	limit = ratelimit.ParseHeaders(ratelimit.METHOD_RATE_LIMIT_TYPE, "10:10,10:20", "1000:10,1000:20", 0.99, time.Second)
	require.NotEmpty(t, limit.Buckets)
	require.Equal(t, ratelimit.METHOD_RATE_LIMIT_TYPE, limit.Type)
}
