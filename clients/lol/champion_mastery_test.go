package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestChampionMasterySummonerMasteries(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.ChampionMasteryDTO{}, &[]lol.ChampionMasteryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[[]lol.ChampionMasteryDTO, []lol.ChampionMasteryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChampionMasteriesURL, "summonerID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.ChampionMasteries.SummonerMasteries(region, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChampionMasteryChampionScore(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.ChampionMasteryDTO{}, &lol.ChampionMasteryDTO{})

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[lol.ChampionMasteryDTO, lol.ChampionMasteryDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChampionMasteriesByChampionURL, "summonerID", 59)
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.ChampionMasteries.ChampionScore(region, "summonerID", 59)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChampionMasteryMasteryScoreSum(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(*new(int), new(int))

	tests[0].Options = map[string]interface{}{"region": lol.BR1}
	tests[1].Options = map[string]interface{}{"region": lol.BR1}

	tests = append(tests, test.TestCase[int, int]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region PBE1 is not available for this method"),
		Options:   map[string]interface{}{"region": lol.PBE1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChampionMasteriesScoresURL, "summonerID")
			region := test.Options["region"].(lol.Region)
			test.MockGetResponse(url, string(region), test.AccessToken)
			gotData, gotErr := client.ChampionMasteries.MasteryScoreSum(region, "summonerID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
