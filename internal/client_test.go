package internal_test

import (
	"context"
	"time"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/clients/lol"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/Kyagara/equinox/v2/ratelimit"
	"github.com/Kyagara/equinox/v2/test/util"
	jsonv2 "github.com/go-json-experiment/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestNewInternalClient(t *testing.T) {
	t.Parallel()

	internalClient := util.NewTestInternalClient(t)
	require.False(t, internalClient.IsCacheEnabled)
	require.False(t, internalClient.IsRateLimitEnabled)
	require.False(t, internalClient.IsRetryEnabled)

	config := util.NewTestEquinoxConfig()

	internalClient = util.NewTestCustomInternalClient(t, config)
	require.False(t, internalClient.IsCacheEnabled)
	require.False(t, internalClient.IsRateLimitEnabled)
	require.False(t, internalClient.IsRetryEnabled)

	config = util.NewTestEquinoxConfig()
	config.Retry.MaxRetries = 1
	config.Key = ""

	_, err := internal.NewInternalClient(config, nil, nil, nil)
	require.Equal(t, internal.ErrKeyNotProvided, err)

	config.Key = "RGAPI-TEST"

	internalClient, err = internal.NewInternalClient(config, nil, &cache.Cache{TTL: 1}, &ratelimit.RateLimit{Enabled: true})
	require.NoError(t, err)

	require.True(t, internalClient.IsCacheEnabled)
	require.True(t, internalClient.IsRateLimitEnabled)
	require.True(t, internalClient.IsRetryEnabled)
}

func TestStatusCodeToError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []int{400, 401, 403, 404, 405, 415, 418, 429, 500, 502, 503, 504}

	internal := util.NewTestInternalClient(t)

	logger := internal.Logger("client_endpoint_method")
	ctx := context.Background()
	urlComponents := []string{"https://", "tests", api.RIOT_API_BASE_URL_FORMAT, "/"}
	equinoxReq, err := internal.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(fmt.Sprint(test), func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://tests.api.riotgames.com/",
				httpmock.NewStringResponder(test, `"response"`).Times(2))

			wantErr := api.StatusCodeToError(test)

			var data string
			err = internal.Execute(ctx, equinoxReq, data)

			if wantErr == nil && test == 418 {
				require.EqualError(t, err, "unexpected status code: 418")
			} else {
				require.Equal(t, wantErr, err)
			}

			_, err = internal.ExecuteBytes(ctx, equinoxReq)
			if wantErr == nil && test == 418 {
				require.EqualError(t, err, "unexpected status code: 418")
			} else {
				require.Equal(t, wantErr, err)
			}
		})
	}
}

func TestRequests(t *testing.T) {
	config := util.NewTestEquinoxConfig()
	client := util.NewTestInternalClient(t)

	expectedURL := "https://cool.and.real.api/post"

	ctx := context.Background()
	logger := client.Logger("client_endpoint_method")
	urlComponents := []string{"https://", "", "cool.and.real.api", "/post"}

	t.Run("Request with body", func(t *testing.T) {
		body := map[string]any{
			"message": "cool",
		}

		expectedBody, err := jsonv2.Marshal(body)
		require.NoError(t, err)

		equinoxReq, err := client.Request(ctx, logger, "POST", urlComponents, "", body)
		require.NoError(t, err)

		require.Equal(t, expectedURL, equinoxReq.Request.URL.String())
		bodyBytes, err := io.ReadAll(equinoxReq.Request.Body)
		require.NoError(t, err)
		require.Equal(t, expectedBody, bodyBytes)
		require.Equal(t, "application/json", equinoxReq.Request.Header.Get("Content-Type"))
		require.Equal(t, "application/json", equinoxReq.Request.Header.Get("Accept"))
		require.Equal(t, config.Key, equinoxReq.Request.Header.Get("X-Riot-Token"))
	})

	t.Run("Request with invalid body", func(t *testing.T) {
		// Errors usually don't have any exported fields to be serialized
		body := []struct {
			Error error
		}{
			{Error: fmt.Errorf("invalid body")},
		}

		_, err := jsonv2.Marshal(body)
		require.Error(t, err)

		equinoxReq, err := client.Request(ctx, logger, "POST", urlComponents, "", body)
		require.Error(t, err)

		require.Nil(t, equinoxReq.Request)
	})

	t.Run("Request without body", func(t *testing.T) {
		equinoxReq, err := client.Request(ctx, logger, "POST", urlComponents, "", nil)
		require.NoError(t, err)

		require.Equal(t, expectedURL, equinoxReq.Request.URL.String())
		require.Nil(t, equinoxReq.Request.Body)
		require.Equal(t, "application/json", equinoxReq.Request.Header.Get("Content-Type"))
		require.Equal(t, "application/json", equinoxReq.Request.Header.Get("Accept"))
		require.Equal(t, config.Key, equinoxReq.Request.Header.Get("X-Riot-Token"))
	})

	t.Run("Request with invalid url", func(t *testing.T) {
		wantErr := &url.Error{
			Op:  "parse",
			URL: "https://----.api.riotgames.com\\:invalid:/=",
			Err: url.InvalidHostError("\\"),
		}

		_, err := client.Request(ctx, logger, http.MethodGet, []string{"https://", "----", api.RIOT_API_BASE_URL_FORMAT, "\\:invalid:/="}, "", nil)
		require.Equal(t, wantErr, err)
	})
}

func TestExecutes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	internalClient := util.NewTestInternalClient(t)

	// Valid responses

	ctx := context.Background()
	logger := internalClient.Logger("client_endpoint_method")
	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}
	equinoxReq, err := internalClient.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewStringResponder(200, `"response"`).Times(2))

	var target string
	err = internalClient.Execute(ctx, equinoxReq, &target)
	require.NoError(t, err)
	require.Equal(t, `response`, target)

	data, err := internalClient.ExecuteBytes(ctx, equinoxReq)
	require.NoError(t, err)
	require.Equal(t, []byte(`"response"`), data)

	// Nil context

	//lint:ignore SA1012 Testing if ctx is nil
	//nolint:staticcheck
	equinoxReq, err = internalClient.Request(nil, logger, http.MethodGet, urlComponents, "", nil)
	require.Error(t, err)
	//lint:ignore SA1012 Testing if ctx is nil
	//nolint:staticcheck
	err = internalClient.Execute(nil, equinoxReq, &target)
	require.Error(t, err)
	//lint:ignore SA1012 Testing if ctx is nil
	//nolint:staticcheck
	_, err = internalClient.ExecuteBytes(nil, equinoxReq)
	require.Error(t, err)

	// Response has invalid json in Execute

	equinoxReq, err = internalClient.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewStringResponder(200, `-{invalid json}-`))

	target = ""
	err = internalClient.Execute(ctx, equinoxReq, &target)
	require.Error(t, err)

	// Body is not marshalled in ExecuteBytes so it wont error

	data, err = internalClient.ExecuteBytes(ctx, equinoxReq)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestExecutesWithCache(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewStringResponder(200, `"response"`))

	httpmock.RegisterResponder("POST", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewStringResponder(200, `"response2"`))

	ctx := context.Background()

	config := util.NewTestEquinoxConfig()
	cache, err := equinox.DefaultCache()
	require.NoError(t, err)
	internalClient, err := internal.NewInternalClient(config, nil, cache, nil)
	require.NoError(t, err)
	require.True(t, internalClient.IsCacheEnabled)

	logger := internalClient.Logger("client_endpoint_method")
	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}

	// Get

	equinoxReq, err := internalClient.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	var res string
	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response`, res)

	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response`, res)

	// Post or any other method

	equinoxReq, err = internalClient.Request(ctx, logger, http.MethodPost, urlComponents, "", nil)
	require.NoError(t, err)
	require.Equal(t, "POST", equinoxReq.Request.Method)

	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response2`, res)
}

func TestRateLimitRetry(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := util.NewTestEquinoxConfig()
	config.Retry = api.Retry{MaxRetries: 1, Jitter: 500 * time.Millisecond}
	internalClient, err := internal.NewInternalClient(config, nil, nil, ratelimit.NewInternalRateLimit(0.99, time.Second))
	require.NoError(t, err)

	require.True(t, internalClient.IsRetryEnabled)

	// All counts start at 1 here because only CheckRetryAfter and client.Do is being tested

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewBytesResponder(429, []byte(`{}`)).HeaderSet(map[string][]string{
			"X-Rate-Limit-Type":         {"application"},
			"X-App-Rate-Limit":          {"100:100"},
			"X-App-Rate-Limit-Count":    {"1:100"},
			"X-Method-Rate-Limit":       {"100:100"},
			"X-Method-Rate-Limit-Count": {"1:100"},
			"Retry-After":               {"1"},
		}).Times(2))

	ctx := context.Background()
	logger := internalClient.Logger("client_endpoint_method")
	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}

	var res lol.StatusPlatformDataV4DTO

	equinoxReq, err := internalClient.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	// Application rate limited
	// This will take 2.5 seconds since it will retry one time. Retry-After + 0.5s of jitter
	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.ErrorIs(t, err, internal.ErrMaxRetries)

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewBytesResponder(429, []byte(`{}`)).HeaderSet(map[string][]string{
			"X-Rate-Limit-Type":         {"method"},
			"X-App-Rate-Limit":          {"100:100"},
			"X-App-Rate-Limit-Count":    {"1:100"},
			"X-Method-Rate-Limit":       {"100:100"},
			"X-Method-Rate-Limit-Count": {"1:100"},
			"Retry-After":               {"1"},
		}).Times(2))

	// Method/Service rate limited
	// This will take 2.5 seconds since it will retry one time. Retry-After + 0.5s of jitter
	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.ErrorIs(t, err, internal.ErrMaxRetries)

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewBytesResponder(429, []byte(`{}`)).HeaderSet(map[string][]string{
			"X-App-Rate-Limit":          {"100:100"},
			"X-App-Rate-Limit-Count":    {"1:100"},
			"X-Method-Rate-Limit":       {"100:100"},
			"X-Method-Rate-Limit-Count": {"1:100"},
		}).Times(1))

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewBytesResponder(200, []byte(`{}`)))

	// Rate limited but deadline exceeds

	ctxWithDeadline, c := context.WithDeadline(ctx, time.Now())
	defer c()

	_, err = internalClient.ExecuteBytes(ctxWithDeadline, equinoxReq)
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	err = internalClient.Execute(ctxWithDeadline, equinoxReq, &res)
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	// Method/Service rate limited
	// This will take 2.5 seconds since it will retry one time and then succeed. DEFAULT_RETRY_AFTER + 0.5s of jitter
	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)

	// Wont block
	_, err = internalClient.ExecuteBytes(ctx, equinoxReq)
	require.NoError(t, err)
}

func TestExponentialBackoffRetry(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewStringResponder(429, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Retry.MaxRetries = 2
	config.Retry.Jitter = 200 * time.Millisecond

	internalClient := util.NewTestCustomInternalClient(t, config)
	require.True(t, internalClient.IsRetryEnabled)

	ctx := context.Background()
	logger := internalClient.Logger("client_endpoint_method")
	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}

	equinoxReq, err := internalClient.Request(ctx, logger, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	// Should error out and take around 6.4 seconds
	_, err = internalClient.ExecuteBytes(ctx, equinoxReq)
	require.Error(t, err)
}
