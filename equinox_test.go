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
	"github.com/stretchr/testify/require"
)

func TestNewEquinoxClient(t *testing.T) {
	_, err := equinox.NewClient("")
	require.Error(t, err)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.NoError(t, err)

	require.NotEmpty(t, client.Cache)
	require.NotEmpty(t, client.LOL)
	require.NotEmpty(t, client.LOR)
	require.NotEmpty(t, client.TFT)
	require.NotEmpty(t, client.VAL)
	require.NotEmpty(t, client.Riot)
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	config := util.NewTestEquinoxConfig()
	config.Key = ""
	_, err := equinox.NewClientWithConfig(config)
	require.Error(t, err)

	config.Key = "RGAPI-TEST"
	config.Cache.TTL = 0
	client, err := equinox.NewClientWithConfig(config)
	require.NoError(t, err)
	require.Equal(t, client.Cache.TTL, time.Duration(0))
}

func TestRateLimitWithMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	headers := map[string][]string{
		ratelimit.APP_RATE_LIMIT_HEADER:          {"6:3"},
		ratelimit.METHOD_RATE_LIMIT_HEADER:       {"6:3"},
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    {"1:3"},
		ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: {"1:3"},
	}

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, []byte(`{}`)).HeaderSet(headers))

	config := util.NewTestEquinoxConfig()
	config.RateLimit = ratelimit.NewInternalRateLimit(0.99, time.Second)
	config.Retry = api.Retry{MaxRetries: 3}

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(t, err)

	for i := 1; i <= 4; i++ {
		ctx := context.Background()
		_, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(t, err)
	}

	// One request left, it should be blocked as it would exceed the rate limit

	// A deadline is set, however, the bucket refill would take longer than the ctx deadline, waiting would exceed that deadline
	// It also shouldn't Reserve() a request from the bucket
	ctx := context.Background()
	ctx, c := context.WithTimeout(ctx, 2*time.Second)
	defer c()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	// This last request should block until rate limit is reset, this test should take around 3 seconds
	ctx = context.Background()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.NoError(t, err)
}
