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
BenchmarkInMemoryCachedSummonerByPUUID-16 158094 6776 ns/op 3265 B/op 12 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 142179 7192 ns/op 3503 B/op 12 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 149342 6798 ns/op 3389 B/op 12 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 151874 6848 ns/op 3352 B/op 12 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 173299 6760 ns/op 3078 B/op 12 allocs/op
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
BenchmarkRedisCachedSummonerByPUUID-16 23100 53341 ns/op 1245 B/op 18 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 22167 52512 ns/op 1245 B/op 18 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 22860 52094 ns/op 1245 B/op 18 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 22894 53063 ns/op 1245 B/op 18 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 21548 53368 ns/op 1245 B/op 18 allocs/op
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
	client := equinox.NewClientWithConfig(config)

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
BenchmarkSummonerByPUUID-16 137259 8428 ns/op 1608 B/op 22 allocs/op
BenchmarkSummonerByPUUID-16 141212 8332 ns/op 1608 B/op 22 allocs/op
BenchmarkSummonerByPUUID-16 134893 8377 ns/op 1608 B/op 22 allocs/op
BenchmarkSummonerByPUUID-16 137270 8503 ns/op 1608 B/op 22 allocs/op
BenchmarkSummonerByPUUID-16 141310 8387 ns/op 1608 B/op 22 allocs/op
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

	client := equinox.NewClientWithConfig(config)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}
