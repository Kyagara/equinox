package equinox_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/clients/lol"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/Kyagara/equinox/v2/ratelimit"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestNewEquinoxClient(t *testing.T) {
	client, err := equinox.NewClient("")
	require.Error(t, err)
	require.Empty(t, client)

	client, err = equinox.NewClient("RGAPI-TEST")
	require.NoError(t, err)

	require.NotEmpty(t, client)
	require.NotEmpty(t, client.Internal)
	require.NotEmpty(t, client.Cache)
	require.NotEmpty(t, client.RateLimit)
	require.NotEmpty(t, client.LOL)
	require.NotEmpty(t, client.LOR)
	require.NotEmpty(t, client.TFT)
	require.NotEmpty(t, client.VAL)
	require.NotEmpty(t, client.Riot)
}

func TestNewCustomClient(t *testing.T) {
	config := util.NewTestEquinoxConfig()
	config.Key = ""

	// Key not provided
	client, err := equinox.NewCustomClient(config, nil, nil, nil)
	require.Equal(t, internal.ErrKeyNotProvided, err)
	require.Empty(t, client)

	config.Key = "RGAPI-TEST"

	client, err = equinox.NewCustomClient(config, nil, &cache.Cache{TTL: 1}, &ratelimit.RateLimit{Enabled: true})
	require.NoError(t, err)
	require.Equal(t, client.Cache.TTL, time.Duration(1))
	require.True(t, client.RateLimit.Enabled)
}

func TestClientMethods(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, err := equinox.NewCustomClient(util.NewTestEquinoxConfig(), nil, nil, nil)
	require.NoError(t, err)

	// Using match.list as an example since its a []string
	httpmock.RegisterResponder("POST", "https://americas.api.riotgames.com/lol/tournament/v5/codes",
		httpmock.NewJsonResponderOrPanic(200, httpmock.File("./test/data/match.list.json")).Once())

	// Post with a body
	ctx := context.Background()
	codes, err := client.LOL.TournamentV5.CreateTournamentCode(ctx, api.AMERICAS, &lol.TournamentCodeParametersV5DTO{MapType: "SUMMONERS_RIFT"}, 420, 20)
	require.NoError(t, err)
	require.NotEmpty(t, codes)
	require.Len(t, codes, 20)

	// Put, also this endpoint method returns nothing. Reusing match.list result as an example, a valid tournamentCode should be used instead
	httpmock.RegisterResponder("PUT", "https://americas.api.riotgames.com/lol/tournament/v5/codes/KR_7050905124",
		httpmock.NewStringResponder(200, "").Once())

	err = client.LOL.TournamentV5.UpdateCode(ctx, api.AMERICAS, &lol.TournamentCodeUpdateParametersV5DTO{MapType: "HOWLING_ABYSS"}, codes[0])
	require.NoError(t, err)
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
	config.Retry = api.Retry{MaxRetries: 3}

	client, err := equinox.NewCustomClient(config, nil, nil, ratelimit.NewInternalRateLimit(0.99, time.Second))
	require.NoError(t, err)

	require.True(t, client.Internal.IsRateLimitEnabled)

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

	// This request should error out, it would block because it exceeds the rate limit but a context with a deadline is set
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	// This last request should block until rate limit is reset, this test should take around 4 seconds
	ctx = context.Background()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.NoError(t, err)
}
