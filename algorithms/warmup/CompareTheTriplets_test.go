package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareValues(t *testing.T) {
	result := calculateScore([]int{0,0,0}, []int{0,0,0});
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 0, result[1])

	result = calculateScore([]int{1,2,3}, []int{1,2,3});
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 0, result[1])

	result = calculateScore([]int{1,2,3}, []int{0,0,0});
	assert.Equal(t, 3, result[0])
	assert.Equal(t, 0, result[1])

	result = calculateScore([]int{0,0,0}, []int{1,2,3});
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 3, result[1])

	result = calculateScore([]int{0,4,0}, []int{1,2,3});
	assert.Equal(t, 1, result[0])
	assert.Equal(t, 2, result[1])
}

