// This package only contains integration tests and are meant to be run manually.
//
// The objective of these tests is to test some methods from different games against the live Riot Games API, making sure the different HTTP methods are working as intended. Ideally, these tests should only contain methods allowed by a development key and should be only a handful of tests to avoid getting rate limited.
//
// After setting up the environment variable "RIOT_GAMES_API_KEY" with your API key, you can run tests using:
//
//	go test -v -tags=integration ./test/integration
package integration
