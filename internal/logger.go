package internal

import (
	"log"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Creates a new zap.Logger from the configuration parameters provided.
func NewLogger(config *api.EquinoxConfig) *zap.Logger {
	zapConfig := zap.NewProductionConfig()

	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))

	logger, err := zapConfig.Build(zap.Fields(zap.Object("equinox", config)))

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	return logger
}

// Used to access the internal logger, this is used to log events from other clients (RiotClient, LOLClient...)
func (c *InternalClient) Logger(client string, endpoint string, method string) *zap.Logger {
	return c.logger.With(zap.String("client", client), zap.String("endpoint", endpoint), zap.String("method", method))
}
