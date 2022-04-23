package equinox

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreeChampionsRotation(t *testing.T) {
	c := NewClient(os.Getenv("RIOT-API-KEY"))

	res, err := c.LOL.Champion.FreeRotation(BR1)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil result")
}
