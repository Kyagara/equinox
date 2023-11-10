//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/tft"
	"github.com/stretchr/testify/require"
)

func TestTFTSummonerByName(t *testing.T) {
	checkIfOnlyDataDragon(t)
	summoner, err := client.TFT.SummonerV1.ByName(tft.BR1, "Phanes")
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, summoner, "expecting non-nil summoner")
	require.Equal(t, "Phanes", summoner.Name, "expecting name to be equal to Phanes")
}
