package lor_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestInventoryCards(t *testing.T) {
	client, err := test.TestingNewLORClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lor.CardDTO{}, &[]lor.CardDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lor.InventoryURL
			test.MockGetResponse(url, string(lor.Americas), test.AccessToken)
			gotData, gotErr := client.Inventory.Cards(lor.Americas, "accessToken")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
