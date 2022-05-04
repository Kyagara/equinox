package lor_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lor"
	"github.com/stretchr/testify/assert"
)

func TestLORClient(t *testing.T) {
	client := lor.NewLORClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil LORClient")
}
