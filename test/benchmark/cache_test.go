package benchmark_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInMemoryCachedSummonerByPUUID-16 171974 6091 ns/op 2997 B/op 8 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 194630 6271 ns/op 2769 B/op 8 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 189826 6242 ns/op 2813 B/op 8 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 188211 6261 ns/op 2828 B/op 8 allocs/op
BenchmarkInMemoryCachedSummonerByPUUID-16 193594 6205 ns/op 2778 B/op 8 allocs/op
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
BenchmarkRedisCachedSummonerByPUUID-16 22797 53065 ns/op 1151 B/op 14 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 23408 52363 ns/op 1152 B/op 14 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 23143 49585 ns/op 1152 B/op 14 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 25165 52813 ns/op 1151 B/op 14 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 23169 51766 ns/op 1152 B/op 14 allocs/op
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

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	config.Cache = cache
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
BenchmarkSummonerByPUUID-16 143720 8108 ns/op 1514 B/op 18 allocs/op
BenchmarkSummonerByPUUID-16 145818 8021 ns/op 1514 B/op 18 allocs/op
BenchmarkSummonerByPUUID-16 144967 7982 ns/op 1514 B/op 18 allocs/op
BenchmarkSummonerByPUUID-16 140066 8181 ns/op 1514 B/op 18 allocs/op
BenchmarkSummonerByPUUID-16 145996 7994 ns/op 1514 B/op 18 allocs/op
*/
func BenchmarkSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := equinox.NewClientWithConfig(config)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.NoError(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}
