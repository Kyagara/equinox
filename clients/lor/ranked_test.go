package lor_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestRankedLeaderboards(t *testing.T) {
	client, err := test.TestingNewLORClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lor.LeaderboardDTO{}, &lor.LeaderboardDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lor.RankedURL
			test.MockGetResponse(url, string(lor.Americas), test.AccessToken)
			gotData, gotErr := client.Ranked.Leaderboards(lor.Americas)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
