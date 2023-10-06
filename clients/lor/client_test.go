package lor_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestLORClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil LORClient")
	client := lor.NewLORClient(c)

	require.NotNil(t, client, "expecting non-nil LORClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil LORClient")
	client = lor.NewLORClient(c)

	require.Nil(t, client, "expecting nil LORClient")
}
