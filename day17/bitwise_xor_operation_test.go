package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoZeros(t *testing.T) {
	assert.Equal(t, int64(0), BitwiseXor(0, 0))
}

func TestZeroAndOne(t *testing.T) {
	assert.Equal(t, int64(1), BitwiseXor(0, 1))
	assert.Equal(t, int64(1), BitwiseXor(1, 0))
}

func TestTwoOne(t *testing.T) {
	assert.Equal(t, int64(0), BitwiseXor(1, 1))
}

func TestProvidedExamples(t *testing.T) {
	assert.Equal(t, int64(26), BitwiseXor(29, 7))
}
