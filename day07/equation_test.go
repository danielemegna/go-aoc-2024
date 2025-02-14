package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTwoOperandsEquation(t *testing.T) {
	var actual = ParseEquation("190: 10 19")
	assert.Equal(t, Equation{operands: []int{10, 19}, total: 190}, actual)
}

func TestParseThreeOperandsEquation(t *testing.T) {
	var actual = ParseEquation("3267: 81 40 27")
	assert.Equal(t, Equation{operands: []int{81, 40, 27}, total: 3267}, actual)
}

func TestParseFourOperandsEquation(t *testing.T) {
	var actual = ParseEquation("21037: 9 7 18 13")
	assert.Equal(t, Equation{operands: []int{9, 7, 18, 13}, total: 21037}, actual)
}
