package bitvector_test

import (
	"github.com/teivah/bitvector"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Len16(t *testing.T) {
	var bv bitvector.Len16

	// Set
	assert.Equal(t, bitvector.Len16(0), bv)
	bv = bv.Set(0, true)
	assert.Equal(t, bitvector.Len16(1), bv)
	bv = bv.Set(15, true)
	assert.Equal(t, bitvector.Len16(32769), bv)
	bv = bv.Set(16, true)
	assert.Equal(t, bitvector.Len16(32769), bv)
	bv = bv.Set(14, true)
	assert.Equal(t, bitvector.Len16(49153), bv)
	bv = bv.Set(14, false)
	assert.Equal(t, bitvector.Len16(32769), bv)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(15))
	assert.False(t, bv.Get(14))

	// Toggle
	bv = 0
	bv = bv.Toggle(15)
	assert.Equal(t, bitvector.Len16(32768), bv)

	// Clear
	bv = 0
	var i uint8
	for ; i < 16; i++ {
		bv = bv.Set(i, true)
	}
	bv = bv.Clear(6, 2)
	assert.Equal(t, bitvector.Len16(math.MaxUint16), bv)
	bv = bv.Clear(2, 6)
	assert.Equal(t, bitvector.Len16(65475), bv)

	// String
	assert.Equal(t, "1111111111000011", bv.String())
}
