# Tests

## Benchmark

> [!WARNING]
> Benchmarks are not automated, take the results in the code comments with a grain of salt.

Keep in mind that since requests are mocked using `httpmock`, results (time, bytes, allocs) will be different from production, specially since you can't make 300000 requests in 1 second to the Riot Games API.

Benchmarks are separated in four files: parallel, data, cache and internal.

- parallel: Variety of benchmarks to test parallelism. Used more as tests to check for race conditions.
- cache: Cache benchmarks for both BigCache and Redis.
- data: Benchmarks that use data from the live Riot Games API.
- internal: Focused on InternalClient functions.

Benchmarks clients should be using a configuration close to the one used in production.

Some benchmarks are mainly used for testing purposes, some have notes in their descriptions.

Benchmarks run under WSL2 with a Ryzen 7 2700 with the command:

```bash
go test -v ./test/benchmark/ -bench=. -benchmem | grep '^Benchmark.*-16'
```

## Results

```
BenchmarkInternals-16                                     319524              3704 ns/op            1418 B/op         17 allocs/op
BenchmarkInternalRequest-16                               658626              1618 ns/op             680 B/op          7 allocs/op
BenchmarkInternalExecute-16                               433489              2422 ns/op             857 B/op         13 allocs/op
BenchmarkInternalExecuteBytes-16                          537806              2005 ns/op            1352 B/op         13 allocs/op
BenchmarkInternalURLWithAuthorizationHash-16             1948528               597.3 ns/op           216 B/op          5 allocs/op
BenchmarkCacheSummonerByPUUIDNoCache-16                   212210              5158 ns/op            1498 B/op         17 allocs/op
BenchmarkCacheBigCacheSummonerByPUUID-16                  343641              3381 ns/op            1008 B/op          7 allocs/op
BenchmarkCacheRedisSummonerByPUUID-16                      22898             55347 ns/op            1212 B/op         14 allocs/op
BenchmarkDataMatchByID-16                                   1940            583458 ns/op           70170 B/op        143 allocs/op
BenchmarkDataMatchTimeline-16                                284           4325004 ns/op         1044960 B/op       1155 allocs/op
BenchmarkDataVALContentAllLocales-16                          22          48749630 ns/op        11581305 B/op     131492 allocs/op
BenchmarkParallelTestRateLimit-16                            100         100045533 ns/op            2812 B/op         31 allocs/op
BenchmarkParallelSummonerByPUUID-16                       338798              3323 ns/op            1496 B/op         17 allocs/op
BenchmarkParallelRedisCachedSummonerByPUUID-16            130627              8568 ns/op            1213 B/op         14 allocs/op
BenchmarkParallelSummonerByAccessToken-16                 318332              3816 ns/op            2144 B/op         25 allocs/op
BenchmarkParallelMatchListByPUUID-16                      277953              4198 ns/op            2768 B/op         32 allocs/op
```

## Integration

> [!NOTE]
> Integration tests are meant to be run manually.

The objective of these tests is to test some methods from different games against the live Riot Games API, making sure the different HTTP methods are working as intended. Ideally, these tests should only contain methods allowed by a development key and should be only a handful of tests to avoid getting rate limited.

Run tests using:

```bash
RIOT_GAMES_API_KEY=RGAPI... go test -v -tags=integration ./test/integration
```

or if using PowerShell:

```powershell
$env:RIOT_GAMES_API_KEY="RGAPI..."; go test -v -tags=integration ./test/integration; Remove-Item Env:RIOT_GAMES_API_KEY
```
