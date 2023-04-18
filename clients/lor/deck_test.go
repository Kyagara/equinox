package lor_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeckList(t *testing.T) {
	client, err := test.TestingNewLORClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases([]lor.DeckDTO{}, &[]lor.DeckDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lor.DeckURL
			test.MockGetResponse(url, string(lor.Americas), test.AccessToken)
			gotData, gotErr := client.Deck.List(lor.Americas, "accessToken")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestDeckCreate(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lor.NewLORClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    string
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: "response",
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lor.Americas)).
				Post(lor.DeckURL).
				Reply(test.code).SetHeader("Authorization", "accessToken").BodyString("response")

			gotData, gotErr := client.Deck.Create(lor.Americas, "accessToken", "code", "name")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
