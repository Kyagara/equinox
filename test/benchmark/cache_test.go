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
BenchmarkSummonerByPUUID-16 203803 5621 ns/op 1498 B/op 17 allocs/op
BenchmarkSummonerByPUUID-16 210289 5591 ns/op 1498 B/op 17 allocs/op
BenchmarkSummonerByPUUID-16 203768 5672 ns/op 1498 B/op 17 allocs/op
BenchmarkSummonerByPUUID-16 206062 5780 ns/op 1498 B/op 17 allocs/op
BenchmarkSummonerByPUUID-16 202708 5710 ns/op 1498 B/op 17 allocs/op
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
	client, err := equinox.NewClientWithConfig(config)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		if err != nil {
			b.Fatal(err)
		}
		if data.ProfileIconID != 1386 {
			b.Fatalf("ProfileIconID != 1386, got %d", data.ProfileIconID)
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInternalCachedSummonerByPUUID-16 316039 3766 ns/op 1024 B/op 7 allocs/op
BenchmarkInternalCachedSummonerByPUUID-16 325017 3594 ns/op 1024 B/op 7 allocs/op
BenchmarkInternalCachedSummonerByPUUID-16 312504 3684 ns/op 1024 B/op 7 allocs/op
BenchmarkInternalCachedSummonerByPUUID-16 311102 3653 ns/op 1024 B/op 7 allocs/op
BenchmarkInternalCachedSummonerByPUUID-16 327607 3721 ns/op 1024 B/op 7 allocs/op
*/
func BenchmarkInternalCachedSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	client, err := equinox.NewClient("RGAPI-TEST")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		if err != nil {
			b.Fatal(err)
		}
		if data.ProfileIconID != 1386 {
			b.Fatalf("ProfileIconID != 1386, got %d", data.ProfileIconID)
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkRedisCachedSummonerByPUUID-16 26541 47643 ns/op 1135 B/op 13 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 27054 46633 ns/op 1136 B/op 13 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 25658 45578 ns/op 1135 B/op 13 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 26203 45947 ns/op 1135 B/op 13 allocs/op
BenchmarkRedisCachedSummonerByPUUID-16 25495 47976 ns/op 1135 B/op 13 allocs/op
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
	client, err := equinox.NewClientWithConfig(config)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		if err != nil {
			b.Fatal(err)
		}
		if data.ProfileIconID != 1386 {
			b.Fatalf("ProfileIconID != 1386, got %d", data.ProfileIconID)
		}
	}
}
