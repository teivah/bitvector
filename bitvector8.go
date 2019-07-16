package bitvector

import (
	"fmt"
	"math"
)

// Len8 is a 8-bit vector
type Len8 uint8

func (bv Len8) String() string {
	return fmt.Sprintf("%b", bv)
}

// Clear bits from index i (included) to index j (excluded)
func (bv Len8) Clear(i, j uint8) Len8 {
	if i > j {
		return bv
	}
	return math.MaxUint8<<j | ((1<<i)-1)&bv
}

// Count the number of bits set to 1
func (bv Len8) Count() uint8 {
	var count uint8
	var index Len8 = 1
	var i uint8
	for {
		if bv&index != 0 {
			count++
		}
		index <<= 1
		i++
		if i == 8 {
			break
		}
	}
	return count
}

// Toggle ith bit
func (bv Len8) Toggle(i uint8) Len8 {
	return bv ^ 1<<i
}

// Get ith bit
func (bv Len8) Get(i uint8) bool {
	return (bv & (1 << i)) != 0
}

// Set ith bit
func (bv Len8) Set(i uint8, b bool) Len8 {
	var value Len8
	if b {
		value = 1
	}
	var mask Len8 = ^(1 << i)
	return (bv & mask) | (value << i)
}
