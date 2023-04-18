package tft_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestMatchList(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]string{}, &[]string{})

	tests[0].Options = map[string]interface{}{"count": 1}
	tests[1].Options = map[string]interface{}{"count": 1}

	tests = append(tests, test.TestCase[[]string, []string]{
		Name:    "default values",
		Code:    http.StatusOK,
		Want:    &[]string{},
		Options: map[string]interface{}{"count": 0},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.MatchListURL, "PUUID")
			test.MockGetResponse(url, string(lol.Americas), test.AccessToken)
			count := test.Options["count"].(int)
			gotData, gotErr := client.Match.List(lol.Americas, "PUUID", count)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestMatchByID(t *testing.T) {
	client, err := test.TestingNewTFTClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(tft.MatchDTO{}, &tft.MatchDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(tft.MatchByIDURL, "matchID")
			test.MockGetResponse(url, string(lol.Americas), test.AccessToken)
			gotData, gotErr := client.Match.ByID(lol.Americas, "matchID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
