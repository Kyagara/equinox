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

func NewLogger(config api.EquinoxConfig) zerolog.Logger {
	if config == (api.EquinoxConfig{}) || config.Logger.Level == zerolog.Disabled {
		return zerolog.Nop()
	}

	zerolog.TimeFieldFormat = config.Logger.TimeFieldFormat

	var logger zerolog.Logger

	if config.Logger.Pretty {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(config.Logger.Level)
	} else {
		logger = zerolog.New(os.Stderr).Level(config.Logger.Level)
	}

	if config.Logger.EnableTimestamp {
		logger = logger.With().Timestamp().Logger()
	}

	if config.Logger.EnableConfigLogging {
		logger = logger.With().Object("equinox", config).Logger()
	}

	return logger
}

// Used to access the internal logger, this is used to log events from other clients.
func (c *Client) Logger(id string) zerolog.Logger {
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
