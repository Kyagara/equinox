package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
)

func TestPlatformStatus(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Status.PlatformStatus(api.LOLRegionNA1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}
