//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/stretchr/testify/require"
)

func TestLOLTournamentStubCreateProvider(t *testing.T) {
	checkIfOnlyDataDragon(t)
	provider, err := client.LOL.TournamentStubV5.RegisterProviderData(api.AMERICAS, &lol.StubProviderRegistrationParametersV5DTO{Region: string(lol.BR1), Url: "http://example.com/"})
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, provider, "expecting non-nil provider")
}
