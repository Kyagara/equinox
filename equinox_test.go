package equinox_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
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
			require.NotEmpty(t, client)
			require.NotEmpty(t, client.Cache)
			require.NotEmpty(t, client.LOL)
			require.NotEmpty(t, client.LOR)
			require.NotEmpty(t, client.TFT)
			require.NotEmpty(t, client.VAL)
			require.NotEmpty(t, client.Riot)
			require.NotEmpty(t, client.CDragon)
			require.NotEmpty(t, client.DDragon)
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
		config  api.EquinoxConfig
	}{
		{
			name:   "success",
			want:   &equinox.Equinox{},
			config: util.NewTestEquinoxConfig(),
		},
		{
			name:   "no cache",
			want:   &equinox.Equinox{},
			config: util.NewTestEquinoxConfig(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := equinox.NewClientWithConfig(test.config)

			if test.wantErr == nil {
				require.NotEmpty(t, client)
			}

			if test.name == "no cache" {
				require.Equal(t, client.Cache.TTL, time.Duration(0))
			}
		})
	}
}

func TestRateLimitWithMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	headers := map[string][]string{
		ratelimit.APP_RATE_LIMIT_HEADER:          {"5:3"},
		ratelimit.METHOD_RATE_LIMIT_HEADER:       {"5:3"},
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    {"1:3"},
		ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: {"1:3"},
	}

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, []byte(`{}`)).HeaderSet(headers))

	config := util.NewTestEquinoxConfig()
	config.RateLimit = ratelimit.NewInternalRateLimit(0, 0.5)
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client := equinox.NewClientWithConfig(config)

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
	_, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	// This last request (5) should block until rate limit is reset, this test should take around 3 seconds
	ctx = context.Background()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.NoError(t, err)
}
