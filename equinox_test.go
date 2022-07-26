package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/allegro/bigcache/v3"
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
			client, err := equinox.NewClient(test.key)

			if test.name != "success" {
				require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

				if test.wantErr == nil {
					require.Equal(t, test.want, client)
				}
			} else {
				require.NotNil(t, client, "expecting non-nil Client")
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
				Timeout:  15,
				Retry:    true,
			},
		},
		{
			name:    "cluster nil",
			wantErr: fmt.Errorf("cluster not provided"),
			config: &api.EquinoxConfig{
				Key:      "RGAPI-KEY",
				LogLevel: api.DebugLevel,
				Timeout:  15,
				Retry:    true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)

			if test.name != "success" {
				require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

				if test.wantErr == nil {
					require.Equal(t, test.want, client)
				}
			} else {
				require.NotNil(t, client, "expecting non-nil Client")
			}
		})
	}
}

func TestEquinoxClientClearCache(t *testing.T) {
	cache, err := cache.NewBigCache(bigcache.DefaultConfig(4 * time.Minute))

	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	config := internal.NewTestEquinoxConfig()

	config.Cache = cache

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(t, err, "expecting nil error")

	account := &riot.AccountDTO{
		PUUID:    "puuid",
		GameName: "gamename",
		TagLine:  "tagline",
	}

	delay := 2 * time.Second

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

	require.Equal(t, account, gotData, fmt.Sprintf("want data %v, got %v", account, gotData))

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	if duration >= 2 {
		t.Error(fmt.Errorf("request took more than 1s, took %ds, request not cached", duration))
	}

	client.ClearCache()

	start = time.Now()
	gotData, gotErr = client.Riot.Account.ByPUUID("puuid")
	duration = int(time.Since(start).Seconds())

	if duration <= 1 {
		t.Error(fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration))
	}

	require.Nil(t, gotData)

	assert.Equal(t, api.ForbiddenError, gotErr, fmt.Sprintf("want err %v, got %v", api.ForbiddenError, gotErr))
}
