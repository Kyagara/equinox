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
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

/*
The only thing that really matters here for benchmark purposes is B/op and allocs/op.

goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelRateLimit-16 100 100046964 ns/op 2817 B/op 31 allocs/op
BenchmarkParallelRateLimit-16 100 100046359 ns/op 2795 B/op 31 allocs/op
BenchmarkParallelRateLimit-16 100 100051542 ns/op 2965 B/op 31 allocs/op
BenchmarkParallelRateLimit-16 100 100048269 ns/op 2755 B/op 31 allocs/op
BenchmarkParallelRateLimit-16 100 100044937 ns/op 2770 B/op 31 allocs/op
*/
func BenchmarkParallelRateLimit(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")).HeaderSet(http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:          {"20:1,40:4"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    {"1:1,1:4"},
			ratelimit.METHOD_RATE_LIMIT_HEADER:       {"1300:60"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: {"1:60"},
		}))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	config.RateLimit = ratelimit.NewInternalRateLimit(0.99, time.Second)
	client := equinox.NewClientWithConfig(config)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			if err != nil {
				b.Fail()
			}
			if data.ProfileIconID != 1386 {
				b.Fail()
			}
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelSummonerByPUUID-16 297469 3876 ns/op 1507 B/op 17 allocs/op
BenchmarkParallelSummonerByPUUID-16 262147 3895 ns/op 1508 B/op 17 allocs/op
BenchmarkParallelSummonerByPUUID-16 295494 3960 ns/op 1508 B/op 17 allocs/op
BenchmarkParallelSummonerByPUUID-16 286360 3919 ns/op 1509 B/op 17 allocs/op
BenchmarkParallelSummonerByPUUID-16 260604 3953 ns/op 1508 B/op 17 allocs/op
*/
func BenchmarkParallelSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := equinox.NewClientWithConfig(config)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			if err != nil {
				b.Fail()
			}
			if data.ProfileIconID != 1386 {
				b.Fail()
			}
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelRedisCachedSummonerByPUUID-16 130348 8137 ns/op 1142 B/op 13 allocs/op
BenchmarkParallelRedisCachedSummonerByPUUID-16 130482 8131 ns/op 1147 B/op 13 allocs/op
BenchmarkParallelRedisCachedSummonerByPUUID-16 137911 8058 ns/op 1147 B/op 13 allocs/op
BenchmarkParallelRedisCachedSummonerByPUUID-16 136819 8066 ns/op 1147 B/op 13 allocs/op
BenchmarkParallelRedisCachedSummonerByPUUID-16 136263 8157 ns/op 1146 B/op 13 allocs/op
*/
func BenchmarkParallelRedisCachedSummonerByPUUID(b *testing.B) {
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

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			if err != nil {
				b.Fail()
			}
			if data.ProfileIconID != 1386 {
				b.Fail()
			}
		}
	})
}

/*
This function clones apiHeaders and adds a new Authorization header.

goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelSummonerByAccessToken-16 242644 4476 ns/op 2182 B/op 25 allocs/op
BenchmarkParallelSummonerByAccessToken-16 234915 4703 ns/op 2183 B/op 25 allocs/op
BenchmarkParallelSummonerByAccessToken-16 259096 4523 ns/op 2182 B/op 25 allocs/op
BenchmarkParallelSummonerByAccessToken-16 262804 4422 ns/op 2182 B/op 25 allocs/op
BenchmarkParallelSummonerByAccessToken-16 251022 4775 ns/op 2184 B/op 25 allocs/op
*/
func BenchmarkParallelSummonerByAccessToken(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/me",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/summoner.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := equinox.NewClientWithConfig(config)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByAccessToken(ctx, lol.BR1, "accesstoken")
			if err != nil {
				b.Fail()
			}
			if data.ProfileIconID != 1386 {
				b.Fail()
			}
		}
	})
}

/*
This function has multiple http queries.

goos: linux
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelMatchListByPUUID-16 240529 5072 ns/op 2802 B/op 33 allocs/op
BenchmarkParallelMatchListByPUUID-16 232051 4994 ns/op 2802 B/op 33 allocs/op
BenchmarkParallelMatchListByPUUID-16 234236 5013 ns/op 2803 B/op 33 allocs/op
BenchmarkParallelMatchListByPUUID-16 218185 5074 ns/op 2804 B/op 33 allocs/op
BenchmarkParallelMatchListByPUUID-16 222524 5005 ns/op 2805 B/op 33 allocs/op
*/
func BenchmarkParallelMatchListByPUUID(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://asia.api.riotgames.com/lol/match/v5/matches/by-puuid/puuid/ids?count=20&queue=420&type=ranked",
		httpmock.NewBytesResponder(200, util.ReadFile(b, "../data/match.list.json")))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := equinox.NewClientWithConfig(config)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.MatchV5.ListByPUUID(ctx, api.ASIA, "puuid", -1, -1, 420, "ranked", -1, 20)
			if err != nil {
				b.Fail()
			}
			if data[0] != "KR_6841523755" {
				b.Fail()
			}
		}
	})
}
