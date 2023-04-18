package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestPlatformStatus(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(api.PlatformDataDTO{}, &api.PlatformDataDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.StatusURL
			test.MockGetResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Status.PlatformStatus(lol.BR1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
