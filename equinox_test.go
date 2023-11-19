package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
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
			key:  "RGAPI-TEST",
		},
		{
			name: "nil key",
			key:  "",
			want: &equinox.Equinox{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClient(test.key)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))
			require.NotNil(t, client, "expecting non-nil Client")
			require.NotNil(t, client.Cache, "expecting non-nil Client")
			require.NotNil(t, client.LOL, "expecting nil Client")
			require.NotNil(t, client.LOR, "expecting nil Client")
			require.NotNil(t, client.TFT, "expecting nil Client")
			require.NotNil(t, client.VAL, "expecting nil Client")
			require.NotNil(t, client.Riot, "expecting nil Client")
			require.NotNil(t, client.CDragon, "expecting non-nil Client")
			require.NotNil(t, client.DDragon, "expecting non-nil Client")
		})
	}
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	emptyKeyConfig := internal.NewTestEquinoxConfig()
	emptyKeyConfig.Key = ""
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
			name:    "no cache",
			want:    &equinox.Equinox{},
			config:  internal.NewTestEquinoxConfig(),
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

			if test.wantErr == nil {
				require.NotNil(t, client, "expecting non-nil client")
			}

			if test.name == "no cache" {
				require.Equal(t, client.Cache.TTL, time.Duration(0), "expecting cache disabled")
			}
		})
	}
}

func TestEquinoxClientClearCache(t *testing.T) {
	config, err := equinox.DefaultConfig("RGAPI-TEST")
	config.LogLevel = api.NOP_LOG_LEVEL
	config.Retry = false
	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, client.Cache.TTL, 4*time.Minute, "expecting cache enabled")

	account := &riot.AccountV1DTO{
		PUUID:    "puuid",
		GameName: "gamename",
		TagLine:  "tagline",
	}

	delay := 2 * time.Second

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(200).
		JSON(account).Delay(delay)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(404).
		JSON(account).Delay(delay)

	gotData, gotErr := client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	require.Equal(t, account, gotData)
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start := time.Now()
	gotData, gotErr = client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	duration := int(time.Since(start).Seconds())

	require.Equal(t, account, gotData, fmt.Sprintf("want data %v, got %v", account, gotData))
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	if duration >= 2 {
		t.Error(fmt.Errorf("request took more than 1s, took %ds, request not cached", duration))
	}

	gotErr = client.Cache.Clear()
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start = time.Now()
	gotData, gotErr = client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	duration = int(time.Since(start).Seconds())

	if duration <= 1 {
		t.Error(fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration))
	}

	require.Nil(t, gotData)
	require.Equal(t, api.ErrNotFound, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrNotFound, gotErr))
}

func TestPutClient(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := lol.NewLOLClient(internalClient)
	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Put("/lol/tournament/v5/codes/tournamentCode").
		Reply(200)

	err = client.TournamentV5.UpdateCode(api.AMERICAS, nil, "tournamentCode")
	require.Nil(t, err, "expecting nil error")
}
