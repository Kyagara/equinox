package internal

import (
	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Creates a new zap.Logger from the configuration parameters provided.
func NewLogger(config *api.EquinoxConfig) (*zap.Logger, error) {
	if config.LogLevel == api.NopLevel {
		return zap.NewNop(), nil
	}

	zapConfig := zap.NewProductionConfig()

	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))

	return zapConfig.Build(zap.Fields(zap.Object("equinox", config)))
}

// Used to access the internal logger, this is used to log events from other clients (RiotClient, LOLClient...).
func (c *InternalClient) Logger(client string, endpoint string, method string) *zap.Logger {
	return c.logger.With(zap.String("client", client), zap.String("endpoint", endpoint), zap.String("method", method))
}
