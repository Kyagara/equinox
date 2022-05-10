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

func TestMatchList(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name       string
		code       int
		want       *[]string
		wantErr    error
		parameters *lol.MatchlistOptions
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]string{},
			parameters: &lol.MatchlistOptions{Start: 0,
				Count: 20},
		},
		{
			name:       "nil options",
			code:       http.StatusOK,
			want:       &[]string{},
			parameters: nil,
		},
		{
			name: "count > 100",
			code: http.StatusOK,
			want: &[]string{},
			parameters: &lol.MatchlistOptions{Start: 0,
				Count: 101},
		},
		{
			name: "all optional fields set",
			code: http.StatusOK,
			want: &[]string{},
			parameters: &lol.MatchlistOptions{
				StartTime: 1,
				EndTime:   1,
				Queue:     420,
				Type:      lol.RankedMatch,
				Start:     1,
				Count:     1,
			},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
			parameters: &lol.MatchlistOptions{Start: 0,
				Count: 20},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.Americas)).
				Get(fmt.Sprintf(lol.MatchListURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.List(lol.Americas, "PUUID", test.parameters)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.Americas)).
				Get(fmt.Sprintf(lol.MatchByIDURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.ByID(lol.Americas, "matchID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.Americas)).
				Get(fmt.Sprintf(lol.MatchTimelineURL, "matchID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Match.Timeline(lol.Americas, "matchID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
