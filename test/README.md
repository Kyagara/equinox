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
go test ./test/benchmark/ -bench=. -benchmem -v | ./test/benchmark.sh
```

### Results

Using WSL2 on a Ryzen 7 2700.

| Benchmark                          |     ops |     ns/op | bytes/op | allocs/op |
| ---------------------------------- | ------: | --------: | -------: | --------: |
| InternalRequest                    | 1086412 |      1101 |      560 |         4 |
| InternalExecute                    |  525964 |      2305 |      809 |        12 |
| InternalExecuteBytes               |  583923 |      1944 |     1320 |        13 |
| CacheDisabledSummonerByPUUID       |  214899 |      5494 |     1498 |        17 |
| CacheBigCacheSummonerByPUUID       |  336384 |      3587 |     1008 |         7 |
| CacheRedisSummonerByPUUID          |   21764 |     56377 |     1212 |        14 |
| CacheGetKey                        |  803038 |      1848 |     1320 |         5 |
| DataMatchByID                      |    1792 |    679945 |    54478 |       166 |
| DataMatchTimeline                  |     187 |   6577700 |  1624708 |      1681 |
| DataVALContentAllLocales           |      16 |  64658314 | 14866372 |    155492 |
| ParallelTestRateLimit              |     100 | 100028519 |     3001 |        31 |
| ParallelSummonerByPUUID            |  351018 |      3563 |     1496 |        17 |
| ParallelRedisCachedSummonerByPUUID |  128907 |      9666 |     1213 |        14 |
| ParallelSummonerByAccessToken      |  300252 |      4185 |     2113 |        26 |
| ParallelMatchListByPUUID           |  273434 |      4643 |     2062 |        26 |

## Integration

> [!NOTE]
> Integration tests are only meant to be run manually.

The objective of these tests is to test some methods from different games against the live Riot Games API, making sure the different HTTP methods are working as intended.

Run tests using:

```bash
RIOT_GAMES_API_KEY=RGAPI... go test -tags=integration ./test/integration -v -failfast
```

or if using PowerShell:

```powershell
$env:RIOT_GAMES_API_KEY="RGAPI..."; go test -tags=integration ./test/integration -v -failfast; Remove-Item Env:RIOT_GAMES_API_KEY
```
