package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a := add(1, 2)
	assert.Equal(t, 3, a)
}
