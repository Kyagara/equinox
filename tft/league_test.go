package tft_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/Kyagara/equinox/tft"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestLeagueEntries(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]tft.LeagueEntryDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]tft.LeagueEntryDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(tft.LeagueEntriesURL, lol.GoldTier, api.I)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Entries(lol.BR1, lol.GoldTier, api.I, 1)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *tft.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &tft.LeagueListDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(tft.LeagueByIDURL, "leagueID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.ByID(lol.BR1, "leagueID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueTopRatedLadder(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]tft.TopRatedLadderEntryDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]tft.TopRatedLadderEntryDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(tft.LeagueRatedLaddersURL, tft.RankedTFTTurbo)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.TopRatedLadder(lol.BR1, tft.RankedTFTTurbo)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueSummonerEntries(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]tft.LeagueEntryDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]tft.LeagueEntryDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(tft.LeagueEntriesBySummonerURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.SummonerEntries(lol.BR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueChallenger(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *tft.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &tft.LeagueListDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(tft.LeagueChallengerURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Challenger(lol.BR1)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueGrandmaster(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *tft.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &tft.LeagueListDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(tft.LeagueGrandmasterURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Grandmaster(lol.BR1)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestLeagueMaster(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *tft.LeagueListDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &tft.LeagueListDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(tft.LeagueMasterURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Master(lol.BR1)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
