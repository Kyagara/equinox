//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/stretchr/testify/require"
)

func TestVALContentAllLocales(t *testing.T) {
	checkIfOnlyDataDragon(t)
	content, err := client.VAL.ContentV1.Content(val.BR, "pt-BR")
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, content, "expecting non-nil content")
	require.NotNil(t, content.Version, "expecting non-nil version")
}

func TestVALRankedLeaderboard(t *testing.T) {
	checkIfOnlyDataDragon(t)
	leaderboard, err := client.VAL.RankedV1.Leaderboard(val.BR, "4401f9fd-4170-2e4c-4bc3-f3b4d7d150d1", 2, 0)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, leaderboard, "expecting non-nil leaderboard")
	require.Equal(t, string(val.BR), leaderboard.Shard, "expecting non-nil leaderboard")
	require.Equal(t, 2, len(leaderboard.Players), "expecting non-nil leaderboard")
	require.Equal(t, "4401f9fd-4170-2e4c-4bc3-f3b4d7d150d1", leaderboard.ActID, "expecting non-nil leaderboard")
}
