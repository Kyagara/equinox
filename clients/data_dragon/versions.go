package data_dragon

import (
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type VersionEndpoint struct {
	internalClient *internal.InternalClient
}

func (e *VersionEndpoint) Latest() (*string, error) {
	logger := e.internalClient.Logger("Data Dragon", "version", "Latest")
	logger.Debug("Method executed")

	var versions *[]string

	err := e.internalClient.DataDragonGet(VersionsURL, &versions, "version", "Latest")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return &(*versions)[0], nil
}

func (e *VersionEndpoint) List() (*[]string, error) {
	logger := e.internalClient.Logger("Data Dragon", "version", "List")
	logger.Debug("Method executed")

	var versions *[]string

	err := e.internalClient.DataDragonGet(VersionsURL, &versions, "version", "List")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return versions, nil
}
