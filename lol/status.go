package lol

import (
	"net/http"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
)

type StatusEndpoint struct {
	internalClient *internal.InternalClient
}

type PlatformDataDTO struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Locales      []string `json:"locales"`
	Maintenances []struct {
		ArchiveAt time.Time `json:"archive_at"`
		Titles    []struct {
			Content string `json:"content"`
			Locale  string `json:"locale"`
		} `json:"titles"`
		UpdatedAt        time.Time `json:"updated_at"`
		IncidentSeverity string    `json:"incident_severity"`
		Platforms        []string  `json:"platforms"`
		Updates          []struct {
			UpdatedAt    string `json:"updated_at"`
			Translations []struct {
				Content string `json:"content"`
				Locale  string `json:"locale"`
			} `json:"translations"`
			Author           string    `json:"author"`
			Publish          bool      `json:"publish"`
			CreatedAt        time.Time `json:"created_at"`
			ID               int       `json:"id"`
			PublishLocations []string  `json:"publish_locations"`
		} `json:"updates"`
		CreatedAt         time.Time `json:"created_at"`
		ID                int       `json:"id"`
		MaintenanceStatus string    `json:"maintenance_status"`
	} `json:"maintenances"`
	Incidents []struct {
		ArchiveAt time.Time `json:"archive_at"`
		Titles    []struct {
			Content string `json:"content"`
			Locale  string `json:"locale"`
		} `json:"titles"`
		UpdatedAt        time.Time `json:"updated_at"`
		IncidentSeverity string    `json:"incident_severity"`
		Platforms        []string  `json:"platforms"`
		Updates          []struct {
			UpdatedAt    time.Time `json:"updated_at"`
			Translations []struct {
				Content string `json:"content"`
				Locale  string `json:"locale"`
			} `json:"translations"`
			Author           string    `json:"author"`
			Publish          bool      `json:"publish"`
			CreatedAt        time.Time `json:"created_at"`
			ID               int       `json:"id"`
			PublishLocations []string  `json:"publish_locations"`
		} `json:"updates"`
		CreatedAt         time.Time `json:"created_at"`
		ID                int       `json:"id"`
		MaintenanceStatus string    `json:"maintenance_status"`
	} `json:"incidents"`
}

// Get Status
func (c *StatusEndpoint) GetStatus(region api.Region) (*PlatformDataDTO, error) {
	res := PlatformDataDTO{}

	err := c.internalClient.Do(http.MethodGet, region, StatusEndpointURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
