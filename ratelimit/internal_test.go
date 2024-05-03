package ratelimit_test

import (
	"context"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/stretchr/testify/require"
)

func TestNewInternalRateLimit(t *testing.T) {
	t.Parallel()

	// Test if invalid values are being replaced with valid ones
	rateLimit := ratelimit.NewInternalRateLimit(-1, -1)
	require.Equal(t, ratelimit.InternalRateLimit, rateLimit.StoreType)
	require.Equal(t, float64(0.99), rateLimit.LimitUsageFactor)
	require.Equal(t, time.Second, rateLimit.IntervalOverhead)
	require.True(t, rateLimit.Enabled)

	err := rateLimit.Reserve(context.Background(), util.NewTestLogger(), "route", "method")
	require.NoError(t, err)
}

func TestLimits(t *testing.T) {
	t.Parallel()

	limits := ratelimit.NewLimits()
	require.NotNil(t, limits)
	require.NotEmpty(t, limits.App)
	require.NotNil(t, limits.Methods)
	require.Empty(t, limits.Methods)
	require.Equal(t, ratelimit.APP_RATE_LIMIT_TYPE, limits.App.Type)

	limits.App = ratelimit.ParseHeaders(ratelimit.APP_RATE_LIMIT_TYPE, "10:1,10:2", "1:1,1:2", 0.99, time.Second)
	require.NotEmpty(t, limits.App.Buckets)
	limits.Methods["method"] = ratelimit.ParseHeaders(ratelimit.METHOD_RATE_LIMIT_TYPE, "10:1,10:2", "1:1,1:2", 0.99, time.Second)
	require.NotEmpty(t, limits.Methods["method"].Buckets)

	limitsMatch := limits.App.LimitsMatch("10:1,10:2")
	require.True(t, limitsMatch)

	ctx := context.Background()
	logger := util.NewTestLogger()

	err := limits.App.CheckBuckets(ctx, logger, "route")
	require.NoError(t, err)

	limits.App = ratelimit.ParseHeaders(ratelimit.APP_RATE_LIMIT_TYPE, "10:10,10:20", "1000:10,1000:20", 0.99, time.Second)
	require.NotEmpty(t, limits.App.Buckets)

	ctx, c := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer c()

	err = limits.App.CheckBuckets(ctx, logger, "route")
	require.Error(t, err)

	limits.App.SetRetryAfter(10 * time.Second)
	err = limits.App.CheckBuckets(ctx, logger, "route")
	require.Error(t, err)

	limits.Methods["method"].SetRetryAfter(10 * time.Second)
	err = limits.Methods["method"].CheckBuckets(ctx, logger, "route")
	require.Error(t, err)

	limits.App.Buckets[0].BaseLimit = 0
	limitsMatch = limits.App.LimitsMatch("10:1,10:2")
	require.False(t, limitsMatch)

	limits.App.Buckets[0] = nil
	limitsMatch = limits.App.LimitsMatch("10:1,10:2")
	require.False(t, limitsMatch)
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
	t.Parallel()

	equinoxReq := api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   util.NewTestLogger(),
	}

	// These tests should take around 2 seconds each
	t.Run("app and method rate limited", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.RATE_LIMIT_TYPE_HEADER:      []string{"application"},
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"19:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		// App rate limited
		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers = http.Header{
			ratelimit.RATE_LIMIT_TYPE_HEADER:         []string{"method"},
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"10:2"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"19:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		// Method rate limited
		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	// Should take 3 seconds, 2 seconds for the app and 1 for the interval overhead
	t.Run("waiting bucket to reset", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.RATE_LIMIT_TYPE_HEADER:      []string{"application"},
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:2"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	// Should take 5 seconds, 2 for each rate limit and 1 for the interval overhead
	t.Run("waiting retry after", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.RATE_LIMIT_TYPE_HEADER:      []string{"application"},
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"10:2"},
		}

		// No Retry after
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

		// Retry after on application rate limit
		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, 2*time.Second)
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers = http.Header{
			ratelimit.RATE_LIMIT_TYPE_HEADER:         []string{"method"},
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"10:2"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"11:2"},
			ratelimit.RETRY_AFTER_HEADER:             []string{"2"},
		}

		// Retry after on method rate limit
		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, 2*time.Second)
		require.NoError(t, err)

		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	t.Run("CheckBuckets failed in reserve because of deadline", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		r := ratelimit.NewInternalRateLimit(0.99, time.Second)

		// Initializing the rate limit
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)

		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:20"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:20"},
		}

		err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
		require.NoError(t, err)

		ctx, c := context.WithDeadline(ctx, time.Now().Add(time.Second))
		defer c()
		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.Error(t, err)
	})
}
