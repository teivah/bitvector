package bitvector

import (
	"fmt"
	"math"
)

// Ascii is a 64-bit vector
type Ascii [2]uint64

func (bv Ascii) String() string {
	return fmt.Sprintf("%064b", bv[1]) + fmt.Sprintf("%064b", bv[0])
}

// Clear bits from index i (included) to index j (excluded)
func (bv Ascii) Clear(i, j uint8) Ascii {
	if i > j {
		return bv
	}
	if i <= 63 && j <= 63 {
		return Ascii{
			math.MaxUint64<<j | ((1<<i)-1)&bv[0],
			bv[1],
		}
	}
	if i >= 63 && j >= 64 {
		return Ascii{
			bv[0],
			math.MaxUint64<<j | ((1<<i)-1)&bv[1],
		}
	}
	return Ascii{
		0 | ((1<<i)-1)&bv[0],
		math.MaxUint64<<j | ((1<<64)-1)&bv[1],
	}
}

// Count the number of bits set to 1
func (bv Ascii) Count() uint8 {
	var count uint8
	var index uint64 = 1
	var i uint8
	for {
		if bv[0]&index != 0 {
			count++
		}
		index <<= 1
		i++
		if i == 64 {
			i = 1
			break
		}
	}
	for {
		if bv[1]&index != 0 {
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
func (bv Ascii) Toggle(i uint8) Ascii {
	if i <= 63 {
		return Ascii{
			bv[0] ^ 1<<i,
			bv[1],
		}
	}
	return Ascii{
		bv[0],
		bv[1] ^ 1<<i,
	}
}

// Get ith bit
func (bv Ascii) Get(i uint8) bool {
	if i <= 63 {
		return (bv[0] & (1 << i)) != 0
	}
	return (bv[1] & (1 << i)) != 0
}

// Set ith bit
func (bv Ascii) Set(i uint8, b bool) Ascii {
	var value uint64
	if b {
		value = 1
	}
	var mask uint64 = ^(1 << i)

	if i <= 63 {
		return Ascii{
			(bv[0] & mask) | (value << i),
			bv[1],
		}
	}

	return Ascii{
		bv[0],
		(bv[1] & mask) | (value << (i - 64)),
	}
}
