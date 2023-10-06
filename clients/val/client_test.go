package val_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestVALClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil VALClient")
	client := val.NewVALClient(c)

	require.NotNil(t, client, "expecting non-nil VALClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil VALClient")
	client = val.NewVALClient(c)

	require.Nil(t, client, "expecting nil VALClient")
}
