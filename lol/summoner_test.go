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
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.SummonerByAccountIDURL, "accountID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByAccountID(lol.BR1, "accountID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.SummonerByNameURL, "summonerName")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByName(lol.BR1, "summonerName")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.SummonerByPUUIDURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByPUUID(lol.BR1, "PUUID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.SummonerByIDURL, "summonerID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Summoner.ByID(lol.BR1, "summonerID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestSummonerByAccessToken(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name        string
		code        int
		want        *lol.SummonerDTO
		wantErr     error
		accessToken string
	}{
		{
			name:        "found",
			code:        http.StatusOK,
			want:        &lol.SummonerDTO{},
			accessToken: "accessToken",
		},
		{
			name:        "not found",
			code:        http.StatusNotFound,
			wantErr:     api.NotFoundError,
			accessToken: "accessToken",
		},
		{
			name:    "accessToken empty",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("accessToken is required"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(lol.SummonerByAccessTokenURL).
				Reply(test.code).
				JSON(test.want).SetHeader("Authorization", "accessToken")

			gotData, gotErr := client.Summoner.ByAccessToken(lol.BR1, test.accessToken)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
