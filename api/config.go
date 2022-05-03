package api

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

// An config object for the EquinoxClient.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Log level. Default: api.FatalLevel
	LogLevel LogLevel
	// Timeout for http.Request in seconds, 0 disables it. Default: 10
	Timeout time.Duration
	// Retry request if it returns a 429 status code. Default: true
	Retry bool
}

func (c *EquinoxConfig) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddBool("retry", c.Retry)
	encoder.AddString("timeout", fmt.Sprintf("%ds", c.Timeout))

	return nil
}
