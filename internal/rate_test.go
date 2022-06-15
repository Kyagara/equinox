package internal_test

import (
	"sync"
	"testing"
	"time"

	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRateLimit(t *testing.T) {
	rateLimit := internal.NewRateLimit()

	assert.NotNil(t, rateLimit, "expecting non-nil RateLimit")
}

func TestRateLimitSetGet(t *testing.T) {
	rateLimit := internal.NewRateLimit()

	require.NotNil(t, rateLimit, "expecting non-nil RateLimit")

	rate := &internal.Rate{
		Seconds: &internal.RateTiming{Time: 100, Limit: 10, Count: 1000, Ticker: &time.Ticker{}, Mutex: &sync.Mutex{}},
		Minutes: &internal.RateTiming{Time: 100, Limit: 10, Count: 1000, Ticker: &time.Ticker{}, Mutex: &sync.Mutex{}},
	}

	rate.Seconds.Ticker = time.NewTicker(time.Duration(rate.Seconds.Time) * time.Second)

	rateLimit.Set("testEndpoint", "testMethod", rate)

	check := rateLimit.Get("testEndpoint", "testMethod")

	require.NotNil(t, check, "expecting non-nil Rate")

	require.Equal(t, 1000, check.Seconds.Count, "expecting non-nil Rate")

	rate.Seconds.Count = 500

	rateLimit.Set("testEndpoint", "testMethod", rate)

	check = rateLimit.Get("testEndpoint", "testMethod")

	require.NotNil(t, check, "expecting non-nil Rate")

	assert.Equal(t, 500, check.Seconds.Count, "expecting non-nil Rate")
}

func TestRateLimitParseHeaders(t *testing.T) {
	rateLimit := internal.NewRateLimit()

	require.NotNil(t, rateLimit, "expecting non-nil RateLimit")

	headers := map[string][]string{}

	headers["X-App-Rate-Limit"] = []string{"1000:10,60000:600"}
	headers["X-App-Rate-Limit-Count"] = []string{"1000:10,60000:600"}

	headers["X-Method-Rate-Limit"] = []string{"1000:10,60000:600"}
	headers["X-Method-Rate-Limit-Count"] = []string{"1000:10,60000:600"}

	rate := internal.ParseHeaders(headers, "X-Method-Rate-Limit", "X-Method-Rate-Limit-Count")

	require.NotNil(t, rate, "expecting non-nil Rate")

	require.Equal(t, 1000, rate.Seconds.Count)

	rate = internal.ParseHeaders(headers, "X-App-Rate-Limit", "X-App-Rate-Limit-Count")

	require.NotNil(t, rate, "expecting non-nil Rate")

	assert.Equal(t, 1000, rate.Seconds.Count)
}
