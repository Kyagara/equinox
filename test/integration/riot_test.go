//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/require"
)

func TestRiotAccountByID(t *testing.T) {
	checkIfOnlyDataDragon(t)
	account, err := client.Riot.AccountV1.ByRiotID(api.AMERICAS, "Loveable Senpai", "SUN")
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, account, "expecting non-nil account")
	require.Equal(t, "Loveable Senpai", account.GameName, "expecting gameName to be equal to Loveable Senpai")
	require.Equal(t, "SUN", account.TagLine, "expecting tagLine to be equal to SUN")
}
