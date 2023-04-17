package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/util"
	"github.com/stretchr/testify/require"
)

func TestPlatformStatus(t *testing.T) {
	client, err := util.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := util.GetEndpointTestCases(api.PlatformDataDTO{}, &api.PlatformDataDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.StatusURL
			test.MockResponse(url, string(lol.BR1), test.AccessToken)
			gotData, gotErr := client.Status.PlatformStatus(lol.BR1)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
