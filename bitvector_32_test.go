package bitvector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Len32(t *testing.T) {
	var bv Len32

	// Set
	assert.Equal(t, Len32(0), bv)
	bv = bv.Set(0, true)
	assert.Equal(t, Len32(1), bv)
	bv = bv.Set(31, true)
	assert.Equal(t, Len32(2147483649), bv)
	bv = bv.Set(32, true)
	assert.Equal(t, Len32(2147483649), bv)
	bv = bv.Set(30, true)
	assert.Equal(t, Len32(3221225473), bv)
	bv = bv.Set(30, false)
	assert.Equal(t, Len32(2147483649), bv)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(31))
	assert.False(t, bv.Get(30))

	// Toggle
	bv = 0
	bv = bv.Toggle(15)
	assert.Equal(t, Len32(32768), bv)

	// Clear
	bv = 0
	var i uint8
	for ; i < 32; i++ {
		bv = bv.Set(i, true)
	}
	bv = bv.Clear(6, 2)
	assert.Equal(t, Len32(math.MaxUint32), bv)
	bv = bv.Clear(2, 6)
	assert.Equal(t, Len32(4294967235), bv)

	// String
	assert.Equal(t, "11111111111111111111111111000011", bv.String())
}
