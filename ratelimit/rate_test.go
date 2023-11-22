package ratelimit_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/stretchr/testify/require"
)

func TestRateLimit(t *testing.T) {
	buckets := ratelimit.NewBuckets()
	require.NotNil(t, buckets)
	require.Empty(t, buckets.App, "app limits not initialized")
	require.Empty(t, buckets.Methods, "methods limits not initialized")
}

func TestRateLimitCheck(t *testing.T) {
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
	}

	t.Run("app rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Buckets)}
		headers := http.Header{
			"X-App-Rate-Limit":          []string{"20:10"},
			"X-App-Rate-Limit-Count":    []string{"19:10"},
			"X-Method-Rate-Limit":       []string{"100:20"},
			"X-Method-Rate-Limit-Count": []string{"1:20"},
		}
		err := r.Update(equinoxReq, &headers)
		require.Nil(t, err)
		err = r.Take(equinoxReq, &headers)
		require.Nil(t, err)
		err = r.Take(equinoxReq, &headers)
		require.Equal(t, fmt.Errorf("app rate limit reached on 'route' route for method 'method'"), err)
	})

	t.Run("method rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Buckets)}
		headers := http.Header{
			"X-App-Rate-Limit":          []string{"20:10,200:100"},
			"X-App-Rate-Limit-Count":    []string{"1:10,1:100"},
			"X-Method-Rate-Limit":       []string{"100:20,200:100"},
			"X-Method-Rate-Limit-Count": []string{"1:20,199:100"},
		}
		err := r.Update(equinoxReq, &headers)
		require.Nil(t, err)
		err = r.Take(equinoxReq, &headers)
		require.Nil(t, err)
		err = r.Take(equinoxReq, &headers)
		require.Equal(t, fmt.Errorf("method rate limit reached on 'route' route for method 'method'"), err)
	})

	t.Run("waiting bucket to reset", func(t *testing.T) {
		r := &ratelimit.RateLimit{Buckets: make(map[any]*ratelimit.Buckets)}
		headers := http.Header{
			"X-App-Rate-Limit":       []string{"20:1"},
			"X-App-Rate-Limit-Count": []string{"20:1"},
		}
		err := r.Update(equinoxReq, &headers)
		require.Nil(t, err)
		err = r.Take(equinoxReq, &headers)
		require.Equal(t, fmt.Errorf("app rate limit reached on 'route' route for method 'method'"), err)
		time.Sleep(2 * time.Second)
		err = r.Take(equinoxReq, &headers)
		require.Nil(t, err)
	})
}
