package test_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestGetEndpointTestCases(t *testing.T) {
	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	require.Equal(t, "found", tests[0].Name, "expecting Name to be equal to found")
}

func TestMockResponse(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.SummonerDTO{}, &lol.SummonerDTO{})

	url := fmt.Sprintf(lol.SummonerByAccountIDURL, "summonerName")
	tests[0].MockResponse(url, string(lol.BR1), "")
	gotData, gotErr := client.Summoner.ByAccountID(lol.BR1, "summonerName")
	tests[0].CheckResponse(t, gotData, gotErr)

	url = lol.SummonerByAccessTokenURL

	tests[0].AccessToken = "accessToken"
	tests[0].MockResponse(url, string(lol.BR1), "accessToken")
	gotData, gotErr = client.Summoner.ByAccessToken(lol.BR1, "accessToken")
	tests[0].CheckResponse(t, gotData, gotErr)
}

func TestNewLOLClientForTesting(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, client, "expecting non-nil client")
}
