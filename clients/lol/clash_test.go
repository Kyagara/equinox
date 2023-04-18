package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestClashTournaments(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.ClashTournamentDTO{}, &[]lol.ClashTournamentDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[[]lol.ClashTournamentDTO, []lol.ClashTournamentDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.ClashURL
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.Clash.Tournaments(region)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestClashSummonerEntries(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.TournamentPlayerDTO{}, &[]lol.TournamentPlayerDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[[]lol.TournamentPlayerDTO, []lol.TournamentPlayerDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ClashSummonerEntriesURL, "summonerID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.Clash.SummonerEntries(region, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestClashTournamentTeamByID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.TournamentTeamDto{}, &lol.TournamentTeamDto{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[lol.TournamentTeamDto, lol.TournamentTeamDto]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ClashTournamentTeamByIDURL, "teamID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.Clash.TournamentTeamByID(region, "teamID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestClashByID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.ClashTournamentDTO{}, &lol.ClashTournamentDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[lol.ClashTournamentDTO, lol.ClashTournamentDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ClashByIDURL, "tournamentID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.Clash.ByID(region, "tournamentID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestClashByTeamID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.ClashTournamentDTO{}, &lol.ClashTournamentDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ClashByTeamIDURL, "teamID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Clash.ByTeamID(lol.BR1, "teamID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
