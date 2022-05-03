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

func TestTournamentStubCreateCodes(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    []string
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: []string{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.RouteAmericas)).
				Post(lol.TournamentStubCodesURL).
				Reply(test.code).
				JSON(test.want)

			options := lol.TournamentCodeParametersDTO{
				MapType:       lol.SummonersRiftMap,
				PickType:      lol.TournamentDraftPick,
				SpectatorType: lol.AllSpectator,
				TeamSize:      5,
			}

			gotData, gotErr := client.TournamentStub.CreateCodes(1, 1, options)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentStubLobbyEvents(t *testing.T) {
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.RouteAmericas)).
				Get(fmt.Sprintf(lol.TournamentStubLobbyEventsURL, "tournamentCode")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.TournamentStub.LobbyEvents("tournamentCode")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentStubCreate(t *testing.T) {
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.RouteAmericas)).
				Post(lol.TournamentStubURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.TournamentStub.Create(1, "name")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestTournamentStubCreateProvider(t *testing.T) {
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, api.RouteAmericas)).
				Post(lol.TournamentStubProvidersURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.TournamentStub.CreateProvider("name", "http://localhost:80")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
