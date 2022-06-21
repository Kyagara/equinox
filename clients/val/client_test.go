package val_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
)

func TestVALClient(t *testing.T) {
	client := val.NewVALClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil VALClient")
}
