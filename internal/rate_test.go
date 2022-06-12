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
	rateLimit := internal.NewRateLimiter()

	assert.NotNil(t, rateLimit, "expecting non-nil RateLimit")
}

func TestRateLimitSetGet(t *testing.T) {
	rateLimit := internal.NewRateLimiter()

	require.NotNil(t, rateLimit, "expecting non-nil RateLimit")

	rate := &internal.Rate{
		Seconds:      100,
		SecondsLimit: 10,
		SecondsCount: 1000,
		Ticker:       &time.Ticker{},
		Mutex:        &sync.Mutex{},
	}

	rate.Ticker = time.NewTicker(time.Duration(rate.Seconds) * time.Second)

	rateLimit.Set("testEndpoint", "testMethod", rate)

	check := rateLimit.Get("testEndpoint", "testMethod")

	require.NotNil(t, check, "expecting non-nil Rate")

	require.Equal(t, 1000, check.SecondsCount, "expecting non-nil Rate")

	rate.SecondsCount = 500

	rateLimit.Set("testEndpoint", "testMethod", rate)

	check = rateLimit.Get("testEndpoint", "testMethod")

	require.NotNil(t, check, "expecting non-nil Rate")

	assert.Equal(t, 500, check.SecondsCount, "expecting non-nil Rate")
}

func TestRateLimitParseHeaders(t *testing.T) {
	rateLimit := internal.NewRateLimiter()

	require.NotNil(t, rateLimit, "expecting non-nil RateLimit")

	headers := map[string][]string{}

	headers["X-App-Rate-Limit"] = []string{"1000:10,60000:600"}
	headers["X-App-Rate-Limit-Count"] = []string{"1000:10,60000:600"}

	headers["X-Method-Rate-Limit"] = []string{"1000:10,60000:600"}
	headers["X-Method-Rate-Limit-Count"] = []string{"1000:10,60000:600"}

	rate := rateLimit.ParseHeaders(headers, "X-Method-Rate-Limit", "X-Method-Rate-Limit-Count")

	require.NotNil(t, rate, "expecting non-nil Rate")

	require.Equal(t, 1000, rate.SecondsCount)

	rate = rateLimit.ParseHeaders(headers, "X-App-Rate-Limit", "X-App-Rate-Limit-Count")

	require.NotNil(t, rate, "expecting non-nil Rate")

	assert.Equal(t, 1000, rate.SecondsCount)
}
