package val_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestPlatformStatus(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(api.PlatformDataDTO{}, &api.PlatformDataDTO{})

	tests[0].Options = map[string]interface{}{"region": val.BR}
	tests[1].Options = map[string]interface{}{"region": val.BR}

	tests = append(tests, test.TestCase[api.PlatformDataDTO, api.PlatformDataDTO]{
		Name:      "invalid region",
		Code:      http.StatusNotFound,
		WantError: fmt.Errorf("the region ESPORTS is not available for this method"),
		Options:   map[string]interface{}{"region": val.ESPORTS},
	})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := val.StatusURL
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			region := test.Options["region"].(val.Shard)
			gotData, gotErr := client.Status.PlatformStatus(region)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
