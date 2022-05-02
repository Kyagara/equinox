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
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		// (Legal values: scheduled, in_progress, complete)
		MaintenanceStatus string `json:"maintenance_status"`
	} `json:"maintenances"`
	Incidents []struct {
		ArchiveAt time.Time `json:"archive_at"`
		Titles    []struct {
			Content string `json:"content"`
			Locale  string `json:"locale"`
		} `json:"titles"`
		UpdatedAt time.Time `json:"updated_at"`
		// (Legal values: info, warning, critical)
		IncidentSeverity string `json:"incident_severity"`
		// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
		Platforms []string `json:"platforms"`
		Updates   []struct {
			UpdatedAt    time.Time `json:"updated_at"`
			Translations []struct {
				Content string `json:"content"`
				Locale  string `json:"locale"`
			} `json:"translations"`
			Author    string    `json:"author"`
			Publish   bool      `json:"publish"`
			CreatedAt time.Time `json:"created_at"`
			ID        int       `json:"id"`
			// (Legal values: riotclient, riotstatus, game)
			PublishLocations []string `json:"publish_locations"`
		} `json:"updates"`
		CreatedAt         time.Time `json:"created_at"`
		ID                int       `json:"id"`
		MaintenanceStatus string    `json:"maintenance_status"`
	} `json:"incidents"`
}

// Get League of Legends status for the given platform.
func (c *StatusEndpoint) PlatformStatus(region api.LOLRegion) (*PlatformDataDTO, error) {
	res := PlatformDataDTO{}

	err := c.internalClient.Do(http.MethodGet, region, StatusURL, nil, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
