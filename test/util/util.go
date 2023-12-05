package util

import (
	"net/http"
	"os"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	jsonv2 "github.com/go-json-experiment/json"
	"github.com/rs/zerolog"
)

// Creates an EquinoxConfig for tests.
//
//   - `LogLevel`   : api.DEBUG_LOG_LEVEL
//   - `HTTPClient` : http.Client{}
//   - `Retry`      : 0
//   - `Cache`      : &cache.Cache{TTL: 0}
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:        "RGAPI-TEST",
		LogLevel:   zerolog.DebugLevel,
		HTTPClient: &http.Client{},
		Retry:      0,
		Cache:      &cache.Cache{TTL: 0},
	}
}

// Reads a json file and unmarshalls it into the target.
func ReadFile(filename string, target any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = jsonv2.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
