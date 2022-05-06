package internal_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestInternalClient(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	assert.NotNil(t, client, "expecting non-nil InternalClient")
}

func TestInternalClientRetries(t *testing.T) {
	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(lol.StatusURL).
		Reply(429).SetHeader("Retry-After", "1").
		JSON(&api.PlatformDataDTO{})

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(lol.StatusURL).
		Reply(200).
		JSON(&api.PlatformDataDTO{})

	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	res := api.PlatformDataDTO{}

	// This will take 1 second.
	err := client.Do(http.MethodGet, lol.BR1, lol.StatusURL, nil, &res, "")

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}

func TestInternalClientFailingRetry(t *testing.T) {
	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	var object api.PlainTextResponse

	// This will take 2 seconds.
	gotErr := client.Do(http.MethodGet, "tests", "/", nil, &object, "")

	wantErr := fmt.Errorf("Retried 2 times, stopping")

	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

func TestInternalClientRetryHeader(t *testing.T) {
	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429)

	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	var object api.PlainTextResponse

	gotErr := client.Do(http.MethodGet, "tests", "/", nil, &object, "")

	wantErr := fmt.Errorf("rate limited but no Retry-After header was found, stopping")

	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

// Testing if InternalClient.Do() can properly decode a plain text response.
func TestInternalClientDoRequest(t *testing.T) {
	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Post("/").
		Reply(200).BodyString("response")

	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	var object api.PlainTextResponse

	err := client.Do(http.MethodPost, "tests", "/", nil, &object, "")

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, object, "expecting non-nil response")
}

// Testing if the InternalClient can properly handle a status code not specified in the Riot API.
func TestInternalClientHandleErrorResponse(t *testing.T) {
	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(418).BodyString("response")

	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	var object api.PlainTextResponse

	gotErr := client.Do(http.MethodGet, "tests", "/", nil, &object, "")

	wantErr := api.ErrorResponse{
		Status: api.Status{
			Message:    "Unknown error",
			StatusCode: 418,
		},
	}

	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

func TestInternalClientNewRequest(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	validReq, _ := client.NewRequest(http.MethodGet, "http://localhost:80", nil)

	tests := []struct {
		name    string
		want    *http.Request
		wantErr error
		method  string
		url     string
	}{
		{
			name:   "success",
			want:   validReq,
			method: http.MethodGet,
			url:    "http://localhost:80",
		},
		{
			name:    "invalid method",
			wantErr: fmt.Errorf("net/http: invalid method \"=\""),
			method:  "=",
			url:     "http://localhost:80",
		},
		{
			name: "invalid url",
			wantErr: &url.Error{
				Op:  "parse",
				URL: "\\:invalid:/=",
				Err: fmt.Errorf("first path segment in URL cannot contain colon"),
			},
			method: http.MethodGet,
			url:    "\\:invalid:/=",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotData, gotErr := client.NewRequest(test.method, test.url, nil)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestInternalClientErrorResponses(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		name    string
		wantErr error
		setup   func()
		region  string
	}{
		{
			name:    "not found",
			wantErr: api.NotFoundError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(404)
			},
			region: "tests",
		},
		{
			name:    "rate limited with retry disabled",
			wantErr: api.RateLimitedError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(429)
			},
			region: "tests",
		},
		{
			name:    "unauthorized",
			wantErr: api.UnauthorizedError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(401)
			},
			region: "tests",
		},
		{
			name:    "forbidden",
			wantErr: api.ForbiddenError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(403)
			},
			region: "tests",
		},
		{
			name:    "region empty",
			wantErr: fmt.Errorf("region is required"),
			setup:   func() {},
			region:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()

			client := internal.NewInternalClient(&api.EquinoxConfig{
				Key:      "RIOT_API_KEY",
				Cluster:  api.Americas,
				LogLevel: api.DebugLevel,
				Timeout:  10,
				Retry:    false,
			})

			var gotData api.PlainTextResponse

			gotErr := client.Do(http.MethodGet, test.region, "/", nil, &gotData, "")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}
