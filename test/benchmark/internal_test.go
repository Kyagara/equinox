package benchmark

import (
	"context"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInternals-16 208984 5425 ns/op 1419 B/op 18 allocs/op
BenchmarkInternals-16 216291 5445 ns/op 1419 B/op 18 allocs/op
BenchmarkInternals-16 208576 5400 ns/op 1419 B/op 18 allocs/op
BenchmarkInternals-16 206632 5434 ns/op 1419 B/op 18 allocs/op
BenchmarkInternals-16 186727 5425 ns/op 1419 B/op 18 allocs/op
*/
func BenchmarkInternals(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "br1", "/lol/summoner/v4/summoners/by-puuid/puuid", "summoner-v4.getByPUUID", nil)
		require.NoError(b, err)
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		require.NoError(b, err)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkRequest-16 301332 3952 ns/op 680 B/op 9 allocs/op
BenchmarkRequest-16 293253 3907 ns/op 680 B/op 9 allocs/op
BenchmarkRequest-16 292212 3997 ns/op 680 B/op 9 allocs/op
BenchmarkRequest-16 271179 4007 ns/op 680 B/op 9 allocs/op
BenchmarkRequest-16 297822 3870 ns/op 680 B/op 9 allocs/op
*/
func BenchmarkRequest(b *testing.B) {
	b.ReportAllocs()

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "br1", "/lol/summoner/v4/summoners/by-puuid/puuid", "summoner-v4.getByPUUID", nil)
		require.NoError(b, err)
		require.Equal(b, "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid", equinoxReq.Request.URL.String())
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkExecute-16 239575 4656 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 252027 4611 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 230380 4605 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 242892 4669 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 248352 4640 ns/op 889 B/op 14 allocs/op
*/
func BenchmarkExecute(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	equinoxReq, err := client.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "br1", "/lol/summoner/v4/summoners/by-puuid/puuid", "summoner-v4.getByPUUID", nil)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		require.NoError(b, err)
		require.Equal(b, `response`, data)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkExecuteRaw-16 269066 4372 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 252100 4355 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 253398 4314 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 271502 4340 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 266028 4332 ns/op 1400 B/op 15 allocs/op
*/
func BenchmarkExecuteRaw(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	equinoxReq, err := client.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "br1", "/lol/summoner/v4/summoners/by-puuid/puuid", "summoner-v4.getByPUUID", nil)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.ExecuteRaw(ctx, equinoxReq)
		require.NoError(b, err)
		require.Equal(b, `"response"`, string(data))
	}
}
