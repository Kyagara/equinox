package equinox_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/ratelimit"
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
	emptyKeyConfig := equinox.NewTestEquinoxConfig()
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
			config: equinox.NewTestEquinoxConfig(),
		},
		{
			name:    "nil config",
			wantErr: fmt.Errorf("equinox configuration not provided"),
			config:  nil,
		},
		{
			name:    "no cache",
			want:    &equinox.Equinox{},
			config:  equinox.NewTestEquinoxConfig(),
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

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
		ratelimit.APP_RATE_LIMIT_HEADER:    "5:5",
		ratelimit.METHOD_RATE_LIMIT_HEADER: "5:5",
	}

	// Mock 5 responses
	// The would be 5 request should be blocked from being created since it would exceed the rate limit
	for i := 1; i <= 5; i++ {
		headers[ratelimit.APP_RATE_LIMIT_COUNT_HEADER] = fmt.Sprintf("%d:5", i)
		headers[ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER] = fmt.Sprintf("%d:5", i)

		gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
			Get("/lol/summoner/v4/summoners/by-puuid/puuid").
			Reply(200).
			SetHeaders(headers).
			JSON(&lol.SummonerV4DTO{})
	}

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(t, err)

	for i := 1; i <= 4; i++ {
		ctx := context.Background()
		_, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.Nil(t, err)
	}

	// One request left, it should be blocked as it would exceed the rate limit

	// A deadline is set, however, the bucket refill would take longer than the ctx deadline, waiting would exceed the deadline
	// It also shouldn't add to the bucket count
	ctx := context.Background()
	ctx, c := context.WithTimeout(ctx, 2*time.Second)
	defer c()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)

	// This last request (5) should block until rate limit is reset
	now := time.Now()
	ctx = context.Background()
	_, err = client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
	require.Nil(t, err)
	require.GreaterOrEqual(t, time.Since(now), 5*time.Second)
}
