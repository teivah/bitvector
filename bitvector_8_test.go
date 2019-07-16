package bitvector_test

import (
	"github.com/teivah/bitvector"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Len8(t *testing.T) {
	var bv bitvector.Len8

	// Set
	assert.Equal(t, bitvector.Len8(0), bv)
	bv = bv.Set(0, true)
	assert.Equal(t, bitvector.Len8(1), bv)
	bv = bv.Set(7, true)
	assert.Equal(t, bitvector.Len8(129), bv)
	bv = bv.Set(8, true)
	assert.Equal(t, bitvector.Len8(129), bv)
	bv = bv.Set(6, true)
	assert.Equal(t, bitvector.Len8(193), bv)
	bv = bv.Set(6, false)
	assert.Equal(t, bitvector.Len8(129), bv)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(7))
	assert.False(t, bv.Get(6))

	// Toggle
	bv = bv.Toggle(7)
	assert.Equal(t, bitvector.Len8(1), bv)

	// Clear
	bv = 0
	var i uint8
	for ; i < 8; i++ {
		bv = bv.Set(i, true)
	}
	bv = bv.Clear(6, 2)
	assert.Equal(t, bitvector.Len8(math.MaxUint8), bv)
	bv = bv.Clear(2, 6)
	assert.Equal(t, bitvector.Len8(195), bv)

	// String
	assert.Equal(t, "11000011", bv.String())
}
