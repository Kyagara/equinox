package util

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/rs/zerolog"
)

// Creates an EquinoxConfig for tests.
//
//   - Key    : "RGAPI-TEST"
//   - Retry  : api.Retry{}
//   - Logger : zerolog.TraceLevel, prettified
func NewTestEquinoxConfig() api.EquinoxConfig {
	return api.EquinoxConfig{
		Key:    "RGAPI-TEST",
		Retry:  api.Retry{},
		Logger: api.Logger{Pretty: true, EnableTimestamp: true, Level: zerolog.TraceLevel, EnableConfigurationLogging: true},
	}
}

// Returns a equinox client, without caching or rate limiting, no Retry and a logger with zerolog.TraceLevel and pretty print.
func NewTestInternalClient(t *testing.T) *internal.Client {
	config := NewTestEquinoxConfig()
	internal, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	return internal
}

// Returns a custom equinox client, without caching or rate limiting.
func NewTestCustomInternalClient(t *testing.T, config api.EquinoxConfig) *internal.Client {
	internal, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	return internal
}

// Returns a equinox client close to the default one but without caching or rate limiting.
func NewBenchmarkEquinoxClient(b *testing.B) *equinox.Equinox {
	b.Helper()
	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := equinox.NewCustomClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}
	return client
}

func NewTestLogger() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.TraceLevel)
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
