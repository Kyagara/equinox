package api

import (
	"time"

	"github.com/rs/zerolog"
)

// Configuration for the equinox client.
type EquinoxConfig struct {
	// Riot API Key.
	Key string
	// Log Level, Pretty print, Timestamp logging, etc.
	Logger Logger
	// Maximum number of retries, Jitter.
	Retry Retry
}

// Retry configuration object.
//
// Retries have a exponential backoff mechanism.
type Retry struct {
	// Maximum number of retries, 0 disables retries.
	MaxRetries int
	// Jitter, in milliseconds, added to the retry interval.
	Jitter time.Duration
}

// Logger configuration object.
type Logger struct {
	TimeFieldFormat string
	Level           zerolog.Level
	// Enables prettified logging.
	Pretty bool
	// Prints the timestamp.
	EnableTimestamp bool
}
