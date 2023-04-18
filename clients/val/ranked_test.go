package val_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestRankedLeaderboardsByActID(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(val.LeaderboardDTO{}, &val.LeaderboardDTO{})

	tests[0].Options = map[string]interface{}{"region": val.BR, "size": uint8(1), "start": 0}
	tests[1].Options = map[string]interface{}{"region": val.BR, "size": uint8(1), "start": 0}

	tests = append(tests, test.TestCase[val.LeaderboardDTO, val.LeaderboardDTO]{
		Name:    "default values",
		Code:    http.StatusOK,
		Want:    &val.LeaderboardDTO{},
		Options: map[string]interface{}{"region": val.BR, "size": uint8(0), "start": -1},
	})

	tests = append(tests, test.TestCase[val.LeaderboardDTO, val.LeaderboardDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region ESPORTS is not available for this method"),
		Options:   map[string]interface{}{"region": val.ESPORTS, "size": uint8(1), "start": 0},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(val.RankedURL, "actID")
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			region := test.Options["region"].(val.Shard)
			size := test.Options["size"].(uint8)
			start := test.Options["start"].(int)
			gotData, gotErr := client.Ranked.LeaderboardsByActID(region, "actID", size, start)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
