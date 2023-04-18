package val_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/test"
	"github.com/stretchr/testify/require"
)

func TestContentAllLocales(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(val.ContentDTO{}, &val.ContentDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := val.ContentURL
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			gotData, gotErr := client.Content.AllLocales(val.BR)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestContentByLocale(t *testing.T) {
	client, err := test.TestingNewVALClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(val.LocalizedContentDTO{}, &val.LocalizedContentDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := val.ContentURL
			test.MockGetResponse(url, string(val.BR), test.AccessToken)
			gotData, gotErr := client.Content.ByLocale(val.BR, val.PortugueseBR)
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}
