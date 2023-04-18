package riot_test

import (
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestAccountPlayerActiveShard(t *testing.T) {
	client, err := test.TestingNewRiotClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(riot.ActiveShardDTO{}, &riot.ActiveShardDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(riot.AccountActiveShardURL, api.VAL, "PUUID")
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.Account.PlayerActiveShard("PUUID", api.VAL)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestAccountByPUUID(t *testing.T) {
	client, err := test.TestingNewRiotClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(riot.AccountDTO{}, &riot.AccountDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(riot.AccountByPUUIDURL, "PUUID")
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.Account.ByPUUID("PUUID")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestAccountByID(t *testing.T) {
	client, err := test.TestingNewRiotClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(riot.AccountDTO{}, &riot.AccountDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(riot.AccountByRiotIDURL, "gameName", "tagLine")
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.Account.ByID("gameName", "tagLine")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestAccountByAccessToken(t *testing.T) {
	client, err := test.TestingNewRiotClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(riot.AccountDTO{}, &riot.AccountDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := riot.AccountByAccessTokenURL
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.Account.ByAccessToken("accessToken")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
