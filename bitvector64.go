package bitvector

import (
	"fmt"
	"math"
)

type bitVector64 struct {
	n uint64
}

func (bv *bitVector64) String() string {
	return fmt.Sprintf("%b", bv.n)
}

func (bv *bitVector64) Clear(i, j uint8) {
	if i > j {
		return
	}
	bv.n = math.MaxUint64<<j | ((1<<i)-1)&bv.n
}

func (bv *bitVector64) Copy() BitVector {
	return &bitVector64{
		n: bv.n,
	}
}

func (bv *bitVector64) Count() uint8 {
	var count uint8
	var index uint64 = 1
	var i uint8
	for {
		if bv.n&index != 0 {
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

func (bv *bitVector64) Toggle(i uint8) {
	bv.n ^= 1 << i
}

func (bv *bitVector64) Get(i uint8) bool {
	return (bv.n & (1 << i)) != 0
}

func (bv *bitVector64) Reset() {
	bv.n = 0
}

func (bv *bitVector64) Set(i uint8, b bool) {
	var value uint64
	if b {
		value = 1
	}
	var mask uint64 = ^(1 << i)
	bv.n = (bv.n & mask) | (value << i)
}

// New64 creates a new 64-bit vector
func New64() BitVector {
	return &bitVector64{}
}
