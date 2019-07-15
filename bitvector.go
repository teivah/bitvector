package bitvector

import (
	"fmt"
	"strings"
)

// Handler represents an 8, 16, 32 or 64 bit vector
type Handler interface {
	fmt.Stringer
	// Clear bits from index i (included) to index j (excluded)
	Clear(i, j uint8) Handler
	// Count the number of bits set to 1
	Count() uint8
	// Get ith bit
	Get(i uint8) bool
	// Set ith bit
	Set(i uint8, b bool) Handler
	// Toggle ith bit
	Toggle(i uint8) Handler
}

type bitVector struct {
	len     int
	bv64    []uint64
	bv32    []uint32
	bv16    []uint16
	bv8     []uint8
	lenBv64 int
	lenBv32 int
	lenBv16 int
	lenBv8  int
}

func New(len int) Handler {
	len64, len32, len16, len8 := calc(len)

	var bv64 []uint64
	if len64 != 0 {
		bv64 = make([]uint64, len64)
	}

	var bv32 []uint32
	if len32 != 0 {
		bv32 = make([]uint32, len32)
	}

	var bv16 []uint16
	if len16 != 0 {
		bv16 = make([]uint16, len16)
	}

	var bv8 []uint8
	if len8 != 0 {
		bv8 = make([]uint8, len8)
	}

	return &bitVector{
		len:     len,
		bv64:    bv64,
		bv32:    bv32,
		bv16:    bv16,
		bv8:     bv8,
		lenBv64: len64,
		lenBv32: len32,
		lenBv16: len16,
		lenBv8:  len8,
	}
}

func calc(len int) (len64, len32, len16, len8 int) {
	len64 = len / 64
	len = len - len64*64

	len32 = len / 32
	len = len - len32*32

	len16 = len / 16
	len = len - len16*16

	if len != 0 {
		len8 = len / 8
		if len%8 != 0 {
			len8++
		}
	}
	return
}

func (bv *bitVector) String() string {
	sb := strings.Builder{}

	for _, b := range bv.bv64 {
		sb.WriteString(fmt.Sprintf("%b", b))
	}

	for _, b := range bv.bv32 {
		sb.WriteString(fmt.Sprintf("%b", b))
	}

	for _, b := range bv.bv16 {
		sb.WriteString(fmt.Sprintf("%b", b))
	}

	for i, b := range bv.bv8 {
		if bv.len < 16 {
			if i+1 == bv.lenBv8 {
				sb.WriteString(fmt.Sprintf("%08b", b
				[]))
			} else {
				sb.WriteString(fmt.Sprintf("%08b", b))
			}
		} else {
			sb.WriteString(fmt.Sprintf("%08b", b))
		}
	}

	return sb.String()
}

func (bv *bitVector) Clear(i, j uint8) Handler {
	panic("implement me")
}

func (bv *bitVector) Count() uint8 {
	panic("implement me")
}

func (bv *bitVector) Get(i uint8) bool {
	panic("implement me")
}

func (bv *bitVector) Set(i uint8, b bool) Handler {
	panic("implement me")
}

func (bv *bitVector) Toggle(i uint8) Handler {
	panic("implement me")
}
