package equinox_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/h2non/gock"
	"github.com/rs/zerolog"
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
			key:  "RGAPI-TEST",
		},
		{
			name: "nil key",
			key:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClient(test.key)
			require.Equal(t, test.wantErr, err)
			require.NotEmpty(t, client, "expecting non-nil Client")
			require.NotEmpty(t, client.Cache, "expecting non-nil Client")
			require.NotEmpty(t, client.LOL, "expecting nil Client")
			require.NotEmpty(t, client.LOR, "expecting nil Client")
			require.NotEmpty(t, client.TFT, "expecting nil Client")
			require.NotEmpty(t, client.VAL, "expecting nil Client")
			require.NotEmpty(t, client.Riot, "expecting nil Client")
			require.NotEmpty(t, client.CDragon, "expecting non-nil Client")
			require.NotEmpty(t, client.DDragon, "expecting non-nil Client")
		})
	}
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	emptyKeyConfig := util.NewTestEquinoxConfig()
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
			config: util.NewTestEquinoxConfig(),
		},
		{
			name:    "nil config",
			wantErr: internal.ErrConfigurationNotProvided,
			config:  nil,
		},
		{
			name:    "no cache",
			want:    &equinox.Equinox{},
			config:  util.NewTestEquinoxConfig(),
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)
			require.Equal(t, test.wantErr, err)

			if test.wantErr == nil {
				require.NotEmpty(t, client, "expecting non-nil client")
			}

			if test.name == "no cache" {
				require.Equal(t, client.Cache.TTL, time.Duration(0), "expecting cache disabled")
			}
		})
	}
}

func TestRateLimitWithMock(t *testing.T) {
	headers := map[string]string{
		ratelimit.APP_RATE_LIMIT_HEADER:    "5:3",
		ratelimit.METHOD_RATE_LIMIT_HEADER: "5:3",
	}

	// Mock 5 responses
	// The would be 5 request should be blocked from being created since it would exceed the rate limit
	for i := 1; i <= 5; i++ {
		headers[ratelimit.APP_RATE_LIMIT_COUNT_HEADER] = fmt.Sprintf("%d:3", i)
		headers[ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER] = fmt.Sprintf("%d:3", i)

		gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1, "/lol/summoner/v4/summoners/by-puuid/puuid")).
			Get("").
			Reply(200).
			SetHeaders(headers).
			JSON(&lol.SummonerV4DTO{})
	}

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(t, err)

	for i := 1; i <= 4; i++ {
		ctx := context.Background()
		_, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(t, err)
	}

	// One request left, it should be blocked as it would exceed the rate limit

	// A deadline is set, however, the bucket refill would take longer than the ctx deadline, waiting would exceed that deadline
	// It also shouldn't Take() a token from the bucket
	ctx := context.Background()
	ctx, c := context.WithTimeout(ctx, 2*time.Second)
	defer c()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	// This last request (5) should block until rate limit is reset, this test should take around 3 seconds
	ctx = context.Background()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.NoError(t, err)
}
