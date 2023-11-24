package cdragon

import (
	"context"

	"github.com/Kyagara/equinox/internal"
)

type VersionEndpoint struct {
	internalClient *internal.InternalClient
}

func (e *VersionEndpoint) Latest(ctx context.Context) (string, error) {
	version, err := e.internalClient.GetDDragonLOLVersions(ctx, "CDragon_Version_Latest")
	if err != nil {
		return "", err
	}
	return version[0], nil
}

func (e *VersionEndpoint) List(ctx context.Context) ([]string, error) {
	return e.internalClient.GetDDragonLOLVersions(ctx, "CDragon_Version_List")
}
