package lor_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestPlatformStatus(t *testing.T) {
	client, err := test.TestingNewLORClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(api.PlatformDataDTO{}, &api.PlatformDataDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lor.StatusURL
			test.MockGetResponse(url, string(lor.Americas), test.AccessToken)
			gotData, gotErr := client.Status.PlatformStatus(lor.Americas)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
