package internal

import (
	"os"
	"strings"
	"sync"

	"github.com/Kyagara/equinox/api"
	"github.com/rs/zerolog"
)

type Loggers struct {
	main    zerolog.Logger
	methods map[string]zerolog.Logger
	mutex   sync.Mutex
}

func NewLogger(config *api.EquinoxConfig) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if config == nil {
		return zerolog.Nop()
	}
	switch config.LogLevel {
	case zerolog.Disabled:
		return zerolog.Nop()
	case zerolog.DebugLevel:
		return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Object("equinox", config).Timestamp().Logger()
	default:
		return zerolog.New(os.Stderr).Level(config.LogLevel).With().Object("equinox", config).Timestamp().Logger()
	}
}

// Used to access the internal logger, this is used to log events from other clients (RiotClient, LOLClient...).
func (c *InternalClient) Logger(id string) zerolog.Logger {
	c.loggers.mutex.Lock()
	defer c.loggers.mutex.Unlock()
	if logger, ok := c.loggers.methods[id]; ok {
		return logger
	}
	names := strings.Split(id, "_")
	logger := c.loggers.main.With().Str("client", names[0]).Str("endpoint", names[1]).Str("method", names[2]).Logger()
	c.loggers.methods[id] = logger
	return logger
}
