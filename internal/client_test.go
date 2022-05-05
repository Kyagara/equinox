package internal_test

import (
	"fmt"
	"net/http"
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

	// This will take 1 second
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

// Testing if the Do() method can properly decode a plain text response
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

// Testing if the client can properly handle a status code not specified in the Riot API
func TestInternalClientSendRequest(t *testing.T) {
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
