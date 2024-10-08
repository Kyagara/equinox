//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/clients/lol"
	"github.com/stretchr/testify/require"
)

func TestLOLChampionRotation(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	rotations, err := client.LOL.ChampionV3.Rotation(ctx, lol.JP1)
	require.NoError(t, err)
	require.NotEmpty(t, rotations, "expecting non-nil rotations")
}

func TestLOLMatchByIDAndCache(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	match, err := client.LOL.MatchV5.ByID(ctx, api.ASIA, "KR_7014499581")
	require.NoError(t, err)
	require.NotEmpty(t, match, "expecting non-nil match")
	require.Equal(t, "KR_7014499581", match.Metadata.MatchID, "expecting match ID to be KR_7014499581")

	ctx = context.Background()
	match, err = client.LOL.MatchV5.ByID(ctx, api.ASIA, "KR_7014499581")
	require.NoError(t, err)
	require.NotEmpty(t, match, "expecting non-nil match")
	require.Equal(t, "KR_7014499581", match.Metadata.MatchID, "expecting match ID to be KR_7014499581")
}
