package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/require"
)

func TestMatchList(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	options := lol.MatchlistOptions{
		Start: 0,
		Count: 20,
	}

	res, err := client.Match.List(api.LOLRegionBR1, "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ", &options)

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

func TestMatchByID(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Match.ByID(api.LOLRegionBR1, "BR1_2499919790")

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

func TestMatchTimeline(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Match.Timeline(api.LOLRegionBR1, "BR1_2499919790")

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
