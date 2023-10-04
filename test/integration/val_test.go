//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/stretchr/testify/require"
)

func TestVALContentAllLocales(t *testing.T) {
	content, err := client.VAL.Content.AllLocales(val.BR)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, content, "expecting non-nil content")

	require.NotNil(t, content.Version, "expecting non-nil version")

	require.NotNil(t, content.Characters[0].LocalizedNames.EnUS, "expecting non-nil localized character name")
}

func TestVALRankedLeaderboard(t *testing.T) {
	leaderboard, err := client.VAL.Ranked.LeaderboardsByActID(val.BR, "3e47230a-463c-a301-eb7d-67bb60357d4f", 2, 0)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, leaderboard, "expecting non-nil leaderboard")

	require.Equal(t, val.BR, leaderboard.Shard, "expecting non-nil leaderboard")

	require.Equal(t, 2, len(leaderboard.Players), "expecting non-nil leaderboard")

	require.Equal(t, "3e47230a-463c-a301-eb7d-67bb60357d4f", leaderboard.ActID, "expecting non-nil leaderboard")
}
