package bitvector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calc(t *testing.T) {
	bv64, bv32, bv16, bv8 := calc(1)
	assert.Equal(t, 0, bv64)
	assert.Equal(t, 0, bv32)
	assert.Equal(t, 0, bv16)
	assert.Equal(t, 1, bv8)

	bv64, bv32, bv16, bv8 = calc(2)
	assert.Equal(t, 0, bv64)
	assert.Equal(t, 0, bv32)
	assert.Equal(t, 0, bv16)
	assert.Equal(t, 1, bv8)

	bv64, bv32, bv16, bv8 = calc(16)
	assert.Equal(t, 0, bv64)
	assert.Equal(t, 0, bv32)
	assert.Equal(t, 1, bv16)
	assert.Equal(t, 0, bv8)

	bv64, bv32, bv16, bv8 = calc(248)
	assert.Equal(t, 3, bv64)
	assert.Equal(t, 1, bv32)
	assert.Equal(t, 1, bv16)
	assert.Equal(t, 1, bv8)
}
