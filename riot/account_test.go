package riot_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/riot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestAccountPlayerActiveShard(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := riot.NewRiotClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *riot.ActiveShardDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &riot.ActiveShardDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Get(fmt.Sprintf(riot.AccountActiveShardURL, api.VAL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Account.PlayerActiveShard("PUUID", api.VAL)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestAccountByPUUID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := riot.NewRiotClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *riot.AccountDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &riot.AccountDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Get(fmt.Sprintf(riot.AccountByPUUIDURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Account.ByPUUID("PUUID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestAccountByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := riot.NewRiotClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *riot.AccountDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &riot.AccountDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Get(fmt.Sprintf(riot.AccountByRiotIDURL, "gameName", "tagLine")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Account.ByID("gameName", "tagLine")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestAccountByAccessToken(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := riot.NewRiotClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *riot.AccountDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &riot.AccountDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Get(riot.AccountByAccessTokenURL).
				Reply(test.code).
				JSON(test.want).SetHeader("Authorization", "accessToken")

			gotData, gotErr := client.Account.ByAccessToken("accessToken")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
