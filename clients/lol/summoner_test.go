package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestSummonerByAccountID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.SummonerByAccountIDURL, "summonerName")
			test.MockResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByAccountID(lol.BR1, "summonerName")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByName(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.SummonerByNameURL, "summonerName")
			test.MockResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByName(lol.BR1, "summonerName")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByPUUID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.SummonerByPUUIDURL, "summonerName")
			test.MockResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByPUUID(lol.BR1, "summonerName")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.SummonerByIDURL, "summonerName")
			test.MockResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Summoner.ByID(lol.BR1, "summonerName")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestSummonerByAccessToken(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})
	tests[0].AccessToken = "accessToken"
	tests[1].AccessToken = "accessToken"

	tests = append(tests, test.TestCase[lol.SummonerDTO, lol.SummonerDTO]{
		Name:      "accessToken empty",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("accessToken is required"),
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.SummonerByAccessTokenURL
			test.MockResponse(url, string(lol.BR1), test.AccessToken)

			gotData, gotErr := client.Summoner.ByAccessToken(lol.BR1, test.AccessToken)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
