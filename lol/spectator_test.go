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

func TestSpectatorFeaturedGames(t *testing.T) {
	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.LOLRegionBR1)).
		Get(lol.SpectatorURL).
		Reply(200).
		JSON(&lol.FeaturedGamesDTO{})

	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	res, err := client.Spectator.FeaturedGames(api.LOLRegionBR1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}

func TestSpectatorCurrentGame(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.CurrentGameInfoDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.CurrentGameInfoDTO{},
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
				Get(fmt.Sprintf(lol.SpectatorCurrentGameURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Spectator.CurrentGame(api.LOLRegionBR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
