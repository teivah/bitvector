package bitvector

import (
	"fmt"
	"math"
)

type bitVector16 struct {
	n uint16
}

func (bv *bitVector16) String() string {
	return fmt.Sprintf("%b", bv.n)
}

func (bv *bitVector16) Clear(i, j uint8) {
	if i > j {
		return
	}
	bv.n = math.MaxUint16<<j | ((1<<i)-1)&bv.n
}

func (bv *bitVector16) Copy() BitVector {
	return &bitVector16{
		n: bv.n,
	}
}

func (bv *bitVector16) Count() uint8 {
	var count uint8
	var index uint16 = 1
	var i uint8
	for {
		if bv.n&index != 0 {
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

func (bv *bitVector16) Toggle(i uint8) {
	bv.n ^= 1 << i
}

func (bv *bitVector16) Get(i uint8) bool {
	return (bv.n & (1 << i)) != 0
}

func (bv *bitVector16) Reset() {
	bv.n = 0
}

func (bv *bitVector16) Set(i uint8, b bool) {
	var value uint16
	if b {
		value = 1
	}
	var mask uint16 = ^(1 << i)
	bv.n = (bv.n & mask) | (value << i)
}

// New16 creates a new 16-bit vector
func New16() BitVector {
	return &bitVector16{}
}
