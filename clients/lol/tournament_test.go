package lol_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestTournamentCreateCodes(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name       string
		code       int
		want       *[]string
		wantErr    error
		count      int
		parameters *lol.TournamentCodeParametersDTO
	}{
		{
			name:  "found",
			code:  http.StatusOK,
			want:  &[]string{},
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
			wantErr: api.ErrNotFound,
			parameters: &lol.TournamentCodeParametersDTO{
				MapType:       lol.SummonersRiftMap,
				PickType:      lol.TournamentDraftPick,
				SpectatorType: lol.AllSpectator,
				TeamSize:      5,
			},
		},
		{
			name:       "default values",
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
			gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
				Post(lol.TournamentCodesURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.CreateCodes(1, test.count, test.parameters)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				require.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestTournamentByCode(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.TournamentCodeDTO{}, &lol.TournamentCodeDTO{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.TournamentByCodeURL, "tournamentCode")
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.Tournament.ByCode("tournamentCode")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestTournamentUpdate(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

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
			wantErr:    api.ErrNotFound,
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
			gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
				Put(fmt.Sprintf(lol.TournamentByCodeURL, "tournamentCode")).
				Reply(test.code)

			gotErr := client.Tournament.Update("tournamentCode", test.parameters)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				require.Equal(t, nil, gotErr)
			}
		})
	}
}

func TestTournamentLobbyEvents(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.LobbyEventDTOWrapper{}, &lol.LobbyEventDTOWrapper{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.TournamentLobbyEventsURL, "tournamentCode")
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.Tournament.LobbyEvents("tournamentCode")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestTournamentCreate(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := []struct {
		name    string
		want    int
		code    int
		wantErr error
	}{
		{
			name: "found",
			code: 200,
			want: 0,
		}, {
			name:    "found",
			code:    404,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
				Post(lol.TournamentURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.Create(1, "name")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				require.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestTournamentCreateProvider(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

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
			wantErr: api.ErrNotFound,
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
			gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
				Post(lol.TournamentProvidersURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Tournament.CreateProvider("name", test.url)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				require.Equal(t, test.want, gotData)
			}
		})
	}
}
