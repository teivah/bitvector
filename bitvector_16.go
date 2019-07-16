package bitvector

import (
	"fmt"
	"math"
)

// Len16 is a 16-bit vector
type Len16 uint16

func (bv Len16) String() string {
	return fmt.Sprintf("%016b", bv)
}

// Clear bits from index i (included) to index j (excluded)
func (bv Len16) Clear(i, j uint8) Len16 {
	if i > j {
		return bv
	}
	return (math.MaxUint16<<j | ((1 << i) - 1)) & bv
}

// Count the number of bits set to 1
func (bv Len16) Count() uint8 {
	var count uint8
	var index Len16 = 1
	var i uint8
	for {
		if bv&index != 0 {
			count++
		}
		index <<= 1
		i++
		if i == 16 {
			break
		}
	}
	return count
}

// Toggle ith bit
func (bv Len16) Toggle(i uint8) Len16 {
	return bv ^ 1<<i
}

// Get ith bit
func (bv Len16) Get(i uint8) bool {
	return (bv & (1 << i)) != 0
}

// Set ith bit
func (bv Len16) Set(i uint8, b bool) Len16 {
	var value Len16
	if b {
		value = 1
	}
	var mask Len16 = ^(1 << i)
	return (bv & mask) | (value << i)
}

// And operator
func (bv Len16) And(bv2 Len16) Len16 {
	return bv & bv2
}

// Or operator
func (bv Len16) Or(bv2 Len16) Len16 {
	return bv | bv2
}

// Xor operator
func (bv Len16) Xor(bv2 Len16) Len16 {
	return bv ^ bv2
}

// AndNot operator
func (bv Len16) AndNot(bv2 Len16) Len16 {
	return bv &^ bv2
}

// Push left shifts the bits
func (bv Len16) Push(i uint8) Len16 {
	return bv << i
}

// Pop right shifts the bits
func (bv Len16) Pop(i uint8) Len16 {
	return bv >> i
}
