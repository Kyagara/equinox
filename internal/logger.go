package internal

import (
	"log"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config *api.EquinoxConfig) *zap.SugaredLogger {
	zapConfig := zap.NewProductionConfig()

	equinoxOptions := &api.EquinoxConfig{
		Retry:   config.Retry,
		Timeout: config.Timeout,
	}

	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel))

	logger, err := zapConfig.Build(zap.Fields(zap.Object("equinox", equinoxOptions)))

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	return logger.Sugar()
}

func (c *InternalClient) Logger() *zap.SugaredLogger {
	return c.logger
}
