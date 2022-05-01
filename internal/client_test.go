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

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
		Get(lol.StatusURL).
		Reply(429).SetHeader("Retry-After", "1").
		JSON(&lol.PlatformDataDTO{})

	config := &api.EquinoxConfig{
		Key:     "RIOT_API_KEY",
		Debug:   true,
		Timeout: 10,
		Retry:   true,
	}

	client := internal.NewInternalClient(config)

	res := lol.PlatformDataDTO{}

	// This will take 1 second
	err := client.Do(http.MethodGet, api.LOLRegionBR1, lol.StatusURL, nil, &res)

	assert.NotNil(t, err, "expecting non-nil error")
}
