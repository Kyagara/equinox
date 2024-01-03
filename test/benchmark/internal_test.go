package benchmark

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInternals-16 221295 5250 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 209899 5375 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 206281 5249 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 215037 5364 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 207663 5274 ns/op 1402 B/op 17 allocs/op
*/
func BenchmarkInternals(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()
	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}

	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
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
BenchmarkRequest-16 304810 3738 ns/op 664 B/op 8 allocs/op
BenchmarkRequest-16 308384 3693 ns/op 664 B/op 8 allocs/op
BenchmarkRequest-16 308928 3684 ns/op 664 B/op 8 allocs/op
BenchmarkRequest-16 301702 3741 ns/op 664 B/op 8 allocs/op
BenchmarkRequest-16 310485 3734 ns/op 664 B/op 8 allocs/op
*/
func BenchmarkRequest(b *testing.B) {
	b.ReportAllocs()

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()
	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}

	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
		require.NoError(b, err)
		// equinoxReq.Request.URL.String() causes 4 allocs
		require.Equal(b, "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid", equinoxReq.Request.URL.String())
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkExecute-16 241426 4647 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 239658 4719 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 244675 4652 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 247274 4616 ns/op 889 B/op 14 allocs/op
BenchmarkExecute-16 245326 4682 ns/op 889 B/op 14 allocs/op
*/
func BenchmarkExecute(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		require.NoError(b, err)
		require.Equal(b, "response", data)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkExecuteRaw-16 262534 4372 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 277652 4298 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 267313 4249 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 265563 4284 ns/op 1400 B/op 15 allocs/op
BenchmarkExecuteRaw-16 264864 4260 ns/op 1400 B/op 15 allocs/op
*/
func BenchmarkExecuteRaw(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client := internal.NewInternalClient(config)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/", "puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.ExecuteRaw(ctx, equinoxReq)
		require.NoError(b, err)
		// string(data) causes an allocation
		require.Equal(b, `"response"`, string(data))
	}
}

func BenchmarkURLWithAuthorizationHash(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()

	req := &http.Request{
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
		Header: http.Header{},
	}

	equinoxReq := api.EquinoxRequest{Request: req}
	equinoxReq.URL = req.URL.String()
	req.Header.Set("Authorization", "7267ee00-5696-47b8-9cae-8db3d49c8c33")

	for i := 0; i < b.N; i++ {
		hash, err := internal.GetURLWithAuthorizationHash(equinoxReq)
		require.NoError(b, err)
		require.Equal(b, "http://example.com/path-45da11db1ebd17ee0c32aca62e08923ea4f15590058ff1e15661bc13ed33df9d", hash)
	}
}
