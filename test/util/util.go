package util

import (
	"net/http"
	"os"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// Creates an EquinoxConfig for tests.
//
//   - `Logger`     : zerolog.TraceLevel
//   - `HTTPClient` : http.Client{}
//   - `Retry`      : 0
//   - `Cache`      : &cache.Cache{TTL: 0}
//   - `RateLimit`  : Disabled
func NewTestEquinoxConfig() api.EquinoxConfig {
	return api.EquinoxConfig{
		Key:        "RGAPI-TEST",
		HTTPClient: &http.Client{},
		Retry:      api.Retry{},
		Cache:      &cache.Cache{},
		RateLimit:  &ratelimit.RateLimit{},
		Logger: api.Logger{
			Level:               zerolog.TraceLevel,
			EnableTimestamp:     true,
			Pretty:              true,
			EnableConfigLogging: true,
		},
	}
}

// Reads a json file and returns its content as []byte.
func ReadFile(b *testing.B, filename string) []byte {
	b.Helper()
	data, err := os.ReadFile(filename)
	require.NoError(b, err)
	return data
}
