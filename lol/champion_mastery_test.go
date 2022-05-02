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
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lol.ChampionMasteryDTO{},
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
				Get(fmt.Sprintf(lol.ChampionMasteriesURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.ChampionMasteries.SummonerMasteries(api.LOLRegionBR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
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
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.ChampionMasteryDTO{},
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
				Get(fmt.Sprintf(lol.ChampionMasteriesByChampionURL, "summonerID", 59)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.ChampionMasteries.ChampionScore(api.LOLRegionBR1, "summonerID", 59)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
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
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: 0,
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
				Get(fmt.Sprintf(lol.ChampionMasteriesScoresURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.ChampionMasteries.MasteryScoreSum(api.LOLRegionBR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
