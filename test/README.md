# Equinox tests

## Integration

> Integration tests are meant to be run manually.

The objective of these tests is to test some methods from different games against the live Riot Games API, making sure the different http methods are working as intended.

Run tests using:

```bash
RIOT_GAMES_API_KEY=RGAPI... go test -v -tags=integration ./test/integration
```

or if using PowerShell:

```powershell
$env:RIOT_GAMES_API_KEY="RGAPI..."; go test -v -tags=integration ./test/integration; Remove-Item Env:RIOT_GAMES_API_KEY
```
