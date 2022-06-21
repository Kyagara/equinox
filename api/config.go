package api

import (
	"go.uber.org/zap/zapcore"
)

// A config object for the EquinoxClient.
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
	// TTL for the cache in seconds, 0 disables caching.
	TTL int
	// Enable or disable rate limiting.
	RateLimit bool
}

func (c *EquinoxConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddBool("retry-enabled", c.Retry)
	encoder.AddInt("client-timeout", c.Timeout)
	encoder.AddInt("cache-ttl", c.TTL)
	encoder.AddBool("rate-limit", c.RateLimit)

	return nil
}
