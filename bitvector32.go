package bitvector

import (
	"fmt"
	"math"
)

// Len32 is a 32-bit vector
type Len32 uint32

func (bv Len32) String() string {
	return fmt.Sprintf("%b", bv)
}

// Clear bits from index i (included) to index j (excluded)
func (bv Len32) Clear(i, j uint8) Len32 {
	if i > j {
		return bv
	}
	return math.MaxUint32<<j | ((1<<i)-1)&bv
}

// Count the number of bits set to 1
func (bv Len32) Count() uint8 {
	var count uint8
	var index Len32 = 1
	var i uint8
	for {
		if bv&index != 0 {
			count++
		}
		index <<= 1
		i++
		if i == 32 {
			break
		}
	}
	return count
}

// Toggle ith bit
func (bv Len32) Toggle(i uint8) Len32 {
	return bv ^ 1<<i
}

// Get ith bit
func (bv Len32) Get(i uint8) bool {
	return (bv & (1 << i)) != 0
}

// Set ith bit
func (bv Len32) Set(i uint8, b bool) Len32 {
	var value Len32
	if b {
		value = 1
	}
	var mask Len32 = ^(1 << i)
	return (bv & mask) | (value << i)
}
