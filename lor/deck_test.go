package lor_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestDeckList(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lor.NewLORClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lor.DeckDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lor.DeckDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lor.Americas)).
				Get(lor.DeckURL).
				Reply(test.code).
				JSON(test.want).SetHeader("Authorization", "accessToken")

			gotData, gotErr := client.Deck.List(lor.Americas, "accessToken")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestDeckCreate(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

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
			wantErr: api.NotFoundError,
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
