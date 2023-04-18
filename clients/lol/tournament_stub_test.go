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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTournamentStubCreateCodes(t *testing.T) {
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
			name:    "parameters with invalid team size",
			count:   1,
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("invalid team size: 0, valid values are 1-5"),
			parameters: &lol.TournamentCodeParametersDTO{TeamSize: 0, MapType: lol.SummonersRiftMap,
				PickType:      lol.TournamentDraftPick,
				SpectatorType: lol.AllSpectator},
		},
		{
			name:       "parameters with only one value set",
			count:      1,
			code:       http.StatusNotFound,
			wantErr:    fmt.Errorf("not all required values are set"),
			parameters: &lol.TournamentCodeParametersDTO{MapType: lol.SummonersRiftMap},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
				Post(lol.TournamentStubCodesURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.TournamentStub.CreateCodes(1, test.count, test.parameters)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestTournamentStubLobbyEvents(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(lol.LobbyEventDTOWrapper{}, &lol.LobbyEventDTOWrapper{})

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := fmt.Sprintf(lol.TournamentStubLobbyEventsURL, "tournamentCode")
			test.MockGetResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.TournamentStub.LobbyEvents("tournamentCode")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestTournamentStubCreate(t *testing.T) {
	client, err := test.TestingNewLOLClient()

	require.Nil(t, err, "expecting nil error")

	tests := test.GetEndpointTestCases(*new(int), new(int))

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			url := lol.TournamentStubURL
			test.MockPostResponse(url, string(api.AmericasCluster), test.AccessToken)
			gotData, gotErr := client.TournamentStub.Create(1, "name")
			test.CheckResponse(t, gotData, gotErr)
		})
	}
}

func TestTournamentStubCreateProvider(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *int
		wantErr error
		url     string
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: new(int),
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
				Post(lol.TournamentStubProvidersURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.TournamentStub.CreateProvider("name", test.url)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
