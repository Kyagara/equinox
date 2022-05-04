package internal_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
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
