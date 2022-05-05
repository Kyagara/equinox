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

func TestChampionMasterySummonerMasteries(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.ChampionMasteryDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &[]lol.ChampionMasteryDTO{},
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
			code:    http.StatusOK,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.ChampionMasteriesURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.ChampionMasteries.SummonerMasteries(test.region, "summonerID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChampionMasteryChampionScore(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.ChampionMasteryDTO
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &lol.ChampionMasteryDTO{},
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
			code:    http.StatusOK,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.ChampionMasteriesByChampionURL, "summonerID", 59)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.ChampionMasteries.ChampionScore(test.region, "summonerID", 59)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChampionMasteryMasteryScoreSum(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    int
		wantErr error
		region  lol.Region
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   0,
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
			code:    http.StatusOK,
			wantErr: fmt.Errorf("the region PBE1 is not available for this method"),
			region:  lol.PBE1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(lol.ChampionMasteriesScoresURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.ChampionMasteries.MasteryScoreSum(test.region, "summonerID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
