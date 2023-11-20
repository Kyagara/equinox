package ddragon_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestDDragonClient(t *testing.T) {
	c := &internal.InternalClient{}
	client := ddragon.NewDDragonClient(c)
	require.NotEmpty(t, client, "expecting non-nil DDragonClient")
}
