package lor_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestMatchList(t *testing.T) {
	client, err := test.TestingNewLORClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]string{}, &[]string{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lor.MatchListURL, "PUUID")
			test.MockGetResponse(url, string(lor.Americas), test.AccessToken)
			gotData, gotErr := client.Match.List(lor.Americas, "PUUID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestMatchByID(t *testing.T) {
	client, err := test.TestingNewLORClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lor.MatchDTO{}, &lor.MatchDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lor.MatchByIDURL, "matchID")
			test.MockGetResponse(url, string(lor.Americas), test.AccessToken)
			gotData, gotErr := client.Match.ByID(lor.Americas, "matchID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
