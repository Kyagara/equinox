# Tests

## Benchmark

> Benchmarks are currently not automated, take the results with a grain of salt.

Benchmarks are separated in three files: parallel, data and cache.

- parallel: Variety of benchmarks to test parallelism. Used more as a test than a benchmark.
- cache: Cache benchmarks for both BigCache and Redis.
- data: Benchmarks that use data from the live Riot Games API, located in the `data` folder.

Benchmarks should be using a configuration closest to the one used in production, with cache disabled when possible and log level set to warn.

## Integration

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
