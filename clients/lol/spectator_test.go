package lol_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestSpectatorFeaturedGames(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.FeaturedGamesDTO{}, &lol.FeaturedGamesDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.SpectatorFeaturedGamesURL
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Spectator.FeaturedGames(lol.BR1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSpectatorCurrentGame(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.CurrentGameInfoDTO{}, &lol.CurrentGameInfoDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.SpectatorCurrentGameURL, "summonerID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Spectator.CurrentGame(lol.BR1, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
