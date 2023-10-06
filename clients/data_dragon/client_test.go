package data_dragon_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestDataDragonClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil DataDragonClient")
	client := data_dragon.NewDataDragonClient(c)

	require.NotNil(t, client, "expecting non-nil DataDragonClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil DataDragonClient")
	client = data_dragon.NewDataDragonClient(c)

	require.NotNil(t, client, "expecting non-nil DataDragonClient")
}
