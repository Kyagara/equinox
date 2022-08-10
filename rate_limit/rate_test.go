package rate_limit_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/rate_limit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateLimitIsDisabled(t *testing.T) {
	rateLimit := rate_limit.RateLimit{Enabled: false, StoreType: rate_limit.InternalRateLimiter}

	_, err := rateLimit.Get("", "", "")

	require.Equal(t, rate_limit.ErrRateLimitingIsDisabled, err, fmt.Sprintf("want err %v, got %v", rate_limit.ErrRateLimitingIsDisabled, err))

	_, err = rateLimit.GetAppRate("")

	require.Equal(t, rate_limit.ErrRateLimitingIsDisabled, err, fmt.Sprintf("want err %v, got %v", rate_limit.ErrRateLimitingIsDisabled, err))

	err = rateLimit.Set("", "", "", &http.Header{})

	require.Equal(t, rate_limit.ErrRateLimitingIsDisabled, err, fmt.Sprintf("want err %v, got %v", rate_limit.ErrRateLimitingIsDisabled, err))

	_, err = rateLimit.IsRateLimited(&rate_limit.Rate{})

	require.Equal(t, rate_limit.ErrRateLimitingIsDisabled, err, fmt.Sprintf("want err %v, got %v", rate_limit.ErrRateLimitingIsDisabled, err))

	err = rateLimit.SetAppRate("", &http.Header{})

	assert.Equal(t, rate_limit.ErrRateLimitingIsDisabled, err, fmt.Sprintf("want err %v, got %v", rate_limit.ErrRateLimitingIsDisabled, err))
}

func TestRateLimitParseHeaders(t *testing.T) {
	headers := &http.Header{}

	headers.Add(rate_limit.AppRateLimitHeader, "20:1,120:60")
	headers.Add(rate_limit.AppRateLimitCountHeader, "1:1,120:60")

	headers.Add(rate_limit.MethodRateLimitHeader, "120:60,1:120")
	headers.Add(rate_limit.MethodRateLimitCountHeader, "1:60,1:120")

	rate, err := rate_limit.ParseHeaders(headers, "X-App-Rate-Limit", rate_limit.AppRateLimitCountHeader)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rate, "expecting non-nil Rate")

	require.Equal(t, 1, rate.Seconds.Count)

	rate, err = rate_limit.ParseHeaders(headers, rate_limit.MethodRateLimitHeader, rate_limit.MethodRateLimitCountHeader)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rate, "expecting non-nil Rate")

	assert.Equal(t, 1, rate.Minutes.Count, "expecting rate minutes to be equal to 1")
}

func TestIsRateLimited(t *testing.T) {
	rateLimit, err := rate_limit.NewInternalRateLimit()

	require.Nil(t, err, "expecting nil error")

	headers := &http.Header{}

	headers.Add("X-App-Rate-Limit", "1300:60,1300:60")
	headers.Add(rate_limit.AppRateLimitCountHeader, "1:60,1300:60")

	headers.Add(rate_limit.MethodRateLimitHeader, "1300:60")
	headers.Add(rate_limit.MethodRateLimitCountHeader, "1300:60")

	err = rateLimit.SetAppRate("testRoute", headers)

	require.Nil(t, err, "expecting nil error")

	rate, err := rateLimit.GetAppRate("testRoute")

	require.Nil(t, err, "expecting nil error")

	isRateLimited, err := rateLimit.IsRateLimited(rate)

	require.Nil(t, err, "expecting nil error")

	assert.Equal(t, true, isRateLimited, "expecting to be rate limited")
}
