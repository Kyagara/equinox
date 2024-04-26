//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/require"
)

func TestRiotAccountByID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	account, err := client.Riot.AccountV1.ByRiotID(ctx, api.AMERICAS, "Kevin", "FFXIV")
	require.NoError(t, err)
	require.NotEmpty(t, account, "expecting non-nil account")
	require.Equal(t, "Kevin", account.GameName, "expecting gameName to be equal to Kevin")
	require.Equal(t, "FFXIV", account.TagLine, "expecting tagLine to be equal to FFXIV")
}
