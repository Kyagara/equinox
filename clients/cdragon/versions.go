package cdragon

import (
	"github.com/Kyagara/equinox/internal"
)

type VersionEndpoint struct {
	internalClient *internal.InternalClient
}

func (e *VersionEndpoint) Latest() (string, error) {
	version, err := e.internalClient.GetDDragonLOLVersions("CDragon_Version_Latest")
	if err != nil {
		return "", err
	}
	return version[0], nil
}

func (e *VersionEndpoint) List() ([]string, error) {
	return e.internalClient.GetDDragonLOLVersions("CDragon_Version_List")
}
