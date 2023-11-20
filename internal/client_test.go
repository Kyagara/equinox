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
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestNewInternalClient(t *testing.T) {
	_, err := internal.NewInternalClient(nil)
	require.NotEmpty(t, err, "expecting non-nil error")

	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	if err != nil {
		require.ErrorContains(t, err, "error initializing logger")
	}

	require.NotEmpty(t, internalClient, "expecting non-nil InternalClient")

	config := equinox.NewTestEquinoxConfig()
	config.Cache.TTL = 1

	internalClient, err = internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, internalClient, "expecting non-nil InternalClient")
}

func TestInternalClientNewRequest(t *testing.T) {
	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
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

func TestInternalClientRequest(t *testing.T) {
	gock.New("https://cool.and.real.api").
		Post("/post").
		Reply(200)

	body := map[string]any{
		"message": "cool",
	}
	config := equinox.NewTestEquinoxConfig()
	client, err := internal.NewInternalClient(config)
	require.Nil(t, err, "expecting nil error")

	url := "https://cool.and.real.api/post"

	t.Run("Request with body", func(t *testing.T) {
		expectedBody, err := json.Marshal(body)
		require.Nil(t, err, "expecting nil error")
		req, err := client.Request("https://cool.and.real.api%v", "POST", "", "/post", body)
		require.Nil(t, err, "expecting nil error")

		if req.URL.String() != url {
			t.Errorf("unexpected URL, got %s, want %s", req.URL.String(), url)
		}

		bodyBytes, err := io.ReadAll(req.Body)
		require.Nil(t, err, "expecting nil error")

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
		req, err := client.Request("https://cool.and.real.api%v", "POST", "", "/post", nil)
		require.Nil(t, err, "expecting nil error")

		if req.URL.String() != url {
			t.Errorf("unexpected URL, got %s, want %s", req.URL.String(), url)
		}
		if req.Body != nil {
			t.Errorf("unexpected body, got %v, want nil", req.Body)
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
}

func TestInternalClientErrorResponses(t *testing.T) {
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
			name:    "rate limited",
			wantErr: api.ErrTooManyRequests,
			code:    429,
		},
		{
			name:    "rate limited but no retry-after header found",
			wantErr: api.ErrRetryAfterHeaderNotFound,
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
			name: "unknown error",
			wantErr: api.ErrorResponse{
				Status: api.Status{
					Message:    "Unknown error",
					StatusCode: 418,
				},
			},
			code: 418,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := equinox.NewTestEquinoxConfig()
			if test.name == "rate limited" {
				config.Retry = 0
			} else if test.name == "rate limited but no retry-after header found" {
				config.Retry = 1
			}

			gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, "tests")).
				Get("/").
				Reply(test.code)

			internalClient, err := internal.NewInternalClient(config)
			require.Nil(t, err, "expecting nil error")
			l := internalClient.Logger("client_endpoint_method")
			request, err := internalClient.Request(api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "tests", "/", nil)
			require.Nil(t, err, "expecting nil error")
			var gotData api.PlainTextResponse
			gotErr := internalClient.Execute(l, request, gotData)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

// Testing if InternalClient.Post() can properly decode a plain text response.
func TestInternalClientPlainTextResponse(t *testing.T) {
	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
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
	require.NotEmpty(t, object, "expecting non-nil response")
}

func TestInternalClientRetries(t *testing.T) {
	config := equinox.NewTestEquinoxConfig()
	config.Retry = 1
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

func TestGetDDragonLOLVersions(t *testing.T) {
	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "/api/versions.json")).
		Get("").
		Reply(200).
		JSON("[\"1.0\"]")

	versions, err := internalClient.GetDDragonLOLVersions("client_endpoint_method")
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, "1.0", versions[0], "expecting nil error")
}
