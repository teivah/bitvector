package bitvector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bv16(t *testing.T) {
	bv := &bitVector16{}

	// Set
	assert.Equal(t, uint16(0), bv.n)
	bv.Set(0, true)
	assert.Equal(t, uint16(1), bv.n)
	bv.Set(15, true)
	assert.Equal(t, uint16(32769), bv.n)
	bv.Set(16, true)
	assert.Equal(t, uint16(32769), bv.n)
	bv.Set(14, true)
	assert.Equal(t, uint16(49153), bv.n)
	bv.Set(14, false)
	assert.Equal(t, uint16(32769), bv.n)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(15))
	assert.False(t, bv.Get(14))

	// Reset
	bv.Reset()
	assert.Equal(t, uint16(0), bv.n)

	// Toggle
	bv.Toggle(15)
	assert.Equal(t, uint16(32768), bv.n)

	// Clear
	bv.Reset()
	var i uint8
	for ; i < 16; i++ {
		bv.Set(i, true)
	}
	bv.Clear(6, 2)
	assert.Equal(t, uint16(math.MaxUint16), bv.n)
	bv.Clear(2, 6)
	assert.Equal(t, uint16(65475), bv.n)

	// Copy
	assert.Equal(t, uint8(12), bv.Copy().Count())

	// String
	assert.Equal(t, "1111111111000011", bv.String())

	// New
	assert.Equal(t, uint8(0), New16().Count())
}
