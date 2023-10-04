package internal_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestInternalClient(t *testing.T) {
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

func TestInternalClientPut(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Put("/").
		Reply(200)

	err = internalClient.Put("tests", "/", nil, "", "")

	require.Nil(t, err, "expecting nil error")
}

func TestInternalClientDataDragonGet(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "/")).
		Get("").
		Reply(200).
		JSON(&data_dragon.DataDragonMetadata{})

	target := &data_dragon.DataDragonMetadata{}

	err = internalClient.DataDragonGet("/", target, "endpoint", "method")

	require.Nil(t, err, "expecting nil error")
}

func TestInternalClientRetries(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.Retry = true

	internalClient, err := internal.NewInternalClient(config)

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(lol.StatusURL).
		Reply(429).SetHeader("Retry-After", "1").
		JSON(&api.PlatformDataDTO{})

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(lol.StatusURL).
		Reply(200).
		JSON(&api.PlatformDataDTO{})

	res := api.PlatformDataDTO{}

	// This will take 1 second
	err = internalClient.Get(lol.BR1, lol.StatusURL, &res, "", "", "")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, res, "expecting non-nil response")
}

func TestInternalClientFailingRetry(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.Retry = true

	internalClient, err := internal.NewInternalClient(config)

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	var object api.PlainTextResponse

	// This will take 2 seconds
	gotErr := internalClient.Get("tests", "/", &object, "", "", "")

	require.Equal(t, api.ErrTooManyRequests, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrTooManyRequests, gotErr))
}

func TestInternalClientRetryHeaderNotFound(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.Retry = true

	internalClient, err := internal.NewInternalClient(config)

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429)

	var object api.PlainTextResponse

	gotErr := internalClient.Get("tests", "/", &object, "", "", "")

	require.Equal(t, api.ErrRetryAfterHeaderNotFound, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrRetryAfterHeaderNotFound, gotErr))
}

// Testing if InternalClient.Post() can properly decode a plain text response.
func TestInternalClientPlainTextResponse(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Post("/").
		Reply(200).BodyString("response")

	var object api.PlainTextResponse

	err = internalClient.Post("tests", "/", nil, &object, "", "", "")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, object, "expecting non-nil response")
}

func TestInternalClientPost(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Post("/").
		Reply(200).JSON(&api.PlatformDataDTO{})

	var object *api.PlatformDataDTO

	err = internalClient.Post("tests", "/", nil, &object, "", "", "")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, object, "expecting non-nil response")
}

// Testing if the InternalClient can properly handle a status code not specified in the Riot api.
func TestInternalClientHandleErrorResponse(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(418).BodyString("response")

	var object api.PlainTextResponse

	gotErr := internalClient.Get("tests", "/", &object, "", "", "")

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
			gotErr := internalClient.Get("----", test.url, nil, "", "", "")

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
				Get("/").
				Reply(test.code)

			internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

			require.Nil(t, err, "expecting nil error")

			var gotData api.PlainTextResponse

			gotErr := internalClient.Get(test.region, "/", &gotData, "", "", "")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

func TestInternalClientRateLimit(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	internalClient, err := internal.NewInternalClient(config)

	require.Nil(t, err, "expecting nil error")

	headers := map[string]string{}

	// The method will be rate limited
	headers[api.RateLimitTypeHeader] = "method"

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Put("/").
		Reply(429).SetHeaders(headers)

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Put("/").
		Reply(200)

	err = internalClient.Put("tests", "/", nil, "", "")

	require.Equal(t, api.ErrTooManyRequests, err, fmt.Sprintf("want err %v, got %v", api.ErrTooManyRequests, err))

	err = internalClient.Put("tests", "/", nil, "", "")

	require.Nil(t, err, "expecting nil error")
}

func TestCacheIsDisabled(t *testing.T) {
	client, err := equinox.NewClientWithConfig(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	err = client.Cache.Clear()

	require.Equal(t, cache.ErrCacheIsDisabled, err, fmt.Sprintf("want err %v, got %v", cache.ErrCacheIsDisabled, err))
}
