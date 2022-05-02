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

func TestSummonerByAccountID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.SummonerDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.SummonerDTO{},
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
				Get(fmt.Sprintf(lol.SummonerByAccountIDURL, "accountID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByAccountID(lol.BR1, "accountID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestSummonerByName(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.SummonerDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.SummonerDTO{},
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
				Get(fmt.Sprintf(lol.SummonerByNameURL, "summonerName")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByName(lol.BR1, "summonerName")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestSummonerByPUUID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.SummonerDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.SummonerDTO{},
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
				Get(fmt.Sprintf(lol.SummonerByPUUIDURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByPUUID(lol.BR1, "PUUID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestSummonerByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.SummonerDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.SummonerDTO{},
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
				Get(fmt.Sprintf(lol.SummonerByID, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByID(lol.BR1, "summonerID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
