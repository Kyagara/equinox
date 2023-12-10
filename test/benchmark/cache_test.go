package benchmark_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInMemoryCachedSummonerByPUUID-16 135530 7820 ns/op 3763 B/op 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 151477 7883 ns/op 3502 B/op 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 131240 8001 ns/op 3845 B/op 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 129679 7736 ns/op 3876 B/op 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 131004 7794 ns/op 3849 B/op 17 allocs/op
*/
func BenchmarkInMemoryCachedSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	client, err := equinox.NewClient("RGAPI-TEST")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkRedisCachedSummonerByPUUID-16 19467 62728 ns/op 1398 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 18762 64878 ns/op 1399 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 16803 65691 ns/op 1400 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 18232 62038 ns/op 1400 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 18590 62258 ns/op 1400 B/op 23 allocs/op
*/
func BenchmarkRedisCachedSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	ctx := context.Background()
	redisConfig := &redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}
	cache, err := cache.NewRedis(ctx, redisConfig, 4*time.Minute)
	require.NoError(b, err)
	config := api.EquinoxConfig{
		Key:      "RGAPI-TEST",
		LogLevel: zerolog.WarnLevel,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Retries: 3,
		Cache:   cache,
	}
	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkSummonerByPUUID-16 121406  9877 ns/op 1739 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 115393 10082 ns/op 1740 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 112934 10016 ns/op 1740 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 116364 10005 ns/op 1739 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 118736 10053 ns/op 1739 B/op 27 allocs/op
*/
func BenchmarkSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}
