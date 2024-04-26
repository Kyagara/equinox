package util

import (
	"net/http"
	"os"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/rs/zerolog"
)

// Creates an EquinoxConfig for tests.
//
//   - Logger     : zerolog.TraceLevel prettified
//   - HTTPClient : http.Client{}
//   - Retry      : api.Retry{}
//   - Cache      : &cache.Cache{}
//   - RateLimit  : &ratelimit.RateLimit{}
func NewTestEquinoxConfig() api.EquinoxConfig {
	return api.EquinoxConfig{
		Key:        "RGAPI-TEST",
		HTTPClient: &http.Client{},
		Retry:      api.Retry{},
		Cache:      &cache.Cache{},
		RateLimit:  &ratelimit.RateLimit{},
		Logger:     TestLogger(),
	}
}

func TestLogger() api.Logger {
	return api.Logger{
		Level:               zerolog.TraceLevel,
		Pretty:              true,
		EnableConfigLogging: true,
		EnableTimestamp:     true,
	}
}

// Reads a json file and returns its content as []byte.
func ReadFile(b *testing.B, filename string) []byte {
	b.Helper()
	data, err := os.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	return data
}
