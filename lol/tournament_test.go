package lol_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestTournamentCreateCodes(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name       string
		code       int
		want       []string
		wantErr    error
		count      int
		parameters *lol.TournamentCodeParametersDTO
	}{
		{
			name:  "found",
			code:  http.StatusOK,
			want:  []string{},
			count: 1,
			parameters: &lol.TournamentCodeParametersDTO{
				MapType:       lol.SummonersRiftMap,
				PickType:      lol.TournamentDraftPick,
				SpectatorType: lol.AllSpectator,
				TeamSize:      5,
			},
		},
		{
			name:    "not found",
			count:   1,
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
			parameters: &lol.TournamentCodeParametersDTO{
				MapType:       lol.SummonersRiftMap,
				PickType:      lol.TournamentDraftPick,
				SpectatorType: lol.AllSpectator,
				TeamSize:      5,
			},
		},
		{
			name:       "count < 0",
			count:      -1,
			code:       http.StatusNotFound,
			wantErr:    fmt.Errorf("count can't be less than 1 or more than 1000"),
			parameters: nil,
		},
		{
			name:       "parameters is nil",
			count:      1,
			code:       http.StatusNotFound,
			wantErr:    fmt.Errorf("parameters are required"),
			parameters: nil,
		},
		{
			name:       "parameters with default value",
			count:      1,
			code:       http.StatusNotFound,
			wantErr:    fmt.Errorf("required values are empty"),
			parameters: &lol.TournamentCodeParametersDTO{},
		},
		{
			name:    "options with invalid team size",
			count:   1,
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("invalid team size: 0, valid values are 1-5"),
			parameters: &lol.TournamentCodeParametersDTO{TeamSize: 0, MapType: lol.SummonersRiftMap,
				PickType:      lol.TournamentDraftPick,
				SpectatorType: lol.AllSpectator},
		},
		{
			name:       "options with only one value set",
			count:      1,
			code:       http.StatusNotFound,
			wantErr:    fmt.Errorf("not all required values are set"),
			parameters: &lol.TournamentCodeParametersDTO{MapType: lol.SummonersRiftMap},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Post(lol.TournamentCodesURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.CreateCodes(1, test.count, test.parameters)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentByCode(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.TournamentCodeDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.TournamentCodeDTO{},
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
				Get(fmt.Sprintf(lol.TournamentByCodeURL, "tournamentCode")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.ByCode("tournamentCode")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentUpdate(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name       string
		code       int
		parameters *lol.TournamentCodeUpdateParametersDTO
		wantErr    error
	}{
		{
			name:       "found",
			code:       http.StatusOK,
			parameters: &lol.TournamentCodeUpdateParametersDTO{},
		},
		{
			name:       "not found",
			code:       http.StatusNotFound,
			wantErr:    api.NotFoundError,
			parameters: &lol.TournamentCodeUpdateParametersDTO{},
		},
		{
			name:       "parameters is nil",
			code:       http.StatusOK,
			wantErr:    fmt.Errorf("parameters are required"),
			parameters: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Put(fmt.Sprintf(lol.TournamentByCodeURL, "tournamentCode")).
				Reply(test.code)

			gotErr := client.Tournament.Update("tournamentCode", test.parameters)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotErr, nil)
			}
		})
	}
}

func TestTournamentLobbyEvents(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.LobbyEventDTOWrapper
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.LobbyEventDTOWrapper{},
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
				Get(fmt.Sprintf(lol.TournamentLobbyEventsURL, "tournamentCode")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.LobbyEvents("tournamentCode")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentCreate(t *testing.T) {
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Post(lol.TournamentURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.Create(1, "name")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentCreateProvider(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    int
		wantErr error
		url     string
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: 0,
			url:  "http://localhost:80",
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
			url:     "http://localhost:80",
		},
		{
			name: "invalid url",
			code: http.StatusOK,
			wantErr: &url.Error{
				Op:  "parse",
				URL: "invalidurl",
				Err: fmt.Errorf("invalid URI for request"),
			},
			url: "invalidurl",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.Americas)).
				Post(lol.TournamentProvidersURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.CreateProvider("name", test.url)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}