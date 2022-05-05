package val_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/val"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestMatchList(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *val.MatchListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &val.MatchListDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, val.BR)).
				Get(fmt.Sprintf(val.MatchListURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.List(val.BR, "PUUID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestMatchByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *val.MatchDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &val.MatchDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, val.BR)).
				Get(fmt.Sprintf(val.MatchByIDURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.ByID(val.BR, "matchID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestMatchRecent(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *val.RecentMatchesDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &val.RecentMatchesDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, val.BR)).
				Get(fmt.Sprintf(val.MatchRecentURL, val.CompetitiveQueue)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.Recent(val.BR, val.CompetitiveQueue)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
