package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestClashTournaments(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.TournamentDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lol.TournamentDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
				Get(lol.ClashURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.Tournaments(api.LOLRegionBR1)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestClashSummonerEntries(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.TournamentPlayerDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lol.TournamentPlayerDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
				Get(fmt.Sprintf(lol.ClashPlayersBySummonerIDURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.SummonerEntries(api.LOLRegionBR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestClashTournamentTeamByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.TournamentTeamDto
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.TournamentTeamDto{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
				Get(fmt.Sprintf(lol.ClashTeamByIDURL, "teamID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.TournamentTeamByID(api.LOLRegionBR1, "teamID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestClashByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.TournamentDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.TournamentDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
				Get(fmt.Sprintf(lol.ClashTournamentByIDURL, "tournamentID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.ByID(api.LOLRegionBR1, "tournamentID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestClashByTeamID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.TournamentDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.TournamentDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
				Get(fmt.Sprintf(lol.ClashTournamentByTeamIDURL, "teamID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.ByTeamID(api.LOLRegionBR1, "teamID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
