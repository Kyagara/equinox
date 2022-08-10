package rate_limit_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/rate_limit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInternalRateLimit(t *testing.T) {
	rateLimit, err := rate_limit.NewInternalRateLimit()

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rateLimit, "expecting non-nil RateLimit")
}

func TestInternalRateLimitSetGet(t *testing.T) {
	rateLimit, err := rate_limit.NewInternalRateLimit()

	require.Nil(t, err, "expecting nil error")

	headers := &http.Header{}

	headers.Add(rate_limit.AppRateLimitHeader, "1300:60,1300:60")
	headers.Add(rate_limit.AppRateLimitCountHeader, "1:60,1300:60")

	headers.Add(rate_limit.MethodRateLimitHeader, "1300:60")
	headers.Add(rate_limit.MethodRateLimitCountHeader, "1300:60")

	err = rateLimit.Set("testRoute", "testEndpoint", "testMethod", headers)

	require.Nil(t, err, "expecting nil error")

	rate, err := rateLimit.Get("testRoute", "testEndpoint", "testMethod")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rate, "expecting non-nil Rate")

	assert.Equal(t, 1300, rate.Seconds.Count, "expecting rate seconds to be equal to 1300")
}

func TestInternalRateLimitAppSetGet(t *testing.T) {
	rateLimit, err := rate_limit.NewInternalRateLimit()

	require.Nil(t, err, "expecting nil error")

	headers := &http.Header{}

	headers.Add(rate_limit.AppRateLimitHeader, "1300:60,1300:60")
	headers.Add(rate_limit.AppRateLimitCountHeader, "1:60,1300:60")

	headers.Add(rate_limit.MethodRateLimitHeader, "1300:60")
	headers.Add(rate_limit.MethodRateLimitCountHeader, "1300:60")

	err = rateLimit.SetAppRate("testRoute", headers)

	require.Nil(t, err, "expecting nil error")

	rate, err := rateLimit.GetAppRate("testRoute")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rate, "expecting non-nil Rate")

	assert.Equal(t, 1, rate.Seconds.Count, "expecting rate seconds to be equal to 1")
}

func TestInternalRateLimitExpiring(t *testing.T) {
	rateLimit, err := rate_limit.NewInternalRateLimit()

	require.Nil(t, err, "expecting nil error")

	headers := &http.Header{}

	headers.Add(rate_limit.AppRateLimitHeader, "1300:2,1300:60")
	headers.Add(rate_limit.AppRateLimitCountHeader, "1300:2,5:60")

	err = rateLimit.SetAppRate("testRoute", headers)

	require.Nil(t, err, "expecting nil error")

	rate, err := rateLimit.GetAppRate("testRoute")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rate, "expecting non-nil Rate")

	limited, err := rateLimit.IsRateLimited(rate)

	require.Nil(t, err, "expecting nil error")

	require.Equal(t, true, limited, "expecting to be rate limited")

	time.Sleep(3 * time.Second)

	limited, err = rateLimit.IsRateLimited(rate)

	require.Nil(t, err, "expecting nil error")

	assert.Equal(t, false, limited, "expecting to not be rate limited")

	// Verifying that the rate was changed

	rate, err = rateLimit.GetAppRate("testRoute")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rate, "expecting non-nil Rate")

	limited, err = rateLimit.IsRateLimited(rate)

	require.Nil(t, err, "expecting nil error")

	assert.Equal(t, false, limited, "expecting to not be rate limited")
}
