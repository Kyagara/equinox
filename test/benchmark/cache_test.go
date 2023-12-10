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
BenchmarkInMemoryCachedSummonerByPUUID-16 133592 7960 ns/op 3799 B/o 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 127503 7960 ns/op 3920 B/o 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 149803 8081 ns/op 3527 B/o 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 149905 8179 ns/op 3525 B/o 17 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 144307 8141 ns/op 3612 B/o 17 allocs/op
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
BenchmarkRedisCachedSummonerByPUUID-16 17214 68915 ns/op 1393 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 19551 63823 ns/op 1393 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 20432 60437 ns/op 1394 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 20166 59177 ns/op 1393 B/op 23 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 19743 59580 ns/op 1393 B/op 23 allocs/op
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
BenchmarkSummonerByPUUID-16 122088 9837 ns/op 1752 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 124603 9537 ns/op 1752 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 125160 9726 ns/op 1752 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 124734 9774 ns/op 1752 B/op 27 allocs/op
BenchmarkSummonerByPUUID-16 115701 9854 ns/op 1752 B/op 27 allocs/op
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
