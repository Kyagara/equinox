package lol_test

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
)

func TestFreeChampionsRotation(t *testing.T) {
	internalClient := internal.NewInternalClient(os.Getenv("RIOT_API_KEY"), true)

	client := lol.NewLOLClient(internalClient)

	res, err := client.Champion.FreeRotation(api.LOLRegionBR1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}
