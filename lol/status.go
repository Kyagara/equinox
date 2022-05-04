package lol

import (
	"net/http"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

// Get League of Legends status for the given platform.
func (s *StatusEndpoint) PlatformStatus(region Region) (*api.PlatformDataDTO, error) {
	logger := s.internalClient.Logger("lol").With("endpoint", "status", "method", "PlatformStatus")

	var status *api.PlatformDataDTO

	err := s.internalClient.Do(http.MethodGet, region, StatusURL, nil, &status, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return status, nil
}
