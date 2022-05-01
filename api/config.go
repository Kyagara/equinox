package api

import (
	"time"
)

// An config object for the EquinoxClient.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Debug mode. Default: false
	Debug bool
	// Timeout for http.Request in seconds, 0 disables it. Default: 10
	Timeout time.Duration
	// Retry request if it returns a 429 status code. Default: true
	Retry bool
}
