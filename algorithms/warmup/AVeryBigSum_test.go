package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumAVeryBigSum(t *testing.T) {
	assert.Equal(t, int64(5000000015), sumAVeryBigSum([]int64{
		1000000001,
		1000000002,
		1000000003,
		1000000004,
		1000000005,
	}))
}
