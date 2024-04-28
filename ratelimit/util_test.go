package ratelimit_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kyagara/equinox/ratelimit"
	"github.com/stretchr/testify/require"
)

func TestCheckRetryAfter(t *testing.T) {
	t.Parallel()

	delay := ratelimit.GetRetryAfterHeader("")
	require.Equal(t, time.Second, delay)

	delay = ratelimit.GetRetryAfterHeader("asdf")
	require.Equal(t, time.Second, delay)

	delay = ratelimit.GetRetryAfterHeader("10")
	require.Equal(t, 10*time.Second, delay)
}

func TestWaitN(t *testing.T) {
	t.Parallel()

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
		ctx, c := context.WithTimeout(ctx, duration)
		defer c()
		estimated := time.Now().Add(10 * time.Second)

		err := ratelimit.WaitN(ctx, estimated, duration)
		require.Equal(t, err, ratelimit.ErrContextDeadlineExceeded)
	})

	t.Run("context canceled", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		estimated := time.Now().Add(10 * time.Second)
		duration := 5 * time.Second

		go func() {
			time.Sleep(2 * time.Second)
			cancel()
		}()

		err := ratelimit.WaitN(ctx, estimated, duration)
		require.Equal(t, context.Canceled, err)
	})
}

func TestParseHeaders(t *testing.T) {
	t.Parallel()

	emptyLimit := ratelimit.NewLimit(ratelimit.APP_RATE_LIMIT_TYPE)
	require.Empty(t, emptyLimit.Buckets)
	require.Len(t, emptyLimit.Buckets, 0)
	require.Equal(t, emptyLimit.RetryAfter, time.Duration(0))
	require.Equal(t, emptyLimit.Type, ratelimit.APP_RATE_LIMIT_TYPE)

	limit := ratelimit.ParseHeaders("", "", ratelimit.APP_RATE_LIMIT_TYPE, 0.99, time.Second)
	require.Equal(t, ratelimit.APP_RATE_LIMIT_TYPE, limit.Type)
	require.Empty(t, limit.Buckets)
	require.Equal(t, emptyLimit, limit)
	require.Len(t, limit.Buckets, 0)

	limit = ratelimit.ParseHeaders("10:10,10:20", "1000:10,1000:20", ratelimit.METHOD_RATE_LIMIT_TYPE, 0.99, time.Second)
	require.NotEmpty(t, limit.Buckets)
	require.Equal(t, ratelimit.METHOD_RATE_LIMIT_TYPE, limit.Type)
}
