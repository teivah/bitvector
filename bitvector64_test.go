package bitvector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bv64(t *testing.T) {
	bv := &bitVector64{}

	// Set
	assert.Equal(t, uint64(0), bv.n)
	bv.Set(0, true)
	assert.Equal(t, uint64(1), bv.n)
	bv.Set(63, true)
	assert.Equal(t, uint64(9223372036854775809), bv.n)
	bv.Set(64, true)
	assert.Equal(t, uint64(9223372036854775809), bv.n)
	bv.Set(62, true)
	assert.Equal(t, uint64(13835058055282163713), bv.n)
	bv.Set(62, false)
	assert.Equal(t, uint64(9223372036854775809), bv.n)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(63))
	assert.False(t, bv.Get(62))

	// Reset
	bv.Reset()
	assert.Equal(t, uint64(0), bv.n)

	// Toggle
	bv.Toggle(63)
	assert.Equal(t, uint64(9223372036854775808), bv.n)

	// Clear
	bv.Reset()
	var i uint8
	for ; i < 64; i++ {
		bv.Set(i, true)
	}
	bv.Clear(6, 2)
	assert.Equal(t, uint64(math.MaxUint64), bv.n)
	bv.Clear(2, 6)
	assert.Equal(t, uint64(18446744073709551555), bv.n)

	// Copy
	assert.Equal(t, uint8(60), bv.Copy().Count())

	// String
	assert.Equal(t, "1111111111111111111111111111111111111111111111111111111111000011", bv.String())

	// New
	assert.Equal(t, uint8(0), New64().Count())
}
