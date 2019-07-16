package bitvector

import (
	"fmt"
	"math"
)

// Len64 is a 64-bit vector
type Len64 uint64

func (bv Len64) String() string {
	return fmt.Sprintf("%064b", bv)
}

// Clear bits from index i (included) to index j (excluded)
func (bv Len64) Clear(i, j uint8) Len64 {
	if i > j {
		return bv
	}
	return (math.MaxUint64<<j | ((1 << i) - 1)) & bv
}

// Count the number of bits set to 1
func (bv Len64) Count() uint8 {
	var count uint8
	var index Len64 = 1
	var i uint8
	for {
		if bv&index != 0 {
			count++
		}
		index <<= 1
		i++
		if i == 64 {
			break
		}
	}
	return count
}

// Toggle ith bit
func (bv Len64) Toggle(i uint8) Len64 {
	return bv ^ 1<<i
}

// Get ith bit
func (bv Len64) Get(i uint8) bool {
	return (bv & (1 << i)) != 0
}

// Set ith bit
func (bv Len64) Set(i uint8, b bool) Len64 {
	var value Len64
	if b {
		value = 1
	}
	var mask Len64 = ^(1 << i)
	return (bv & mask) | (value << i)
}

// And operator
func (bv Len64) And(bv2 Len64) Len64 {
	return bv & bv2
}

// Or operator
func (bv Len64) Or(bv2 Len64) Len64 {
	return bv | bv2
}

// Xor operator
func (bv Len64) Xor(bv2 Len64) Len64 {
	return bv ^ bv2
}

// AndNot operator
func (bv Len64) AndNot(bv2 Len64) Len64 {
	return bv &^ bv2
}

// Push left shifts the bits
func (bv Len64) Push(i uint8) Len64 {
	return bv << i
}

// Pop right shifts the bits
func (bv Len64) Pop(i uint8) Len64 {
	return bv >> i
}
