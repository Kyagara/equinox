package data_dragon_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
)

func TestDataDragonClient(t *testing.T) {
	client := data_dragon.NewDataDragonClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil DataDragonClient")
}
