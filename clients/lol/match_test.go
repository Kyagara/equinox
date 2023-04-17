package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/util"
	"github.com/stretchr/testify/require"
)

func TestMatchList(t *testing.T) {
	client, err := util.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := util.GetEndpointTestCases([]string{}, &lol.MatchlistOptions{})

	tests[0].Parameters = &lol.MatchlistOptions{Start: 0,
		Count: 20}
	tests[1].Parameters = &lol.MatchlistOptions{Start: 0,
		Count: 20}

	tests = append(tests, util.TestCase[[]string, lol.MatchlistOptions]{
		Name: "all optional fields set",
		Code: http.StatusOK,
		Want: &[]string{},
		Parameters: &lol.MatchlistOptions{
			StartTime: 1,
			EndTime:   1,
			Queue:     420,
			Type:      lol.RankedMatch,
			Start:     1,
			Count:     1,
		},
	})

	tests = append(tests, util.TestCase[[]string, lol.MatchlistOptions]{
		Name:       "nil options",
		Code:       http.StatusOK,
		Want:       &[]string{},
		Parameters: nil,
	})

	tests = append(tests, util.TestCase[[]string, lol.MatchlistOptions]{
		Name: "count > 100",
		Code: http.StatusOK,
		Want: &[]string{},
		Parameters: &lol.MatchlistOptions{Start: 0,
			Count: 101},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.MatchListURL, "PUUID")
			test.MockResponse(url, string(lol.Americas), test.AccessToken)
			gotData, gotErr := client.Match.List(lol.Americas, "PUUID", test.Parameters)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestMatchByID(t *testing.T) {
	client, err := util.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := util.GetEndpointTestCases(lol.MatchDTO{}, &lol.MatchDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.MatchByIDURL, "matchID")
			test.MockResponse(url, string(lol.Americas), test.AccessToken)
			gotData, gotErr := client.Match.ByID(lol.Americas, "matchID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestMatchTimeline(t *testing.T) {
	client, err := util.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := util.GetEndpointTestCases(lol.MatchTimelineDTO{}, &lol.MatchTimelineDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.MatchByIDURL, "matchID")
			test.MockResponse(url, string(lol.Americas), test.AccessToken)
			gotData, gotErr := client.Match.Timeline(lol.Americas, "matchID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
