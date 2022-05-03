package lol

import (
	"net/http"
	"time"

	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

type PlatformDataDTO struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Locales      []string    `json:"locales"`
	Maintenances []StatusDTO `json:"maintenances"`
	Incidents    []StatusDTO `json:"incidents"`
}

type ContentDTO struct {
	Content string `json:"content"`
	Locale  string `json:"locale"`
}

type StatusDTO struct {
	ArchiveAt time.Time    `json:"archive_at"`
	Titles    []ContentDTO `json:"titles"`
	UpdatedAt time.Time    `json:"updated_at"`
	// (Legal values: info, warning, critical)
	IncidentSeverity string `json:"incident_severity"`
	// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms         []string    `json:"platforms"`
	Updates           []UpdateDTO `json:"updates"`
	CreatedAt         time.Time   `json:"created_at"`
	ID                int         `json:"id"`
	MaintenanceStatus string      `json:"maintenance_status"`
}

type UpdateDTO struct {
	UpdatedAt    time.Time    `json:"updated_at"`
	Translations []ContentDTO `json:"translations"`
	Author       string       `json:"author"`
	Publish      bool         `json:"publish"`
	CreatedAt    time.Time    `json:"created_at"`
	ID           int          `json:"id"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string `json:"publish_locations"`
}

// Get League of Legends status for the given platform.
func (s *StatusEndpoint) PlatformStatus(region Region) (*PlatformDataDTO, error) {
	logger := s.internalClient.Logger().With("endpoint", "status", "method", "PlatformStatus")

	var status *PlatformDataDTO

	err := s.internalClient.Do(http.MethodGet, region, StatusURL, nil, &status, "")

	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	return status, nil
}
