package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestLeagueEntries(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.LeagueEntryDTO{}, &[]lol.LeagueEntryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": 1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": 1}

	tests = append(tests, test.TestCase[[]lol.LeagueEntryDTO, []lol.LeagueEntryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1, "tier": lol.BronzeTier, "page": 1},
	})

	tests = append(tests, test.TestCase[[]lol.LeagueEntryDTO, []lol.LeagueEntryDTO]{
		Name:      "invalid tier",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the tier specified is an apex tier, please use the corresponded method instead"),
		Options:   map[string]interface{}{"region": lol.BR1, "tier": lol.ChallengerTier, "page": 1},
	})

	tests = append(tests, test.TestCase[[]lol.LeagueEntryDTO, []lol.LeagueEntryDTO]{
		Name:    "default values",
		Code:    http.StatusOK,
		Want:    &[]lol.LeagueEntryDTO{},
		Options: map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": 0},
	})

	tests = append(tests, test.TestCase[[]lol.LeagueEntryDTO, []lol.LeagueEntryDTO]{
		Name:    "invalid page",
		Code:    http.StatusOK,
		Want:    &[]lol.LeagueEntryDTO{},
		Options: map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": -1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tier := test.Options["tier"].(lol.Tier)
			region := test.Options["region"].(lol.Region)
			page := test.Options["page"].(int)
			url := fmt.Sprintf(lol.LeagueEntriesURL, lol.Solo5x5Queue, tier, api.I)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.Entries(region, lol.Solo5x5Queue, tier, api.I, page)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueByID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.LeagueListDTO{}, &lol.LeagueListDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.LeagueByIDURL, "leagueID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.League.ByID(lol.BR1, "leagueID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueSummonerEntries(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.LeagueEntryDTO{}, &[]lol.LeagueEntryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[[]lol.LeagueEntryDTO, []lol.LeagueEntryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.LeagueEntriesBySummonerURL, "summonerID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.SummonerEntries(region, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueChallengerByQueue(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.LeagueListDTO{}, &lol.LeagueListDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[lol.LeagueListDTO, lol.LeagueListDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.LeagueChallengerURL, lol.Solo5x5Queue)
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.ChallengerByQueue(region, lol.Solo5x5Queue)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueGrandmasterByQueue(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.LeagueListDTO{}, &lol.LeagueListDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.LeagueGrandmasterURL, lol.Solo5x5Queue)
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.League.GrandmasterByQueue(lol.BR1, lol.Solo5x5Queue)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueMasterByQueue(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.LeagueListDTO{}, &lol.LeagueListDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.LeagueMasterURL, lol.Solo5x5Queue)
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.League.MasterByQueue(lol.BR1, lol.Solo5x5Queue)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
