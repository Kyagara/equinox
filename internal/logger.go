package internal

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Loggers struct {
	main    *zap.Logger
	methods map[string]*zap.Logger
	mutex   sync.RWMutex
}

// Creates a new zap.Logger from the configuration parameters provided.
func NewLogger(config *api.EquinoxConfig) (*zap.Logger, error) {
	if config == nil {
		return nil, fmt.Errorf("error initializing logger, equinox configuration not provided")
	}
	var zapConfig zap.Config
	switch config.LogLevel {
	case api.NOP_LOG_LEVEL:
		return zap.NewNop(), nil
	case api.DEBUG_LOG_LEVEL:
		zapConfig = zap.NewDevelopmentConfig()
	default:
		zapConfig = zap.NewProductionConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))
	}
	return zapConfig.Build(zap.Fields(zap.Object("equinox", config)))
}

// Used to access the internal logger, this is used to log events from other clients (RiotClient, LOLClient...).
func (c *InternalClient) Logger(id string) *zap.Logger {
	if logger, ok := c.loggers.methods[id]; ok {
		return logger
	}
	names := strings.Split(id, "_")
	logger := c.loggers.main.With(zap.String("client", names[0]), zap.String("endpoint", names[1]), zap.String("method", names[2]))
	c.loggers.mutex.Lock()
	c.loggers.methods[id] = logger
	c.loggers.mutex.Unlock()
	return logger
}
