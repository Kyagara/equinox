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

Command to run benchmarks and generate a markdown table:

```bash
go test -v ./test/benchmark/ -bench=. -benchmem | ./benchmark.sh
```

### Results

Using WSL2 on a Ryzen 7 2700.

| Benchmark                          |     ops |     ns/op | bytes/op | allocs/op |
| ---------------------------------- | ------: | --------: | -------: | --------: |
| Internals                          |  325980 |      3739 |     1418 |        17 |
| InternalRequest                    | 1000000 |      1105 |      560 |         4 |
| InternalExecute                    |  451694 |      2495 |      857 |        13 |
| InternalExecuteBytes               |  594973 |      2035 |     1352 |        13 |
| InternalGetCacheKey                |  758509 |      1585 |     1321 |         5 |
| CacheDisabledSummonerByPUUID       |  217107 |      5299 |     1499 |        17 |
| CacheBigCacheSummonerByPUUID       |  339098 |      3470 |     1008 |         7 |
| CacheRedisSummonerByPUUID          |   22164 |     56318 |     1212 |        14 |
| DataMatchByID                      |    1869 |    684389 |    70330 |       166 |
| DataMatchTimeline                  |     180 |   6420727 |  1624715 |      1681 |
| DataVALContentAllLocales           |      18 |  62570122 | 14865476 |    155491 |
| ParallelTestRateLimit              |     100 | 100028115 |     2813 |        31 |
| ParallelSummonerByPUUID            |  354766 |      3285 |     1496 |        17 |
| ParallelRedisCachedSummonerByPUUID |  131782 |      9168 |     1213 |        14 |
| ParallelSummonerByAccessToken      |  272334 |      4571 |     2178 |        26 |

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
