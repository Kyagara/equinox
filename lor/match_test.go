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

func TestMatchList(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lor.NewLORClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    []string
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: []string{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lor.Americas)).
				Get(fmt.Sprintf(lor.MatchListURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.List(lor.Americas, "PUUID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestMatchByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lor.NewLORClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lor.MatchDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lor.MatchDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lor.Americas)).
				Get(fmt.Sprintf(lor.MatchByIDURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.ByID(lor.Americas, "matchID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
