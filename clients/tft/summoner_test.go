package tft_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestSummonerByAccountID(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.SummonerByAccountIDURL, "accountID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByAccountID(lol.BR1, "accountID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByName(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.SummonerByNameURL, "summonerName")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByName(lol.BR1, "summonerName")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByPUUID(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.SummonerByPUUIDURL, "PUUID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByPUUID(lol.BR1, "PUUID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByID(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.SummonerByIDURL, "summonerID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByID(lol.BR1, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByAccessToken(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := tft.SummonerByAccessTokenURL
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByAccessToken(lol.BR1, "accessToken")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
