package val_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestMatchList(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(val.MatchListDTO{}, &val.MatchListDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(val.MatchListURL, "PUUID")
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			gotData, gotErr := client.Match.List(val.BR, "PUUID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestMatchByID(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(val.MatchDTO{}, &val.MatchDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(val.MatchByIDURL, "matchID")
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			gotData, gotErr := client.Match.ByID(val.BR, "matchID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestMatchRecent(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(val.RecentMatchesDTO{}, &val.RecentMatchesDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(val.MatchRecentURL, val.CompetitiveQueue)
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			gotData, gotErr := client.Match.Recent(val.BR, val.CompetitiveQueue)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
