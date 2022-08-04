package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestLeagueEntries(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.LeagueEntryDTO
		wantErr error
		region  lol.Region
		tier    lol.Tier
		page    int
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]lol.LeagueEntryDTO{},
			region: lol.BR1,
			tier:   lol.BronzeTier,
			page:   1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			region:  lol.BR1,
			tier:    lol.BronzeTier,
			page:    1,
		},
		{
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
			tier:    lol.BronzeTier,
			page:    1,
		},
		{
			name:    "invalid tier",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the tier specified is an apex tier, please use the corresponded method instead"),
			region:  lol.BR1,
			tier:    lol.ChallengerTier,
			page:    1,
		},
		{
			name:   "invalid page",
			code:   http.StatusOK,
			want:   &[]lol.LeagueEntryDTO{},
			region: lol.BR1,
			tier:   lol.BronzeTier,
			page:   0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.LeagueEntriesURL, lol.Solo5x5Queue, test.tier, api.I)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Entries(test.region, lol.Solo5x5Queue, test.tier, api.I, test.page)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestLeagueByID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

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
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.LeagueByIDURL, "leagueID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.ByID(lol.BR1, "leagueID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestLeagueSummonerEntries(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.LeagueEntryDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]lol.LeagueEntryDTO{},
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
				Get(fmt.Sprintf(lol.LeagueEntriesBySummonerURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.SummonerEntries(test.region, "summonerID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestLeagueChallengerByQueue(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.LeagueListDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &lol.LeagueListDTO{},
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
				Get(fmt.Sprintf(lol.LeagueChallengerURL, lol.Solo5x5Queue)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.ChallengerByQueue(test.region, lol.Solo5x5Queue)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestLeagueGrandmasterByQueue(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

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
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.LeagueGrandmasterURL, lol.Solo5x5Queue)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.GrandmasterByQueue(lol.BR1, lol.Solo5x5Queue)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestLeagueMasterByQueue(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

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
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.LeagueMasterURL, lol.Solo5x5Queue)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.MasterByQueue(lol.BR1, lol.Solo5x5Queue)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
