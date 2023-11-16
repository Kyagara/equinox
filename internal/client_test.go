package internal_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestInternalClient(t *testing.T) {
	_, err := internal.NewInternalClient(nil)
	require.NotNil(t, err, "expecting non-nil error")

	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	if err != nil {
		require.ErrorContains(t, err, "error initializing logger")
	}

	require.NotNil(t, internalClient, "expecting non-nil InternalClient")
	require.False(t, internalClient.IsCacheEnabled)
	require.False(t, internalClient.IsRetryEnabled)

	config := internal.NewTestEquinoxConfig()
	config.Cache.TTL = 1
	config.Retry = true

	internalClient, err = internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, internalClient, "expecting non-nil InternalClient")
	require.True(t, internalClient.IsCacheEnabled)
	require.True(t, internalClient.IsRetryEnabled)
}

func TestInternalClientDDragonGet(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "/")).
		Get("").
		Reply(200).
		JSON(&ddragon.Champion{})

	target := &ddragon.Champion{}
	request, err := internalClient.Request(api.D_DRAGON_BASE_URL_FORMAT, http.MethodGet, "", "", nil)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	err = internalClient.Execute(l, request, target)
	require.Nil(t, err, "expecting nil error")
}

func TestGetDDragonLOLVersions(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "/api/versions.json")).
		Get("").
		Reply(200).
		JSON("[\"1.0\"]")

	versions, err := internalClient.GetDDragonLOLVersions("client_endpoint_method")
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, "1.0", versions[0], "expecting nil error")
}

func TestInternalClientRetries(t *testing.T) {
	config := internal.NewTestEquinoxConfig()
	config.Retry = true
	internalClient, err := internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/status/v4/platform-data").
		Reply(429).SetHeader("Retry-After", "1").
		JSON(&lol.PlatformDataV4DTO{})

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/status/v4/platform-data").
		Reply(200).
		JSON(&lol.PlatformDataV4DTO{})

	res := lol.PlatformDataV4DTO{}

	// This will take 1 second
	request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, lol.BR1, "/lol/status/v4/platform-data", nil)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	err = internalClient.Execute(l, request, &res)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, res, "expecting non-nil response")
}

func TestInternalClientFailingRetry(t *testing.T) {
	config := internal.NewTestEquinoxConfig()
	config.Retry = true
	internalClient, err := internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	var object api.PlainTextResponse
	// This will take 2 seconds
	request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "tests", "/", nil)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	gotErr := internalClient.Execute(l, request, &object)
	require.Equal(t, api.ErrTooManyRequests, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrTooManyRequests, gotErr))
}

func TestInternalClientRetryHeaderNotFound(t *testing.T) {
	config := internal.NewTestEquinoxConfig()
	config.Retry = true
	internalClient, err := internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Get("/").
		Reply(429)

	var object api.PlainTextResponse
	request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "tests", "/", nil)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	gotErr := internalClient.Execute(l, request, &object)
	require.Equal(t, api.ErrRetryAfterHeaderNotFound, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrRetryAfterHeaderNotFound, gotErr))
}

// Testing if InternalClient.Post() can properly decode a plain text response.
func TestInternalClientPlainTextResponse(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Post("/").
		Reply(200).BodyString("response")

	var object api.PlainTextResponse
	request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodPost, "tests", "/", nil)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	err = internalClient.Execute(l, request, &object)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, object, "expecting non-nil response")
}

// Testing if the InternalClient can properly handle a status code not specified in the Riot api.
func TestInternalClientHandleErrorResponse(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Get("/").
		Reply(418).BodyString("response")

	var object api.PlainTextResponse
	request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "tests", "/", nil)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	gotErr := internalClient.Execute(l, request, &object)
	wantErr := api.ErrorResponse{
		Status: api.Status{
			Message:    "Unknown error",
			StatusCode: 418,
		},
	}
	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

func TestInternalClientNewRequest(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
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
			_, gotErr := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "----", test.url, nil)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

func TestInternalClientErrorResponses(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
		code    int
		region  string
	}{
		{
			name:    "bad request",
			wantErr: api.ErrBadRequest,
			code:    400,
			region:  "tests",
		},
		{
			name:    "unauthorized",
			wantErr: api.ErrUnauthorized,
			code:    401,
			region:  "tests",
		},
		{
			name:    "forbidden",
			wantErr: api.ErrForbidden,
			code:    403,
			region:  "tests",
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			code:    404,
			region:  "tests",
		},
		{
			name:    "method not allowed",
			wantErr: api.ErrMethodNotAllowed,
			code:    405,
			region:  "tests",
		},
		{
			name:    "unsupported media type",
			wantErr: api.ErrUnsupportedMediaType,
			code:    415,
			region:  "tests",
		},
		{
			name:    "rate limited with retry disabled",
			wantErr: api.ErrTooManyRequests,
			code:    429,
			region:  "tests",
		},
		{
			name:    "internal server error",
			wantErr: api.ErrInternalServer,
			code:    500,
			region:  "tests",
		},
		{
			name:    "bad gateway",
			wantErr: api.ErrBadGateway,
			code:    502,
			region:  "tests",
		},
		{
			name:    "service unavailable",
			wantErr: api.ErrServiceUnavailable,
			code:    503,
			region:  "tests",
		},
		{
			name:    "gateway timeout",
			wantErr: api.ErrGatewayTimeout,
			code:    504,
			region:  "tests",
		},
	}

	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
				Get("/").
				Reply(test.code)

			request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, test.region, "/", nil)
			require.Nil(t, err, "expecting nil error")
			var gotData api.PlainTextResponse
			gotErr := internalClient.Execute(l, request, gotData)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

func TestInternalClientRateLimit(t *testing.T) {
	config := internal.NewTestEquinoxConfig()
	internalClient, err := internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")
	l := internalClient.Logger("client_endpoint_method")
	headers := map[string]string{}
	// The method will be rate limited
	headers[api.X_RATE_LIMIT_TYPE_HEADER] = "method"

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Put("/").
		Reply(429).SetHeaders(headers)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
		Put("/").
		Reply(200)

	request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodPut, "tests", "/", nil)
	require.Nil(t, err, "expecting nil error")
	err = internalClient.Execute(l, request, nil)
	require.Equal(t, api.ErrTooManyRequests, err, fmt.Sprintf("want err %v, got %v", api.ErrTooManyRequests, err))
	err = internalClient.Execute(l, request, nil)
	require.Nil(t, err, "expecting nil error")
}

func TestCacheIsDisabled(t *testing.T) {
	client, err := equinox.NewClientWithConfig(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	err = client.Cache.Clear()
	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", cache.ErrCacheIsDisabled, err))
}

func TestInternalClientRequest(t *testing.T) {
	base := "https://cool.and.real.api/api/%v"
	route := "users"
	method := "POST"
	url := "/post"
	body := map[string]any{
		"message": "cool",
	}
	config := internal.NewTestEquinoxConfig()
	config.Key = "RGAPI-TEST"
	c, err := internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")
	t.Run("Request with body", func(t *testing.T) {
		expectedURL := fmt.Sprintf(base, route) + url
		expectedBody, _ := json.Marshal(body)
		req, err := c.Request(base, method, route, url, body)
		require.Nil(t, err, "expecting nil error")
		if req.URL.String() != expectedURL {
			t.Errorf("unexpected URL, got %s, want %s", req.URL.String(), expectedURL)
		}

		bodyBytes, _ := io.ReadAll(req.Body)
		if string(bodyBytes) != string(expectedBody) {
			t.Errorf("unexpected body, got %s, want %s", string(bodyBytes), string(expectedBody))
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected Content-Type header, got %s, want application/json", req.Header.Get("Content-Type"))
		}
		if req.Header.Get("Accept") != "application/json" {
			t.Errorf("unexpected Accept header, got %s, want application/json", req.Header.Get("Accept"))
		}
		if req.Header.Get("X-Riot-Token") != config.Key {
			t.Errorf("unexpected X-Riot-Token header, got %s, want %s", req.Header.Get("X-Riot-Token"), config.Key)
		}
		if req.Header.Get("User-Agent") != "equinox - https://github.com/Kyagara/equinox" {
			t.Errorf("unexpected User-Agent header, got %s, want equinox - https://github.com/Kyagara/equinox", req.Header.Get("User-Agent"))
		}
	})

	t.Run("Request without body", func(t *testing.T) {
		expectedURL := fmt.Sprintf(base, route) + url
		req, err := c.Request(base, method, route, url, nil)
		require.Nil(t, err, "expecting nil error")

		if req.URL.String() != expectedURL {
			t.Errorf("unexpected URL, got %s, want %s", req.URL.String(), expectedURL)
		}
		if req.Body != nil {
			t.Errorf("unexpected body, got %v, want nil", req.Body)
		}
		if req.Header.Get("Content-Type") != "" {
			t.Errorf("unexpected Content-Type header, got %s, want empty", req.Header.Get("Content-Type"))
		}
		if req.Header.Get("Accept") != "application/json" {
			t.Errorf("unexpected Accept header, got %s, want application/json", req.Header.Get("Accept"))
		}
		if req.Header.Get("X-Riot-Token") != config.Key {
			t.Errorf("unexpected X-Riot-Token header, got %s, want %s", req.Header.Get("X-Riot-Token"), config.Key)
		}
		if req.Header.Get("User-Agent") != "equinox - https://github.com/Kyagara/equinox" {
			t.Errorf("unexpected User-Agent header, got %s, want equinox - https://github.com/Kyagara/equinox", req.Header.Get("User-Agent"))
		}
	})
}
