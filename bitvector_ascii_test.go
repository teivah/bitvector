package bitvector_test

import (
	"github.com/teivah/bitvector"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ascii(t *testing.T) {
	var bv bitvector.Ascii

	// Set
	assert.Equal(t, bitvector.Ascii{0, 0}, bv)
	bv = bv.Set(0, true)
	assert.Equal(t, bitvector.Ascii{1, 0}, bv)
	bv = bv.Set(127, true)
	assert.Equal(t, bitvector.Ascii{1, 0x8000000000000000}, bv)
	bv = bv.Set(128, true)
	assert.Equal(t, bitvector.Ascii{1, 0x8000000000000000}, bv)
	bv = bv.Set(126, true)
	assert.Equal(t, bitvector.Ascii{1, 0xc000000000000000}, bv)
	bv = bv.Set(126, false)
	assert.Equal(t, bitvector.Ascii{1, 0x8000000000000000}, bv)

	// Count
	assert.Equal(t, uint8(2), bv.Count())

	// Get
	assert.True(t, bv.Get(0))
	assert.True(t, bv.Get(127))
	assert.False(t, bv.Get(126))

	// Toggle
	bv = bitvector.NewAscii()
	bv = bv.Toggle(61)
	assert.Equal(t, bitvector.Ascii{0x2000000000000000, 0x0}, bv)
	bv = bv.Toggle(127)
	assert.Equal(t, bitvector.Ascii{0x2000000000000000, 9223372036854775808}, bv)

	// Clear
	bv = bitvector.NewAscii()
	var i uint8
	for ; i < 128; i++ {
		bv = bv.Set(i, true)
	}
	assert.Equal(t, bitvector.Ascii{math.MaxUint64, math.MaxUint64}, bv)
	bv = bv.Clear(6, 2)
	assert.Equal(t, bitvector.Ascii{math.MaxUint64, math.MaxUint64}, bv)
	bv = bv.Clear(2, 6)
	assert.Equal(t, bitvector.Ascii{0xffffffffffffffc3, 0xffffffffffffffff}, bv)
	bv = bv.Clear(125, 127)
	assert.Equal(t, bitvector.Ascii{0xffffffffffffffc3, 0x9fffffffffffffff}, bv)
	bv = bv.Clear(62, 67)
	assert.Equal(t, bitvector.Ascii{0x3fffffffffffffc3, 0x9ffffffffffffff8}, bv)

	// String
	assert.Equal(t, "10011111111111111111111111111111111111111111111111111111111110000011111111111111111111111111111111111111111111111111111111000011", bv.String())
}
