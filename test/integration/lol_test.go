//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/stretchr/testify/require"
)

func TestLOLChampionRotation(t *testing.T) {
	checkIfOnlyDataDragon(t)
	ctx := context.Background()
	rotations, err := client.LOL.ChampionV3.Rotation(ctx, lol.BR1)
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, rotations, "expecting non-nil rotations")
}

func TestLOLMatchByID(t *testing.T) {
	checkIfOnlyDataDragon(t)
	ctx := context.Background()
	match, err := client.LOL.MatchV5.ByID(ctx, api.AMERICAS, "BR1_2744215970")
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, match, "expecting non-nil match")
	require.Equal(t, "BR1_2744215970", match.Metadata.MatchID, "expecting match ID to be BR1_2744215970")
}
