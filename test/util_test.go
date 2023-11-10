package test_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestGetEndpointTestCases(t *testing.T) {
	tests := test.GetEndpointTestCases(lol.SummonerV4DTO{}, &lol.SummonerV4DTO{})
	require.Equal(t, "found", tests[0].Name, "expecting Name to be equal to found")
}

func TestMockGetResponse(t *testing.T) {
	client, err := test.TestingNewLOLClient()
	require.Nil(t, err, "expecting nil error")
	tests := test.GetEndpointTestCases(lol.SummonerV4DTO{}, &lol.SummonerV4DTO{})
	url := fmt.Sprintf("/lol/summoner/v4/summoners/by-name/%v", "summonerName")
	tests[0].MockGetResponse(url, string(lol.BR1), "")
	gotData, gotErr := client.SummonerV4.ByName(lol.BR1, "summonerName")
	tests[0].CheckResponse(t, gotData, gotErr)
	tests[0].AccessToken = "accessToken"
	tests[0].MockGetResponse("/lol/summoner/v4/summoners/me", string(lol.BR1), "accessToken")
	gotData, gotErr = client.SummonerV4.ByAccessToken(lol.BR1, "accessToken")
	tests[0].CheckResponse(t, gotData, gotErr)
}

func TestMockPostResponse(t *testing.T) {
	client, err := test.TestingNewLOLClient()
	require.Nil(t, err, "expecting nil error")
	tests := test.GetEndpointTestCases(*new(int32), new(int32))
	tests[0].MockPostResponse("/lol/tournament/v5/tournaments", string(api.AMERICAS), "")
	gotData, gotErr := client.TournamentV5.RegisterTournament(api.AMERICAS, &lol.TournamentRegistrationParametersV5DTO{})
	tests[0].CheckResponse(t, &gotData, gotErr)
}
