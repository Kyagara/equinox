//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/stretchr/testify/require"
)

func TestVALContent(t *testing.T) {
	checkIfOnlyDataDragon(t)
	content, err := client.VAL.ContentV1.Content(val.BR, "pt-BR")
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, content, "expecting non-nil content")
	require.NotEmpty(t, content.Version, "expecting non-nil version")
}

func TestVALRankedLeaderboard(t *testing.T) {
	checkIfOnlyDataDragon(t)
	leaderboard, err := client.VAL.RankedV1.Leaderboard(val.BR, "4401f9fd-4170-2e4c-4bc3-f3b4d7d150d1", 2, 0)
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, leaderboard, "expecting non-nil leaderboard")
	require.Equal(t, string(val.BR), leaderboard.Shard, "expecting shard to be 'br'")
	require.Equal(t, 2, len(leaderboard.Players), "expecting players to be equal to 2")
	require.Equal(t, "4401f9fd-4170-2e4c-4bc3-f3b4d7d150d1", leaderboard.ActID, "expecting act ID to be 4401f9fd-4170-2e4c-4bc3-f3b4d7d150d1")
}
