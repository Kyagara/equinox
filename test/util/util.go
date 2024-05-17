package util

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/cache"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/redis/go-redis/v9"
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

// Returns a equinox client close to the default one but without caching and rate limiting.
func NewBenchmarkEquinoxClient(b *testing.B) *equinox.Equinox {
	b.Helper()
	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := equinox.NewCustomClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}
	return client
}

// Returns a equinox client close to the default one but with RedisCache, no rate limiting.
func NewBenchmarkRedisCacheEquinoxClient(b *testing.B) *equinox.Equinox {
	b.Helper()
	redisConfig := &redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}

	ctx := context.Background()
	cache, err := cache.NewRedis(ctx, redisConfig, 4*time.Minute)
	if err != nil {
		b.Fatal(err)
	}

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := equinox.NewCustomClient(config, nil, cache, nil)
	if err != nil {
		b.Fatal(err)
	}
	return client
}

func NewTestLogger() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.TraceLevel)
}
