//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLOLChampionsRotations(t *testing.T) {
	if OnlyDataDragon {
		t.SkipNow()
	}

	rotations, err := client.LOL.Champion.Rotations(lol.BR1)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rotations, "expecting non-nil rotations")

	assert.Greater(t, len(rotations.FreeChampionIDs), 0, "expecting free champions length greater than 0")
}

func TestLOLTournamentStubCreateProvider(t *testing.T) {
	provider, err := client.LOL.TournamentStub.CreateProvider(lol.BR, "http://example.com/")

	require.Nil(t, err, "expecting nil error")

	assert.NotNil(t, provider, "expecting non-nil provider")
}
