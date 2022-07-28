package api

import (
	"github.com/Kyagara/equinox/cache"
	"go.uber.org/zap/zapcore"
)

// Configuration for the Equinox client.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Cluster name, using the nearest cluster to you is recommended.
	Cluster Cluster
	// Log level.
	LogLevel LogLevel
	// Timeout for the http.Client in seconds, 0 disables the timeout.
	Timeout int
	// Enable or disable retrying a request if it returns a 429 status code.
	Retry bool
	// The cache used to store all GET requests done by the client.
	Cache *cache.Cache
	// Enable or disable rate limiting.
	RateLimit bool
}

func (c *EquinoxConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddBool("retry-enabled", c.Retry)
	encoder.AddInt("client-timeout", c.Timeout)
	encoder.AddBool("rate-limit", c.RateLimit)
	encoder.AddDuration("cache-ttl", c.Cache.TTL)

	return nil
}
