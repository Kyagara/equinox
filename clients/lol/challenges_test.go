package lol_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChallengesList(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.ChallengeConfigInfoDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]lol.ChallengeConfigInfoDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(lol.ChallengesConfigurationsURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Challenges.List(lol.BR1)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChallengesByID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.ChallengeConfigInfoDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.ChallengeConfigInfoDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.ChallengesConfigurationByIDURL, 1)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Challenges.ByID(lol.BR1, 1)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChallengesPercentiles(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *map[int64]lol.PercentileDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &map[int64]lol.PercentileDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(lol.ChallengesPercentilesURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Challenges.Percentiles(lol.BR1)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChallengesPercentilesByID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.PercentileDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.PercentileDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.ChallengesPercentileByIDURL, 1)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Challenges.PercentilesByID(lol.BR1, 1)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChallengesLeaderboards(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]lol.ApexPlayerInfoDTO
		wantErr error
		limit   int
	}{
		{
			name:  "found",
			code:  http.StatusOK,
			want:  &[]lol.ApexPlayerInfoDTO{},
			limit: 0,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			limit:   0,
		},
		{
			name:    "limit > 0",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			limit:   1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.ChallengesLeaderboardsByLevelURL, 1, lol.MasterLevel)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Challenges.Leaderboards(lol.BR1, 1, lol.MasterLevel, test.limit)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestChallengesByPUUID(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lol.PlayerInfoDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lol.PlayerInfoDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
				Get(fmt.Sprintf(lol.ChallengesByPUUIDURL, "PUUID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Challenges.ByPUUID(lol.BR1, "PUUID")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}
