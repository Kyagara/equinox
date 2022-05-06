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
		region  lol.Region
		tier    lol.Tier
		page    int
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]tft.LeagueEntryDTO{},
			region: lol.BR1,
			tier:   lol.BronzeTier,
			page:   1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
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
			name:   "default values",
			code:   http.StatusOK,
			want:   &[]tft.LeagueEntryDTO{},
			region: lol.BR1,
			tier:   lol.BronzeTier,
			page:   0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(tft.LeagueEntriesURL, test.tier, api.I)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Entries(test.region, test.tier, api.I, test.page)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &tft.LeagueListDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
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
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(tft.LeagueByIDURL, "leagueID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.ByID(test.region, "leagueID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
		queue   tft.QueueType
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]tft.TopRatedLadderEntryDTO{},
			queue:  tft.RankedTFTTurboQueue,
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
			queue:   tft.RankedTFTTurboQueue,
			region:  lol.BR1,
		},
		{
			name:    "invalid queue",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the queue specified is not available for the top rated ladder endpoint, please use the RankedTFTTurbo queue"),
			queue:   tft.RankedTFTQueue,
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
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(tft.LeagueRatedLaddersURL, test.queue)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.TopRatedLadder(test.region, test.queue)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]tft.LeagueEntryDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
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
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(tft.LeagueEntriesBySummonerURL, "summonerID")).
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

func TestLeagueChallenger(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := tft.NewTFTClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *tft.LeagueListDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &tft.LeagueListDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
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
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(tft.LeagueChallengerURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Challenger(test.region)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &tft.LeagueListDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
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
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(tft.LeagueGrandmasterURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Grandmaster(test.region)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &tft.LeagueListDTO{},
			region: lol.BR1,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
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
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(tft.LeagueMasterURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.League.Master(test.region)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
