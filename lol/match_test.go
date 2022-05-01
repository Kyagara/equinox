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

func TestMatchlist(t *testing.T) {
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
				Get(fmt.Sprintf(lol.MatchlistURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			options := lol.MatchlistOptions{
				Start: 0,
				Count: 20,
			}

			gotData, gotErr := client.Match.ListByPUUID(api.RouteAmericas, "PUUID", &options)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestMatchByID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.MatchDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.MatchDTO{},
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
				Get(fmt.Sprintf(lol.MatchURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.ByID(api.RouteAmericas, "matchID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

func TestMatchTimeline(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.MatchTimelineDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.MatchTimelineDTO{},
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
				Get(fmt.Sprintf(lol.MatchTimelineURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.Timeline(api.RouteAmericas, "matchID")

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
