package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClashTournaments(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.ClashTournamentDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]lol.ClashTournamentDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			region:  lol.BR1,
		}, {
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(lol.ClashURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.Tournaments(test.region)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestClashSummonerEntries(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.TournamentPlayerDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]lol.TournamentPlayerDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			region:  lol.BR1,
		}, {
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.ClashSummonerEntriesURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.SummonerEntries(test.region, "summonerID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestClashTournamentTeamByID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.TournamentTeamDto
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &lol.TournamentTeamDto{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			region:  lol.BR1,
		},
		{
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.ClashTournamentTeamByIDURL, "teamID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.TournamentTeamByID(test.region, "teamID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestClashByID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.ClashTournamentDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &lol.ClashTournamentDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			region:  lol.BR1,
		},
		{
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.ClashByIDURL, "tournamentID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.ByID(test.region, "tournamentID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestClashByTeamID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.ClashTournamentDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.ClashTournamentDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.ClashByTeamIDURL, "teamID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Clash.ByTeamID(lol.BR1, "teamID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
