package bitvector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bv8(t *testing.T) {
	bv := &bitVector8{}

	// Set
	assert.Equal(t, uint8(0), bv.n)
	bv.Set(0, true)
	assert.Equal(t, uint8(1), bv.n)
	bv.Set(7, true)
	assert.Equal(t, uint8(129), bv.n)
	bv.Set(8, true)
	assert.Equal(t, uint8(129), bv.n)
	bv.Set(6, true)
	assert.Equal(t, uint8(193), bv.n)
	bv.Set(6, false)
	assert.Equal(t, uint8(129), bv.n)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(7))
	assert.False(t, bv.Get(6))

	// Reset
	bv.Reset()
	assert.Equal(t, uint8(0), bv.n)

	// Toggle
	bv.Toggle(7)
	assert.Equal(t, uint8(128), bv.n)

	// Clear
	bv.Reset()
	var i uint8
	for ; i < 8; i++ {
		bv.Set(i, true)
	}
	bv.Clear(6, 2)
	assert.Equal(t, uint8(math.MaxUint8), bv.n)
	bv.Clear(2, 6)
	assert.Equal(t, uint8(195), bv.n)

	// Copy
	assert.Equal(t, uint8(4), bv.Copy().Count())

	// String
	assert.Equal(t, "11000011", bv.String())

	// New
	assert.Equal(t, uint8(0), New8().Count())
}
