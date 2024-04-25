package ratelimit_test

import (
	"context"
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

func TestNewLimits(t *testing.T) {
	t.Parallel()

	limits := ratelimit.NewLimits()
	require.NotNil(t, limits)
	require.NotEmpty(t, limits.App)
	require.NotNil(t, limits.Methods)
}

func TestNewInternalRateLimit(t *testing.T) {
	t.Parallel()

	rateLimit := ratelimit.NewInternalRateLimit(0.99, time.Second)
	require.NotNil(t, rateLimit)
	require.Empty(t, rateLimit.Route)
	require.True(t, rateLimit.Enabled)
	require.Nil(t, rateLimit.Route["route"])

	err := rateLimit.Reserve(context.Background(), zerolog.Nop(), "route", "method")
	require.NoError(t, err)
	require.NotNil(t, rateLimit.Route["route"].App)
	require.NotEmpty(t, rateLimit.Route["route"])
	require.NotNil(t, rateLimit.Route["route"].Methods)
	require.NotEmpty(t, rateLimit.Route["route"].Methods["method"])

	// Test invalid values are being replaced with valid ones
	rateLimit = ratelimit.NewInternalRateLimit(-1, -1)
	require.Equal(t, float64(0.99), rateLimit.LimitUsageFactor)
	require.Equal(t, time.Second, rateLimit.IntervalOverhead)
}

func TestRateLimitCheck(t *testing.T) {
	client, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.NoError(t, err)
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   client.Logger("client_endpoint_method"),
	}

	t.Run("buckets not created", func(t *testing.T) {
		t.Parallel()

		r := &ratelimit.RateLimit{
			Route: make(map[string]*ratelimit.Limits),
		}
		require.Nil(t, r.Route[equinoxReq.Route])
		ctx := context.Background()
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
		require.NotNil(t, r.Route[equinoxReq.Route].Methods)
		require.NotNil(t, r.Route[equinoxReq.Route].Methods[equinoxReq.MethodID])
	})

	// These tests should take around 2 seconds each

	t.Run("app rate limited", func(t *testing.T) {
		t.Parallel()

		r := &ratelimit.RateLimit{Route: make(map[string]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"19:2"},
		}
		ctx := context.Background()
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
		r.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers)
		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	t.Run("method rate limited", func(t *testing.T) {
		t.Parallel()

		r := &ratelimit.RateLimit{Route: make(map[string]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"100:2,200:2"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"1:2,199:2"},
		}
		ctx := context.Background()
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
		r.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers)
		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})

	t.Run("waiting bucket to reset", func(t *testing.T) {
		t.Parallel()

		r := &ratelimit.RateLimit{Route: make(map[string]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:2"},
		}
		ctx := context.Background()
		err := r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
		r.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers)
		err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		require.NoError(t, err)
	})
}

func TestLimitsDontMatch(t *testing.T) {
	t.Parallel()

	config := util.NewTestEquinoxConfig()
	config.RateLimit = ratelimit.NewInternalRateLimit(0.99, time.Second)
	client, err := internal.NewInternalClient(config)
	require.NoError(t, err)
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   client.Logger("client_endpoint_method"),
	}

	r := config.RateLimit
	headers := http.Header{
		ratelimit.APP_RATE_LIMIT_HEADER:          []string{"20:2"},
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    []string{"1:2"},
		ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"20:2"},
		ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"1:2"},
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	// Used here just to populate the Limits
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.NoError(t, err)

	r.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers)

	// Both requests shouldn't be rate limited or block
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.NoError(t, err)
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.NoError(t, err)

	headers = http.Header{
		ratelimit.APP_RATE_LIMIT_HEADER:       []string{"4:2"},
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"1:2"},
	}
	// Updating limits
	r.Update(equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers)

	// Should have no issue
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.NoError(t, err)

	// The buckets should've been updated, so this request should be rate limited and block
	// Since we have a deadline, this should return a DeadlineExceeded error
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)
}

func TestCheckRetryAfter(t *testing.T) {
	t.Parallel()

	r := &ratelimit.RateLimit{
		Route: map[string]*ratelimit.Limits{
			"route": {
				App: ratelimit.NewLimit(ratelimit.APP_RATE_LIMIT_TYPE),
				Methods: map[string]*ratelimit.Limit{
					"method": ratelimit.NewLimit(ratelimit.METHOD_RATE_LIMIT_TYPE),
				},
			},
		},
	}

	headers := http.Header{}
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
	}

	headers.Set(ratelimit.RETRY_AFTER_HEADER, "")
	delay := r.CheckRetryAfter(equinoxReq.Route, equinoxReq.MethodID, headers)
	require.Equal(t, 2*time.Second, delay)

	headers.Set(ratelimit.RETRY_AFTER_HEADER, "asdf")
	delay = r.CheckRetryAfter(equinoxReq.Route, equinoxReq.MethodID, headers)
	require.Equal(t, 2*time.Second, delay)

	headers.Set(ratelimit.RETRY_AFTER_HEADER, "10")
	delay = r.CheckRetryAfter(equinoxReq.Route, equinoxReq.MethodID, headers)
	require.Equal(t, 10*time.Second, delay)
}

func TestWaitN(t *testing.T) {
	t.Run("deadline not exceeded", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		estimated := time.Now().Add(time.Second)
		duration := 2 * time.Second
		err := ratelimit.WaitN(ctx, estimated, duration)
		require.NoError(t, err)
	})

	t.Run("deadline exceeded", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		duration := time.Second
		ctx, cancel := context.WithTimeout(ctx, duration)
		estimated := time.Now().Add(2 * time.Second)
		err := ratelimit.WaitN(ctx, estimated, duration)
		require.Equal(t, err, ratelimit.ErrContextDeadlineExceeded)
		cancel()
	})

	t.Run("deadline exceeded", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		estimated := time.Now().Add(10 * time.Second)
		duration := 5 * time.Second
		cancel()
		err := ratelimit.WaitN(ctx, estimated, duration)
		require.Equal(t, err, context.Canceled)
	})
}
