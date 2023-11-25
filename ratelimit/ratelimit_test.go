package ratelimit_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
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
	client, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   client.Logger("client_endpoint_method"),
	}

	t.Run("buckets not created", func(t *testing.T) {
		r := &ratelimit.RateLimit{
			Limits: make(map[any]*ratelimit.Limits),
		}
		require.Nil(t, r.Limits[equinoxReq.Route])
		ctx := context.Background()
		err := r.Take(ctx, equinoxReq)
		require.Nil(t, err)
		require.NotNil(t, r.Limits[equinoxReq.Route].Methods)
		require.NotNil(t, r.Limits[equinoxReq.Route].Methods[equinoxReq.MethodID])
	})

	t.Run("app rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Limits: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"19:2"},
		}
		ctx := context.Background()
		r.Update(equinoxReq, &headers)
		err := r.Take(ctx, equinoxReq)
		require.Nil(t, err)
	})

	t.Run("method rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Limits: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"100:2,200:2"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"1:2,199:2"},
		}
		ctx := context.Background()
		r.Update(equinoxReq, &headers)
		err := r.Take(ctx, equinoxReq)
		require.Nil(t, err)
	})

	t.Run("waiting bucket to reset", func(t *testing.T) {
		r := &ratelimit.RateLimit{Limits: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:2"},
		}
		ctx := context.Background()
		r.Update(equinoxReq, &headers)
		// This test should take around 2 seconds
		err := r.Take(ctx, equinoxReq)
		require.Nil(t, err)
	})
}
