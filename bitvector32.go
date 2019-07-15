package bitvector

import (
	"fmt"
	"math"
)

type bitVector32 struct {
	n uint32
}

func (bv *bitVector32) String() string {
	return fmt.Sprintf("%b", bv.n)
}

func (bv *bitVector32) Clear(i, j uint8) {
	if i > j {
		return
	}
	bv.n = math.MaxUint32<<j | ((1<<i)-1)&bv.n
}

func (bv *bitVector32) Copy() BitVector {
	return &bitVector32{
		n: bv.n,
	}
}

func (bv *bitVector32) Count() uint8 {
	var count uint8
	var index uint32 = 1
	var i uint8
	for {
		if bv.n&index != 0 {
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

func (bv *bitVector32) Toggle(i uint8) {
	bv.n ^= 1 << i
}

func (bv *bitVector32) Get(i uint8) bool {
	return (bv.n & (1 << i)) != 0
}

func (bv *bitVector32) Reset() {
	bv.n = 0
}

func (bv *bitVector32) Set(i uint8, b bool) {
	var value uint32
	if b {
		value = 1
	}
	var mask uint32 = ^(1 << i)
	bv.n = (bv.n & mask) | (value << i)
}

// New32 creates a new 32-bit vector
func New32() BitVector {
	return &bitVector32{}
}
