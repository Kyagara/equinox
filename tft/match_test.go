package tft_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/tft"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestMatchList(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    []string
		wantErr error
		count   int
	}{
		{
			name:  "found",
			code:  http.StatusOK,
			want:  []string{},
			count: 1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
			count:   1,
		},
		{
			name:  "default values",
			code:  http.StatusOK,
			want:  []string{},
			count: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Get(fmt.Sprintf(tft.MatchListURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.List("PUUID", test.code)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestMatchByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *tft.MatchDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &tft.MatchDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Get(fmt.Sprintf(tft.MatchByIDURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.ByID("matchID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
