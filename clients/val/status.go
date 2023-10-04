package val

import (
	"fmt"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

// Get VALORANT status for the given platform.
func (e *StatusEndpoint) PlatformStatus(shard Shard) (*api.PlatformDataDTO, error) {
	logger := e.internalClient.Logger("VAL", "val-status-v1", "PlatformStatus")
	logger.Debug("Method executed")

	if shard == ESPORTS {
		return nil, fmt.Errorf("the region ESPORTS is not available for this method")
	}

	var status *api.PlatformDataDTO

	err := e.internalClient.Get(shard, StatusURL, &status, "val-status-v1", "PlatformStatus", "")
	if err != nil {
		logger.Error("Method failed", zap.Error(err))
		return nil, err
	}

	return status, nil
}
