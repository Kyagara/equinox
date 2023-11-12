package cdragon_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/cdragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestChampionByName(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := cdragon.NewCDragonClient(internalClient)
	data := &cdragon.ChampionData{Name: "Jarvan IV"}

	tests := []struct {
		name    string
		code    int
		want    *cdragon.ChampionData
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: data,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.C_DRAGON_BASE_URL_FORMAT, "")).
				Get(fmt.Sprintf(cdragon.ChampionURL, "1.0", "JarvanIV")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Champion.ByName("1.0", "JarvanIV")
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
			if test.wantErr == nil {
				require.Equal(t, data, gotData)
			}
		})
	}
}

func TestChampionByID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := cdragon.NewCDragonClient(internalClient)
	data := &cdragon.ChampionData{Name: "Jarvan IV"}

	tests := []struct {
		name    string
		code    int
		want    *cdragon.ChampionData
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: data,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.C_DRAGON_BASE_URL_FORMAT, "")).
				Get(fmt.Sprintf(cdragon.ChampionURL, "1.0", 223)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Champion.ByID("1.0", 223)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
			if test.wantErr == nil {
				require.Equal(t, data, gotData)
			}
		})
	}
}
