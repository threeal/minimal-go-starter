package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOdd(t *testing.T) {
	assert.True(t, IsOdd(7))
	assert.False(t, IsOdd(8))
}
