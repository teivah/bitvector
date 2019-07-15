package bitvector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBV32(t *testing.T) {
	bv := &bitVector32{}

	// Set
	assert.Equal(t, uint32(0), bv.n)
	bv.Set(0, true)
	assert.Equal(t, uint32(1), bv.n)
	bv.Set(31, true)
	assert.Equal(t, uint32(2147483649), bv.n)
	bv.Set(32, true)
	assert.Equal(t, uint32(2147483649), bv.n)
	bv.Set(30, true)
	assert.Equal(t, uint32(3221225473), bv.n)
	bv.Set(30, false)
	assert.Equal(t, uint32(2147483649), bv.n)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(31))
	assert.False(t, bv.Get(30))

	// Reset
	bv.Reset()
	assert.Equal(t, uint32(0), bv.n)

	// Toggle
	bv.Toggle(31)
	assert.Equal(t, uint32(2147483648), bv.n)

	// Clear
	bv.Reset()
	var i uint8
	for ; i < 32; i++ {
		bv.Set(i, true)
	}
	bv.Clear(6, 2)
	assert.Equal(t, uint32(math.MaxUint32), bv.n)
	bv.Clear(2, 6)
	assert.Equal(t, uint32(4294967235), bv.n)

	// Copy
	assert.Equal(t, uint8(28), bv.Copy().Count())

	// String
	assert.Equal(t, "11111111111111111111111111000011", bv.String())

	// New
	assert.Equal(t, uint8(0), New32().Count())
}
