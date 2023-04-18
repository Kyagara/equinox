package tft_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestLeagueEntries(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]tft.LeagueEntryDTO{}, &[]tft.LeagueEntryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": 1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": 1}

	tests = append(tests, test.TestCase[[]tft.LeagueEntryDTO, []tft.LeagueEntryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1, "tier": lol.BronzeTier, "page": 1},
	})

	tests = append(tests, test.TestCase[[]tft.LeagueEntryDTO, []tft.LeagueEntryDTO]{
		Name:      "invalid tier",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the tier specified is an apex tier, please use the corresponded method instead"),
		Options:   map[string]interface{}{"region": lol.BR1, "tier": lol.ChallengerTier, "page": 1},
	})

	tests = append(tests, test.TestCase[[]tft.LeagueEntryDTO, []tft.LeagueEntryDTO]{
		Name:    "default values",
		Code:    http.StatusOK,
		Want:    &[]tft.LeagueEntryDTO{},
		Options: map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": 0},
	})

	tests = append(tests, test.TestCase[[]tft.LeagueEntryDTO, []tft.LeagueEntryDTO]{
		Name:    "invalid page",
		Code:    http.StatusOK,
		Want:    &[]tft.LeagueEntryDTO{},
		Options: map[string]interface{}{"region": lol.BR1, "tier": lol.BronzeTier, "page": -1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tier := test.Options["tier"].(lol.Tier)
			region := test.Options["region"].(lol.Region)
			page := test.Options["page"].(int)
			url := fmt.Sprintf(tft.LeagueEntriesURL, tier, api.I)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.Entries(region, tier, api.I, page)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueByID(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(tft.LeagueListDTO{}, &tft.LeagueListDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[tft.LeagueListDTO, tft.LeagueListDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.LeagueByIDURL, "leagueID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.ByID(region, "leagueID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueTopRatedLadder(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]tft.TopRatedLadderEntryDTO{}, &[]tft.TopRatedLadderEntryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1, "queue": tft.RankedTFTTurboQueue}
	tests[1].Options = map[string]interface{}{"region": lol.BR1, "queue": tft.RankedTFTTurboQueue}

	tests = append(tests, test.TestCase[[]tft.TopRatedLadderEntryDTO, []tft.TopRatedLadderEntryDTO]{
		Name:      "invalid queue",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the queue specified is not available for the top rated ladder endpoint, please use the RankedTFTTurbo queue"),
		Options:   map[string]interface{}{"region": lol.BR1, "queue": tft.RankedTFTQueue},
	})

	tests = append(tests, test.TestCase[[]tft.TopRatedLadderEntryDTO, []tft.TopRatedLadderEntryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1, "queue": tft.RankedTFTTurboQueue},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			queue := test.Options["queue"].(tft.QueueType)
			region := test.Options["region"].(lol.Region)
			url := fmt.Sprintf(tft.LeagueRatedLaddersURL, queue)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.TopRatedLadder(region, queue)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueSummonerEntries(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]tft.LeagueEntryDTO{}, &[]tft.LeagueEntryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[[]tft.LeagueEntryDTO, []tft.LeagueEntryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.LeagueEntriesBySummonerURL, "summonerID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.SummonerEntries(region, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueChallenger(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(tft.LeagueListDTO{}, &tft.LeagueListDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[tft.LeagueListDTO, tft.LeagueListDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := tft.LeagueChallengerURL
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.Challenger(region)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueGrandmaster(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(tft.LeagueListDTO{}, &tft.LeagueListDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[tft.LeagueListDTO, tft.LeagueListDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := tft.LeagueGrandmasterURL
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.Grandmaster(region)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestLeagueMaster(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(tft.LeagueListDTO{}, &tft.LeagueListDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[tft.LeagueListDTO, tft.LeagueListDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := tft.LeagueMasterURL
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.League.Master(region)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
