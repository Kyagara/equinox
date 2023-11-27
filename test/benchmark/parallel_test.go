package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

// These parallel benchmarks are more of a test less of an benchmark.
// Used to see how the library reacts to some parallel work.

// This revealed problems with multiple limits in a bucket, only the first bucket was being respected.
func BenchmarkParallelRateLimit(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).SetHeaders(map[string]string{
		ratelimit.APP_RATE_LIMIT_HEADER:       "20:1,50:3,80:5",
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER: "1:1,1:3,1:5",
	}).
		JSON(summoner)

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.Nil(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelCachedSummonerByPUUID-16 254956 4763 ns/op 2762 B/op 20 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 245898 4688 ns/op 2810 B/op 20 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 255544 4860 ns/op 2759 B/op 20 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 250755 4800 ns/op 2784 B/op 20 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 235132 5019 ns/op 2873 B/op 20 allocs/op
*/
func BenchmarkParallelCachedSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).
		JSON(summoner)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.Nil(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.Nil(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelSummonerByPUUID-16 138342  7883 ns/op 3005 B/op 41 allocs/op
BenchmarkParallelSummonerByPUUID-16 119114  9325 ns/op 3134 B/op 42 allocs/op
BenchmarkParallelSummonerByPUUID-16 111822 10600 ns/op 3396 B/op 43 allocs/op
BenchmarkParallelSummonerByPUUID-16 101041 11860 ns/op 3398 B/op 43 allocs/op
BenchmarkParallelSummonerByPUUID-16  96073 12417 ns/op 3915 B/op 44 allocs/op
*/
func BenchmarkParallelSummonerByPUUID(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).
		JSON(summoner)

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.Nil(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}
