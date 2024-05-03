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
BenchmarkInternals-16                                     325980              3739 ns/op            1418 B/op         17 allocs/op
BenchmarkInternalRequest-16                               682623              1630 ns/op             680 B/op          7 allocs/op
BenchmarkInternalExecute-16                               451694              2495 ns/op             857 B/op         13 allocs/op
BenchmarkInternalExecuteBytes-16                          546505              2016 ns/op            1352 B/op         13 allocs/op
BenchmarkInternalURLWithAuthorizationHash-16             1977759               602.6 ns/op           216 B/op          5 allocs/op
BenchmarkCacheSummonerByPUUIDNoCache-16                   217107              5299 ns/op            1499 B/op         17 allocs/op
BenchmarkCacheBigCacheSummonerByPUUID-16                  339098              3470 ns/op            1008 B/op          7 allocs/op
BenchmarkCacheRedisSummonerByPUUID-16                      22164             56318 ns/op            1212 B/op         14 allocs/op
BenchmarkDataMatchByID-16                                   1869            684389 ns/op           70330 B/op        166 allocs/op
BenchmarkDataMatchTimeline-16                                180           6420727 ns/op         1624715 B/op       1681 allocs/op
BenchmarkDataVALContentAllLocales-16                          18          62570122 ns/op        14865476 B/op     155491 allocs/op
BenchmarkParallelTestRateLimit-16                            100         100028115 ns/op            2813 B/op         31 allocs/op
BenchmarkParallelSummonerByPUUID-16                       354766              3285 ns/op            1496 B/op         17 allocs/op
BenchmarkParallelRedisCachedSummonerByPUUID-16            131782              9168 ns/op            1213 B/op         14 allocs/op
BenchmarkParallelSummonerByAccessToken-16                 280219              4000 ns/op            2144 B/op         25 allocs/op
BenchmarkParallelMatchListByPUUID-16                      274206              4173 ns/op            2768 B/op         32 allocs/op
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
