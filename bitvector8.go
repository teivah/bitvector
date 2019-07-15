package bitvector

import (
	"fmt"
	"math"
)

type bitVector8 struct {
	n uint8
}

func (bv *bitVector8) String() string {
	return fmt.Sprintf("%b", bv.n)
}

func (bv *bitVector8) Clear(i, j uint8) {
	if i > j {
		return
	}
	bv.n = math.MaxUint8<<j | ((1<<i)-1)&bv.n
}

func (bv *bitVector8) Copy() Handler {
	return &bitVector8{
		n: bv.n,
	}
}

func (bv *bitVector8) Count() uint8 {
	var count uint8
	var index uint8 = 1
	var i uint8
	for {
		if bv.n&index != 0 {
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

func (bv *bitVector8) Toggle(i uint8) {
	bv.n ^= 1 << i
}

func (bv *bitVector8) Get(i uint8) bool {
	return (bv.n & (1 << i)) != 0
}

func (bv *bitVector8) Reset() {
	bv.n = 0
}

func (bv *bitVector8) Set(i uint8, b bool) {
	var value uint8
	if b {
		value = 1
	}
	var mask uint8 = ^(1 << i)
	bv.n = (bv.n & mask) | (value << i)
}

// New8 creates a new 8-bit vector
func New8() Handler {
	return &bitVector8{}
}
