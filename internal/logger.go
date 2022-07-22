package internal

import (
	"log"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Creates a new zap.Logger from the configuration parameters provided.
func NewLogger(config *api.EquinoxConfig) *zap.SugaredLogger {
	zapConfig := zap.NewProductionConfig()

	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))

	logger, err := zapConfig.Build(zap.Fields(zap.Object("equinox", config)))

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	return logger.Sugar()
}

// Used to access the logger from the InternalClient, this is used to log events from other clients (RiotClient, LOLClient...)
func (c *InternalClient) Logger(client string, endpoint string, method string) *zap.SugaredLogger {
	logger := c.logger.With("client", client, "endpoint", endpoint, "method", method)

	return logger
}
