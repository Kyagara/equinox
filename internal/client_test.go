package internal_test

import (
	"context"
	"time"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/allegro/bigcache/v3"
	jsonv2 "github.com/go-json-experiment/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestNewInternalClient(t *testing.T) {
	t.Parallel()
	internalClient := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.NotEmpty(t, internalClient)
	require.False(t, internalClient.IsCacheEnabled)
	require.False(t, internalClient.IsRateLimitEnabled)
	require.False(t, internalClient.IsRetryEnabled)

	config := util.NewTestEquinoxConfig()
	config.HTTPClient = nil
	config.Cache = nil
	config.RateLimit = nil

	internalClient = internal.NewInternalClient(config)
	require.NotEmpty(t, internalClient)
	require.False(t, internalClient.IsCacheEnabled)
	require.False(t, internalClient.IsRateLimitEnabled)
	require.False(t, internalClient.IsRetryEnabled)

	config = util.NewTestEquinoxConfig()
	config.Cache.TTL = 1
	config.Retry.MaxRetries = 1
	config.RateLimit.Enabled = true
	config.Key = ""

	internalClient = internal.NewInternalClient(config)
	require.NotEmpty(t, internalClient)
	require.True(t, internalClient.IsCacheEnabled)
	require.True(t, internalClient.IsRateLimitEnabled)
	require.True(t, internalClient.IsRetryEnabled)

	l := internalClient.Logger("client_endpoint_method")
	ctx := context.Background()
	_, gotErr := internalClient.Request(ctx, l, http.MethodGet, []string{"https://", "", api.RIOT_API_BASE_URL_FORMAT}, "", nil)
	require.Error(t, gotErr)

	_, gotErr = internalClient.Request(ctx, l, http.MethodGet, []string{"https://", "", api.D_DRAGON_BASE_URL_FORMAT}, "", nil)
	require.NoError(t, gotErr)
}

func TestInternalClientNewRequest(t *testing.T) {
	internal := internal.NewInternalClient(util.NewTestEquinoxConfig())
	l := internal.Logger("client_endpoint_method")
	tests := []struct {
		name    string
		want    *http.Request
		wantErr error
		method  string
		url     string
	}{
		{
			name: "invalid url",
			wantErr: &url.Error{
				Op:  "parse",
				URL: "https://----.api.riotgames.com\\:invalid:/=",
				Err: url.InvalidHostError("\\"),
			},
			url: "\\:invalid:/=",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			_, gotErr := internal.Request(ctx, l, http.MethodGet, []string{"https://", "----", api.RIOT_API_BASE_URL_FORMAT, test.url}, "", nil)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

func TestInternalClientRequest(t *testing.T) {
	body := map[string]any{
		"message": "cool",
	}
	config := util.NewTestEquinoxConfig()
	client := internal.NewInternalClient(config)

	url := "https://cool.and.real.api/post"

	t.Run("Request with body", func(t *testing.T) {
		expectedBody, err := jsonv2.Marshal(body)
		require.NoError(t, err)
		logger := client.Logger("client_endpoint_method")
		ctx := context.Background()
		urlComponents := []string{"https://", "", "cool.and.real.api", "/post"}
		equinoxReq, err := client.Request(ctx, logger, "POST", urlComponents, "", body)
		require.NoError(t, err)

		if equinoxReq.Request.URL.String() != url {
			t.Errorf("unexpected URL, got %s, want %s", equinoxReq.Request.URL.String(), url)
		}

		bodyBytes, err := io.ReadAll(equinoxReq.Request.Body)
		require.NoError(t, err)

		if string(bodyBytes) != string(expectedBody) {
			t.Errorf("unexpected body, got %s, want %s", string(bodyBytes), string(expectedBody))
		}
		if equinoxReq.Request.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected Content-Type header, got %s, want application/json", equinoxReq.Request.Header.Get("Content-Type"))
		}
		if equinoxReq.Request.Header.Get("Accept") != "application/json" {
			t.Errorf("unexpected Accept header, got %s, want application/json", equinoxReq.Request.Header.Get("Accept"))
		}
		if equinoxReq.Request.Header.Get("X-Riot-Token") != config.Key {
			t.Errorf("unexpected X-Riot-Token header, got %s, want %s", equinoxReq.Request.Header.Get("X-Riot-Token"), config.Key)
		}
		if equinoxReq.Request.Header.Get("User-Agent") != "equinox - https://github.com/Kyagara/equinox" {
			t.Errorf("unexpected User-Agent header, got %s, want equinox - https://github.com/Kyagara/equinox", equinoxReq.Request.Header.Get("User-Agent"))
		}
	})

	t.Run("Request without body", func(t *testing.T) {
		logger := client.Logger("client_endpoint_method")
		ctx := context.Background()
		urlComponents := []string{"https://", "", "cool.and.real.api", "/post"}
		equinoxReq, err := client.Request(ctx, logger, "POST", urlComponents, "", nil)
		require.NoError(t, err)

		if equinoxReq.Request.URL.String() != url {
			t.Errorf("unexpected URL, got %s, want %s", equinoxReq.Request.URL.String(), url)
		}
		if equinoxReq.Request.Body != nil {
			t.Errorf("unexpected body, got %v, want nil", equinoxReq.Request.Body)
		}
		if equinoxReq.Request.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected Content-Type header, got %s, want application/json", equinoxReq.Request.Header.Get("Content-Type"))
		}
		if equinoxReq.Request.Header.Get("Accept") != "application/json" {
			t.Errorf("unexpected Accept header, got %s, want application/json", equinoxReq.Request.Header.Get("Accept"))
		}
		if equinoxReq.Request.Header.Get("X-Riot-Token") != config.Key {
			t.Errorf("unexpected X-Riot-Token header, got %s, want %s", equinoxReq.Request.Header.Get("X-Riot-Token"), config.Key)
		}
		if equinoxReq.Request.Header.Get("User-Agent") != "equinox - https://github.com/Kyagara/equinox" {
			t.Errorf("unexpected User-Agent header, got %s, want equinox - https://github.com/Kyagara/equinox", equinoxReq.Request.Header.Get("User-Agent"))
		}
	})
}

func TestInternalClientErrorResponses(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name    string
		wantErr error
		code    int
	}{
		{
			name:    "bad request",
			wantErr: api.ErrBadRequest,
			code:    400,
		},
		{
			name:    "unauthorized",
			wantErr: api.ErrUnauthorized,
			code:    401,
		},
		{
			name:    "forbidden",
			wantErr: api.ErrForbidden,
			code:    403,
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			code:    404,
		},
		{
			name:    "method not allowed",
			wantErr: api.ErrMethodNotAllowed,
			code:    405,
		},
		{
			name:    "unsupported media type",
			wantErr: api.ErrUnsupportedMediaType,
			code:    415,
		},
		{
			name:    "too many requests",
			wantErr: api.ErrTooManyRequests,
			code:    429,
		},
		{
			name:    "internal server error",
			wantErr: api.ErrInternalServer,
			code:    500,
		},
		{
			name:    "bad gateway",
			wantErr: api.ErrBadGateway,
			code:    502,
		},
		{
			name:    "service unavailable",
			wantErr: api.ErrServiceUnavailable,
			code:    503,
		},
		{
			name:    "gateway timeout",
			wantErr: api.ErrGatewayTimeout,
			code:    504,
		},
		{
			name:    "unknown error",
			wantErr: fmt.Errorf("unexpected status code: %d", 418),
			code:    418,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://tests.api.riotgames.com/",
				httpmock.NewBytesResponder(test.code, []byte(`{}`)))

			internal := internal.NewInternalClient(util.NewTestEquinoxConfig())
			l := internal.Logger("client_endpoint_method")
			ctx := context.Background()
			urlComponents := []string{"https://", "tests", api.RIOT_API_BASE_URL_FORMAT, "/"}
			equinoxReq, err := internal.Request(ctx, l, http.MethodGet, urlComponents, "", nil)
			require.NoError(t, err)
			require.Equal(t, "GET", equinoxReq.Request.Method)
			var data interface{}
			err = internal.Execute(ctx, equinoxReq, data)
			require.Equal(t, test.wantErr, err)
			_, err = internal.ExecuteRaw(ctx, equinoxReq)
			require.Equal(t, test.wantErr, err)
		})
	}
}

func TestInternalClientRetryableErrors(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := util.NewTestEquinoxConfig()
	config.Retry = api.Retry{MaxRetries: 1, Jitter: 500 * time.Millisecond}
	config.RateLimit = ratelimit.NewInternalRateLimit(0.99, 1*time.Second)
	i := internal.NewInternalClient(config)
	require.True(t, i.IsRetryEnabled)

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

	res := lol.PlatformDataV4DTO{}
	l := i.Logger("client_endpoint_method")

	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}

	//lint:ignore SA1012 Testing if ctx is nil
	//nolint:staticcheck
	equinoxReq, err := i.Request(nil, l, http.MethodGet, urlComponents, "", nil)
	require.Error(t, err)
	//lint:ignore SA1012 Testing if ctx is nil
	//nolint:staticcheck
	err = i.Execute(nil, equinoxReq, &res)
	require.Error(t, err)
	//lint:ignore SA1012 Testing if ctx is nil
	//nolint:staticcheck
	_, err = i.ExecuteRaw(nil, equinoxReq)
	require.Error(t, err)

	ctx := context.Background()
	equinoxReq, err = i.Request(ctx, l, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	// Application rate limited
	// This will take 2.5 seconds since it will retry one time. Retry-After + 0.5s of jitter
	err = i.Execute(ctx, equinoxReq, &res)
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
	err = i.Execute(ctx, equinoxReq, &res)
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

	// Method/Service rate limited
	// This will take 2.5 seconds since it will retry one time and then succeed. DEFAULT_RETRY_AFTER + 0.5s of jitter
	err = i.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
}

func TestGetDDragonLOLVersions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	internal := internal.NewInternalClient(util.NewTestEquinoxConfig())

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/api/versions.json",
		httpmock.NewStringResponder(200, `["1.0"]`).Times(1))

	ctx := context.Background()
	versions, err := internal.GetDDragonLOLVersions(ctx, "client_endpoint_method")
	require.NoError(t, err)
	require.Equal(t, "1.0", versions[0])

	versions, err = internal.GetDDragonLOLVersions(ctx, "client_endpoint_method")
	require.Error(t, err)
	require.Empty(t, versions)

	httpmock.RegisterResponder("GET", "https://ddragon.leagueoflegends.com/api/versions.json",
		httpmock.NewStringResponder(404, `["1.0"]`).Times(1))

	versions, err = internal.GetDDragonLOLVersions(ctx, "client_endpoint_method")
	require.Error(t, err)
	require.Empty(t, versions)
}

func TestGetURLWithAuthorizationHash(t *testing.T) {
	t.Parallel()
	req := &http.Request{
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
		Header: http.Header{},
	}

	equinoxReq := api.EquinoxRequest{Request: req}
	equinoxReq.URL = req.URL.String()

	hash, err := internal.GetURLWithAuthorizationHash(equinoxReq)
	require.NoError(t, err)
	require.Equal(t, "http://example.com/path", hash)

	req.Header.Set("Authorization", "7267ee00-5696-47b8-9cae-8db3d49c8c33")
	hash, err = internal.GetURLWithAuthorizationHash(equinoxReq)
	require.NoError(t, err)
	require.Equal(t, "http://example.com/path-45da11db1ebd17ee0c32aca62e08923ea4f15590058ff1e15661bc13ed33df9d", hash)
}

func TestInternalClientExecutes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/status/v4/platform-data",
		httpmock.NewStringResponder(200, `"response"`))

	internal := internal.NewInternalClient(util.NewTestEquinoxConfig())

	ctx := context.Background()

	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}
	equinoxReq, err := internal.Request(ctx, internal.Logger("client_endpoint_method"), http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)

	data, err := internal.ExecuteRaw(ctx, equinoxReq)
	require.NoError(t, err)
	require.Equal(t, []byte(`"response"`), data)

	var res string
	err = internal.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response`, res)
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
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(5*time.Minute))
	require.NoError(t, err)
	config.Cache = cache
	internalClient := internal.NewInternalClient(config)
	require.NotEmpty(t, internalClient)
	require.True(t, internalClient.IsCacheEnabled)

	l := internalClient.Logger("client_endpoint_method")
	urlComponents := []string{"https://", lol.BR1.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}
	equinoxReq, err := internalClient.Request(ctx, l, http.MethodGet, urlComponents, "", nil)
	require.NoError(t, err)
	require.Equal(t, "GET", equinoxReq.Request.Method)

	var res string
	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response`, res)

	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response`, res)

	equinoxReq, err = internalClient.Request(ctx, l, http.MethodPost, urlComponents, "", nil)
	require.NoError(t, err)
	require.Equal(t, "POST", equinoxReq.Request.Method)

	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response2`, res)
	err = internalClient.Execute(ctx, equinoxReq, &res)
	require.NoError(t, err)
	require.Equal(t, `response2`, res)
}
