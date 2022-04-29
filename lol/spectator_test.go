package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpectatorFeaturedGames(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Spectator.FeaturedGames(api.LOLRegionNA1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}

func TestSpectatorCurrentGame(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Spectator.CurrentGame(api.LOLRegionBR1, "mqk6ubCanzRDH9PPSLMNhIi1PAvjAYh9hTip8daGU2aACQ")

	// What should be done in cases where a 404 is a valid response?

	// If there's an error, it could be that no summoner was in a match when this was called.

	// How can we test this?
	if err != nil && err == api.NotFoundError {
		require.NotNil(t, err, "expecting non-nil error")
	}

	if err == nil {
		require.NotNil(t, res, "expecting non-nil response")
	}
}
