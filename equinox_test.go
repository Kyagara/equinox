package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/riot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestNewEquinoxClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		key     string
	}{
		{
			name: "success",
			want: &equinox.Equinox{},
			key:  "RGAPI-KEY",
		},
		{
			name:    "nil key",
			wantErr: fmt.Errorf("API Key not provided"),
			key:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotData, gotErr := equinox.NewClient(test.key)

			if test.name != "success" {
				require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

				if test.wantErr == nil {
					assert.Equal(t, test.want, gotData)
				}
			} else {
				require.NotEmpty(t, gotData, "expecting not empty client")
			}
		})
	}
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		config  *api.EquinoxConfig
	}{
		{
			name:   "success",
			want:   &equinox.Equinox{},
			config: internal.NewTestEquinoxConfig(),
		},
		{
			name:    "nil config",
			wantErr: fmt.Errorf("equinox configuration not provided"),
			config:  nil,
		},
		{
			name:    "api key nil",
			wantErr: fmt.Errorf("API Key not provided"),
			config: &api.EquinoxConfig{
				LogLevel: api.DebugLevel,
				Timeout:  10,
				Retry:    true,
			},
		},
		{
			name:    "cluster nil",
			wantErr: fmt.Errorf("cluster not provided"),
			config: &api.EquinoxConfig{
				Key:      "RGAPI-KEY",
				LogLevel: api.DebugLevel,
				Timeout:  10,
				Retry:    true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotData, gotErr := equinox.NewClientWithConfig(test.config)

			if test.name != "success" {
				require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

				if test.wantErr == nil {
					assert.Equal(t, test.want, gotData)
				}
			} else {
				require.NotEmpty(t, gotData, "expecting not empty client")
			}
		})
	}
}

func TestEquinoxClientClearCache(t *testing.T) {
	config := &api.EquinoxConfig{
		Key:       "RGAPI-KEY",
		Cluster:   api.AmericasCluster,
		LogLevel:  api.DebugLevel,
		Timeout:   10,
		TTL:       120,
		Retry:     false,
		RateLimit: false,
	}

	client, _ := equinox.NewClientWithConfig(config)

	delay := 2 * time.Second

	account := &riot.AccountDTO{
		PUUID:    "puuid",
		GameName: "gamename",
		TagLine:  "tagline",
	}

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
		Get(fmt.Sprintf(riot.AccountByPUUIDURL, "puuid")).
		Reply(200).
		JSON(account).Delay(delay)

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
		Get(fmt.Sprintf(riot.AccountByPUUIDURL, "puuid")).
		Reply(403).
		JSON(account).Delay(delay)

	gotData, gotErr := client.Riot.Account.ByPUUID("puuid")

	require.Equal(t, account, gotData)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start := time.Now()
	gotData, gotErr = client.Riot.Account.ByPUUID("puuid")
	duration := int(time.Since(start).Seconds())

	require.Equal(t, account, gotData)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	if duration >= 2 {
		gotErr = fmt.Errorf("request took more than 1s, took %ds, request not cached", duration)

		require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))
	}

	client.ClearCache()

	start = time.Now()
	gotData, gotErr = client.Riot.Account.ByPUUID("puuid")
	duration = int(time.Since(start).Seconds())

	require.Equal(t, api.ForbiddenError, gotErr, fmt.Sprintf("want err %v, got %v", api.ForbiddenError, gotErr))

	if duration <= 1 {
		gotErr = fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration)

		require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))
	}

	if gotErr == nil {
		assert.Equal(t, account, gotData)
	}
}
