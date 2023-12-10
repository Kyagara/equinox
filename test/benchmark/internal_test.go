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
goos: linux
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInternals-16 170773 6469 ns/op 1562 B/op 22 allocs/op
BenchmarkInternals-16 164850 6328 ns/op 1562 B/op 22 allocs/op
BenchmarkInternals-16 180541 6293 ns/op 1562 B/op 22 allocs/op
BenchmarkInternals-16 172447 6285 ns/op 1562 B/op 22 allocs/op
BenchmarkInternals-16 177984 6503 ns/op 1562 B/op 22 allocs/op
*/
func BenchmarkInternals(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewBytesResponder(200, []byte(`{}`)))

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	client, err := internal.NewInternalClient(config)
	require.NoError(b, err)

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, "br1", "/lol/summoner/v4/summoners/by-puuid/puuid", "summoner-v4.getByPUUID", nil)
		require.NoError(b, err)
		var data struct{}
		err = client.Execute(ctx, equinoxReq, &data)
		require.NoError(b, err)
	}
}
