package bitvector_test

import (
	"github.com/teivah/bitvector"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Len64(t *testing.T) {
	var bv bitvector.Len64

	// Set
	assert.Equal(t, bitvector.Len64(0), bv)
	bv = bv.Set(0, true)
	assert.Equal(t, bitvector.Len64(1), bv)
	bv = bv.Set(63, true)
	assert.Equal(t, bitvector.Len64(9223372036854775809), bv)
	bv = bv.Set(64, true)
	assert.Equal(t, bitvector.Len64(9223372036854775809), bv)
	bv = bv.Set(62, true)
	assert.Equal(t, bitvector.Len64(13835058055282163713), bv)
	bv = bv.Set(62, false)
	assert.Equal(t, bitvector.Len64(9223372036854775809), bv)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(63))
	assert.False(t, bv.Get(62))

	// Toggle
	bv = 0
	bv = bv.Toggle(63)
	assert.Equal(t, bitvector.Len64(9223372036854775808), bv)

	// Clear
	bv = 0
	var i uint8
	for ; i < 64; i++ {
		bv = bv.Set(i, true)
	}
	bv = bv.Clear(6, 2)
	assert.Equal(t, bitvector.Len64(math.MaxUint64), bv)
	bv = bv.Clear(2, 6)
	assert.Equal(t, bitvector.Len64(18446744073709551555), bv)
	bv = bv.Clear(60, 62)
	bv = bv.Clear(63, 64)
	bv = bv.Clear(60, 64)
	bv = bv.Clear(20, 34)

	// String
	assert.Equal(t, "0000111111111111111111111111110000000000000011111111111111000011", bv.String())
}
