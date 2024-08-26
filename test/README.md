# Tests

## Benchmark

Keep in mind that since requests are mocked using `httpmock`, results (time, bytes, allocs) will be different from production, specially since you can't make 300000 requests in 1 second to the Riot Games API.

Benchmarks clients should be using a configuration close to the one used in production. Some benchmarks are mainly used for testing purposes, check comments.

Benchmarks are separated in four files: parallel, data, cache and internal.

- parallel: Variety of benchmarks to test parallelism. Used more as tests to check for race conditions.
- cache: Cache benchmarks for both BigCache and Redis.
- data: Benchmarks that use data from the live Riot Games API.
- internal: Focused on InternalClient functions.

Command to run benchmarks and generate a markdown table:

```bash
go test ./test/benchmark/ -bench=. -benchmem -v | ./test/benchmark.sh
```

### Results

Using WSL2 on a Ryzen 7 2700.

| Benchmark                          |     ops |     ns/op | bytes/op | allocs/op |
| ---------------------------------- | ------: | --------: | -------: | --------: |
| InternalRequest                    | 1066546 |      1131 |      592 |         4 |
| InternalExecute                    |  526686 |      2151 |      769 |        10 |
| InternalExecuteBytes               |  596833 |      1818 |     1280 |        11 |
| CacheDisabledSummonerByPUUID       |  225356 |      5108 |     1458 |        15 |
| CacheBigCacheSummonerByPUUID       |  338178 |      3459 |     1040 |         7 |
| CacheRedisSummonerByPUUID          |   22681 |     52650 |     1244 |        14 |
| CacheGetKey                        |  790894 |      1846 |     1320 |         5 |
| DataMatchByID                      |    1750 |    685300 |    59216 |       164 |
| DataMatchTimeline                  |     184 |   6472210 |  1627204 |      1680 |
| DataVALContentAllLocales           |      19 |  64913132 | 14872587 |    155696 |
| ParallelTestRateLimit              |     100 | 100036531 |     2936 |        30 |
| ParallelSummonerByPUUID            |  358532 |      3287 |     1456 |        15 |
| ParallelRedisCachedSummonerByPUUID |  125478 |     10090 |     1245 |        14 |
| ParallelSummonerByAccessToken      |  316081 |      3688 |     2096 |        24 |
| ParallelMatchListByPUUID           |  289644 |      3965 |     1904 |        23 |

## Integration

> [!NOTE]
> Integration tests are only meant to be run manually.

The objective of these tests is to run some methods from different games against the live Riot Games API, making sure different HTTP methods are working as intended.

Run tests using:

```bash
RIOT_GAMES_API_KEY=RGAPI... go test -tags=integration ./test/integration -v -failfast
```

or if using PowerShell:

```powershell
$env:RIOT_GAMES_API_KEY="RGAPI..."; go test -tags=integration ./test/integration -v -failfast; Remove-Item Env:RIOT_GAMES_API_KEY
```
