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

func TestLeagueEntries(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.LeagueEntryDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lol.LeagueEntryDTO{},
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
				Get(fmt.Sprintf(lol.LeagueURL, api.I, api.LOLTierGold, api.RankedFlexSRQueueType)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Entries(api.LOLRegionBR1, api.I, api.LOLTierGold, api.RankedFlexSRQueueType, 1)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.LeagueListDTO{},
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
				Get(fmt.Sprintf(lol.LeagueByID, "leagueID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.ByID(api.LOLRegionBR1, "leagueID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueEntriesBySummoner(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.LeagueEntryDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lol.LeagueEntryDTO{},
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
				Get(fmt.Sprintf(lol.LeagueEntriesBySummonerURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.EntriesBySummonerID(api.LOLRegionBR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueChallengerByQueue(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.LeagueListDTO{},
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
				Get(fmt.Sprintf(lol.LeagueChallengerURL, api.RankedFlexSRQueueType)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.ChallengerByQueue(api.LOLRegionBR1, api.RankedFlexSRQueueType)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueGrandmasterByQueue(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.LeagueListDTO{},
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
				Get(fmt.Sprintf(lol.LeagueGrandmasterURL, api.RankedFlexSRQueueType)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.GrandmasterByQueue(api.LOLRegionBR1, api.RankedFlexSRQueueType)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueMasterByQueue(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.LeagueListDTO{},
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
				Get(fmt.Sprintf(lol.LeagueMasterURL, api.RankedFlexSRQueueType)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.MasterByQueue(api.LOLRegionBR1, api.RankedFlexSRQueueType)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
