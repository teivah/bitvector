package bitvector

import (
	"fmt"
)

// Handler represents an 8, 16, 32 or 64 bit vector
type Handler interface {
	fmt.Stringer
	// Clear bits from index i (included) to index j (excluded)
	Clear(i, j uint8)
	// Copy creates another bit vector structure based on the same length
	Copy() Handler
	// Count the number of bits set to 1
	Count() uint8
	// Get ith bit
	Get(i uint8) bool
	// Reset the bit vector to 0
	Reset()
	// Set ith bit
	Set(i uint8, b bool)
	// Toggle ith bit
	Toggle(i uint8)
}
