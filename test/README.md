# Tests

## Benchmark

Keep in mind that since requests are mocked using `httpmock`, results (time, bytes, allocs) will be different from production, specially since you can't make 300000 requests in 1 second to the Riot Games API.

Benchmarks clients should be using a configuration close to the one used in production, some are mainly used for testing purposes, check comments.

Benchmarks are separated in four files: parallel, data, cache and internal.

- parallel: Variety of benchmarks to test parallelism. Used more as tests to check for race conditions.
- cache: Cache benchmarks for both BigCache and Redis.
- data: Benchmarks that use data from the live Riot Games API.
- internal: Focused on InternalClient functions.

Command to run benchmarks and generate a markdown table:

```bash
go test -v ./test/benchmark/ -bench=. -benchmem | ./benchmark.sh
```

### Results

Using WSL2 on a Ryzen 7 2700.

| Benchmark                          |     ops |     ns/op | bytes/op | allocs/op |
| ---------------------------------- | ------: | --------: | -------: | --------: |
| Internals                          |  314580 |      3791 |     1418 |        17 |
| InternalRequest                    | 1086412 |      1101 |      560 |         4 |
| InternalExecute                    |  462622 |      2490 |      857 |        13 |
| InternalExecuteBytes               |  547890 |      2070 |     1352 |        13 |
| CacheDisabledSummonerByPUUID       |  214899 |      5494 |     1498 |        17 |
| CacheBigCacheSummonerByPUUID       |  336384 |      3587 |     1008 |         7 |
| CacheRedisSummonerByPUUID          |   21764 |     56377 |     1212 |        14 |
| CacheGetKey                        |  803038 |      1848 |     1320 |         5 |
| DataMatchByID                      |    1744 |    658620 |    70331 |       166 |
| DataMatchTimeline                  |     187 |   6577700 |  1624708 |      1681 |
| DataVALContentAllLocales           |      16 |  64658314 | 14866372 |    155492 |
| ParallelTestRateLimit              |     100 | 100028519 |     3001 |        31 |
| ParallelSummonerByPUUID            |  351018 |      3563 |     1496 |        17 |
| ParallelRedisCachedSummonerByPUUID |  128907 |      9666 |     1213 |        14 |
| ParallelSummonerByAccessToken      |  308845 |      4221 |     2152 |        26 |
| ParallelMatchListByPUUID           |  266372 |      4869 |     2768 |        32 |

## Integration

> [!NOTE]
> Integration tests are only meant to be run manually.

The objective of these tests is to test some methods from different games against the live Riot Games API, making sure the different HTTP methods are working as intended.

Run tests using:

```bash
RIOT_GAMES_API_KEY=RGAPI... go test -v -tags=integration ./test/integration
```

or if using PowerShell:

```powershell
$env:RIOT_GAMES_API_KEY="RGAPI..."; go test -v -tags=integration ./test/integration; Remove-Item Env:RIOT_GAMES_API_KEY
```
