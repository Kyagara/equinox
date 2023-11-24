package ratelimit_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/stretchr/testify/require"
)

func TestRateLimit(t *testing.T) {
	buckets := ratelimit.NewLimits()
	require.NotNil(t, buckets)
	require.Empty(t, buckets.App, "app limits not initialized")
	require.Empty(t, buckets.Methods, "methods limits not initialized")
}

func TestRateLimitCheck(t *testing.T) {
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
	}

	t.Run("buckets not created", func(t *testing.T) {
		r := &ratelimit.RateLimit{
			Buckets: make(map[any]*ratelimit.Limits),
		}
		require.Nil(t, r.Buckets[equinoxReq.Route])
		ctx := context.Background()
		err := r.Check(ctx, equinoxReq)
		require.Nil(t, err)
		require.NotNil(t, r.Buckets[equinoxReq.Route].Methods)
		require.NotNil(t, r.Buckets[equinoxReq.Route].Methods[equinoxReq.MethodID])
	})

	t.Run("app rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:10"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"19:10"},
		}
		ctx := context.Background()
		r.Update(equinoxReq, &headers)
		err := r.Check(ctx, equinoxReq)
		require.Equal(t, fmt.Errorf("app rate limit reached on 'route' route for method 'method'. exceeded the bucket's limit 20"), err)
	})

	t.Run("method rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"100:20,200:100"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"1:20,199:100"},
		}
		ctx := context.Background()
		r.Update(equinoxReq, &headers)
		err := r.Check(ctx, equinoxReq)
		require.Equal(t, fmt.Errorf("method rate limit reached on 'route' route for method 'method'. exceeded the bucket's limit 200"), err)
	})

	t.Run("waiting bucket to reset", func(t *testing.T) {
		r := &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:1"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:1"},
		}
		ctx := context.Background()
		r.Update(equinoxReq, &headers)
		err := r.Check(ctx, equinoxReq)
		require.Equal(t, fmt.Errorf("app rate limit reached on 'route' route for method 'method'. exceeded the bucket's limit 20"), err)
		time.Sleep(2 * time.Second)
		err = r.Check(ctx, equinoxReq)
		require.Nil(t, err)
	})
}
