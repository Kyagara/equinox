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
	"github.com/stretchr/testify/require"
)

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

	err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
	require.NoError(t, err)

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
	err = r.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, headers, time.Duration(0))
	require.NoError(t, err)

	// Should have no issue
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.NoError(t, err)

	// The buckets should've been updated, so this request should be rate limited and block
	// Since we have a deadline, this should return a DeadlineExceeded error
	err = r.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)
}
