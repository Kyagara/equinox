package ratelimit_test

import (
	"context"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNewInternalRateLimit(t *testing.T) {
	t.Parallel()

	rateLimit := ratelimit.NewInternalRateLimit(0.99, time.Second)
	require.NotNil(t, rateLimit)

	err := rateLimit.Reserve(context.Background(), zerolog.Nop(), "route", "method")
	require.NoError(t, err)

	// Test invalid values are being replaced with valid ones
	rateLimit = ratelimit.NewInternalRateLimit(-1, -1)
	require.Equal(t, float64(0.99), rateLimit.LimitUsageFactor)
	require.Equal(t, time.Second, rateLimit.IntervalOverhead)
	require.True(t, rateLimit.Enabled)
}

func TestLimits(t *testing.T) {
	t.Parallel()

	limits := ratelimit.NewLimits()
	require.NotNil(t, limits)
	require.NotEmpty(t, limits.App)
	require.NotNil(t, limits.Methods)
}

func TestBucket(t *testing.T) {
	t.Parallel()

	bucket := ratelimit.NewBucket(2*time.Second, 500*time.Millisecond, 10, int(math.Max(1, 10.0*0.99)), 0)
	require.NotNil(t, bucket)
	require.Equal(t, 10, bucket.BaseLimit)
	require.Equal(t, int(math.Max(1, 10.0*0.99)), bucket.Limit)
	require.Equal(t, 0, bucket.Tokens)
	require.Equal(t, 2*time.Second, bucket.Interval)
	require.Equal(t, 500*time.Millisecond, bucket.IntervalOverhead)
	require.Greater(t, bucket.Next, time.Now())
	require.False(t, bucket.IsRateLimited())

	bucket.BaseLimit = 0
	require.False(t, bucket.IsRateLimited())

	bucket.Tokens = 20
	bucket.BaseLimit = 10
	require.True(t, bucket.IsRateLimited())

	bucket.Next = time.Time{}
	bucket.Check()
	require.False(t, bucket.IsRateLimited())
}

func TestReserveAndUpdate(t *testing.T) {
	client, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.NoError(t, err)
	equinoxReq := api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   client.Logger("client_endpoint_method"),
	}

	t.Run("buckets not created", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		require.True(t, r.Enabled)
		require.Equal(t, r.LimitUsageFactor, 0.99)
		require.Equal(t, r.IntervalOverhead, time.Second)
	})

	// These tests should take around 2 seconds each

	t.Run("app rate limited", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"19:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	t.Run("method rate limited", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"100:2,200:2"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"1:2,199:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	t.Run("waiting bucket to reset", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	t.Run("waiting retry after", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"10:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers = http.Header{
			ratelimit.RATE_LIMIT_TYPE_HEADER:      []string{"application"},
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"11:2"},
			ratelimit.RETRY_AFTER_HEADER:          []string{"2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, 2*time.Second)
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})
}
