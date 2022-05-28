package api

import (
	"go.uber.org/zap/zapcore"
)

// An config object for the EquinoxClient.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Cluster name, using the nearest cluster to you is recommended.
	Cluster Cluster
	// Log level. Defaults to api.FatalLevel
	LogLevel LogLevel
	// Timeout for HTTP Request in seconds, 0 disables it. Defaults to 10
	Timeout int
	// TTL for the cache in seconds, 0 disables caching. Defaults to 120 seconds
	TTL int
	// Retry request if it returns a 429 status code. Defaults to true
	Retry bool
}

func (c *EquinoxConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddBool("retry", c.Retry)
	encoder.AddInt("timeout", c.Timeout)
	encoder.AddInt("ttl", c.TTL)

	return nil
}
