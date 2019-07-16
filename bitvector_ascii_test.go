package bitvector

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ascii(t *testing.T) {
	var bv Ascii

	// Set
	assert.Equal(t, Ascii{0, 0}, bv)
	bv = bv.Set(0, true)
	assert.Equal(t, Ascii{1, 0}, bv)
	bv = bv.Set(127, true)
	assert.Equal(t, Ascii{1, 0x8000000000000000}, bv)
	bv = bv.Set(128, true)
	assert.Equal(t, Ascii{1, 0x8000000000000000}, bv)
	fmt.Printf("%v\n", bv)
	bv = bv.Set(126, true)
	fmt.Printf("%v\n", bv)
	assert.Equal(t, Ascii{1, 0xc000000000000000}, bv)
	bv = bv.Set(126, false)
	fmt.Printf("%v\n", bv)
	assert.Equal(t, Ascii{1, 0x8000000000000000}, bv)
	//
	//// Count
	//assert.Equal(t, uint8(2), bv.Count())
	//
	//// Get
	//assert.True(t, bv.Get(0))
	//assert.True(t, bv.Get(127))
	//assert.False(t, bv.Get(126))

	// Toggle
	//bv = Ascii{0, 0}
	//bv = bv.Toggle(127)
	//assert.Equal(t, Ascii{0, 9223372036854775808}, bv)
	//
	//// Clear
	//bv = 0
	//var i uint8
	//for ; i < 64; i++ {
	//	bv = bv.Set(i, true)
	//}
	//bv = bv.Clear(6, 2)
	//assert.Equal(t, Ascii(math.MaxUint64), bv)
	//bv = bv.Clear(2, 6)
	//assert.Equal(t, Ascii(18446744073709551555), bv)
	//
	//// String
	//assert.Equal(t, "1111111111111111111111111111111111111111111111111111111111000011", bv.String())
}
