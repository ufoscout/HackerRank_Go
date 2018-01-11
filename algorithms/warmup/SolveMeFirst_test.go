package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMeFirst(t *testing.T) {

	assert.Equal(t, uint32(0), solveMeFirst(0, 0))
	assert.Equal(t, uint32(10), solveMeFirst(10, 0))
	assert.Equal(t, uint32(10), solveMeFirst(0, 10))
	assert.Equal(t, uint32(3), solveMeFirst(1, 2))

}
