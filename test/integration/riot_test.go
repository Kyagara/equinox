//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRiotAccountByID(t *testing.T) {
	account, err := client.Riot.Account.ByID("Loveable Senpai", "SUN")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, account, "expecting non-nil account")

	require.Equal(t, "Loveable Senpai", account.GameName, "expecting gameName to be equal to Loveable Senpai")

	assert.Equal(t, "SUN", account.TagLine, "expecting tagLine to be equal to SUN")
}
