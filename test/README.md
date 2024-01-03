# Tests

## Benchmark

> Benchmarks are currently not automated, take the results with a grain of salt.

Benchmarks are separated in four files: parallel, data, cache and internal.

- parallel: Variety of benchmarks to test parallelism. Used more as a test than a benchmark.
- cache: Cache benchmarks for both BigCache and Redis.
- data: Benchmarks that use data from the live Riot Games API, located in the `data` folder.
- internal: Focused mainly on Request and Execute/ExecuteRaw and internal client functions.

Benchmarks should be using a configuration close to the one used in production. `HTTPClient` timeout is disabled as I believe the context should be used instead (the request is even created with `http.NewRequestWithContext(ctx, ...)`).

Keep in mind that helpers like `httpmock` and `testify`, depending on how they are used, might make some allocations not found in production.

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
