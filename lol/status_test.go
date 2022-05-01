package lol_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestPlatformStatus(t *testing.T) {
	defer gock.Disable()

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
		Get(lol.StatusURL).
		Reply(200).
		JSON(&lol.PlatformDataDTO{})

	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Status.PlatformStatus(api.LOLRegionBR1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}
