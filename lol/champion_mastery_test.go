package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/require"
)

func TestChampionMasteriesBySummonerID(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.ChampionMasteries.BySummonerID(api.LOLRegionBR1, "mqk6ubCanzRDH9PPSLMNhIi1PAvjAYh9hTip8daGU2aACQ")

	if err != nil && err == api.NotFoundError {
		require.NotNil(t, err, "expecting non-nil error")
	}

	if err == nil {
		require.NotNil(t, res, "expecting non-nil response")
	}
}

func TestChampionMasteriesChampionScore(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.ChampionMasteries.ChampionScore(api.LOLRegionBR1, "mqk6ubCanzRDH9PPSLMNhIi1PAvjAYh9hTip8daGU2aACQ", 59)

	if err != nil && err == api.NotFoundError {
		require.NotNil(t, err, "expecting non-nil error")
	}

	if err == nil {
		require.NotNil(t, res, "expecting non-nil response")
	}
}

func TestChampionMasteriesMasteryScore(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.ChampionMasteries.MasteryScore(api.LOLRegionBR1, "mqk6ubCanzRDH9PPSLMNhIi1PAvjAYh9hTip8daGU2aACQ")

	if err != nil && err == api.NotFoundError {
		require.NotNil(t, err, "expecting non-nil error")
	}

	if err == nil {
		require.NotNil(t, res, "expecting non-nil response")
	}
}
