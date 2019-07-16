package bitvector

// Ascii is a 64-bit vector
type Ascii [2]Len64

func NewAscii() Ascii {
	return Ascii{}
}

func (bv Ascii) String() string {
	return bv[1].String() + bv[0].String()
}

// Clear bits from index i (included) to index j (excluded)
func (bv Ascii) Clear(i, j uint8) Ascii {
	if i > j {
		return bv
	}
	if i <= 63 && j <= 63 {
		return Ascii{
			bv[0].Clear(i, j),
			bv[1],
		}
	}
	if i >= 64 && j >= 64 {
		return Ascii{
			bv[0],
			bv[1].Clear(i-64, j-64),
		}
	}
	return Ascii{
		bv[0].Clear(i, 64),
		bv[1].Clear(0, j-64),
	}
}

// Count the number of bits set to 1
func (bv Ascii) Count() uint8 {
	return bv[0].Count() + bv[1].Count()
}

// Toggle ith bit
func (bv Ascii) Toggle(i uint8) Ascii {
	if i <= 63 {
		return Ascii{
			bv[0].Toggle(i),
			bv[1],
		}
	}
	return Ascii{
		bv[0],
		bv[1].Toggle(i - 64),
	}
}

// Get ith bit
func (bv Ascii) Get(i uint8) bool {
	if i <= 63 {
		return bv[0].Get(i)
	}
	return bv[1].Get(i - 64)
}

// Set ith bit
func (bv Ascii) Set(i uint8, b bool) Ascii {
	if i <= 63 {
		return Ascii{
			bv[0].Set(i, b),
			bv[1],
		}
	}
	return Ascii{
		bv[0],
		bv[1].Set(i-64, b),
	}
}

// And operator
func (bv Ascii) And(bv2 Ascii) Ascii {
	return Ascii{
		bv[0] & bv2[0],
		bv[1] & bv2[1],
	}
}

// Or operator
func (bv Ascii) Or(bv2 Ascii) Ascii {
	return Ascii{
		bv[0] | bv2[0],
		bv[1] | bv2[1],
	}
}

// Xor operator
func (bv Ascii) Xor(bv2 Ascii) Ascii {
	return Ascii{
		bv[0] ^ bv2[0],
		bv[1] ^ bv2[1],
	}
}

// AndNot operator
func (bv Ascii) AndNot(bv2 Ascii) Ascii {
	return Ascii{
		bv[0] &^ bv2[0],
		bv[1] & bv2[1],
	}
}
