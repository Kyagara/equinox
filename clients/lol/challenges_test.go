package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestChallengesList(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.ChallengeConfigInfoDTO{}, &[]lol.ChallengeConfigInfoDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.ChallengesConfigurationsURL
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Challenges.List(lol.BR1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChallengesByID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.ChallengeConfigInfoDTO{}, &lol.ChallengeConfigInfoDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChallengesConfigurationByIDURL, 1)
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Challenges.ByID(lol.BR1, 1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChallengesPercentiles(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(map[int64]lol.PercentileDTO{}, &map[int64]lol.PercentileDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.ChallengesPercentilesURL
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Challenges.Percentiles(lol.BR1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChallengesPercentilesByID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.PercentileDTO{}, &lol.PercentileDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChallengesPercentileByIDURL, 1)
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Challenges.PercentilesByID(lol.BR1, 1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChallengesLeaderboards(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lol.ApexPlayerInfoDTO{}, &[]lol.ApexPlayerInfoDTO{})

	tests[0].Options = map[string]interface{}{"limit": 0}
	tests[0].Options = map[string]interface{}{"limit": 0}

	tests = append(tests, test.TestCase[[]lol.ApexPlayerInfoDTO, []lol.ApexPlayerInfoDTO]{
		Name:    "limit > 0",
		Code:    http.StatusOK,
		Want:    &[]lol.ApexPlayerInfoDTO{},
		Options: map[string]interface{}{"limit": 1},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChallengesLeaderboardsByLevelURL, 1, lol.MasterLevel)
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			limit := 0
			if test.Options["limit"] != nil {
				limit = test.Options["limit"].(int)
			}
			gotData, gotErr := client.Challenges.Leaderboards(lol.BR1, 1, lol.MasterLevel, limit)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestChallengesByPUUID(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.PlayerInfoDTO{}, &lol.PlayerInfoDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.ChallengesByPUUIDURL, "PUUID")
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Challenges.ByPUUID(lol.BR1, "PUUID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
