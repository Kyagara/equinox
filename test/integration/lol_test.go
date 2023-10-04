//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/stretchr/testify/require"
)

func TestLOLChampionsRotations(t *testing.T) {
	checkIfOnlyDataDragon(t)
	rotations, err := client.LOL.Champion.Rotations(lol.BR1)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rotations, "expecting non-nil rotations")

	require.Greater(t, len(rotations.FreeChampionIDs), 0, "expecting free champions length greater than 0")
}

func TestLOLTournamentStubCreateProvider(t *testing.T) {
	checkIfOnlyDataDragon(t)
	provider, err := client.LOL.TournamentStub.CreateProvider(lol.BR, "http://example.com/")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, provider, "expecting non-nil provider")
}
