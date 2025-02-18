package day07

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRecognizeTrueTheSumOfTwoOperands(t *testing.T) {
	var equation = Equation{
		operands: []int{1, 2},
		total:    1 + 2,
	}

	assert.True(t, CanBeTrue(equation))
}

func TestRecognizeTrueTheSumOfManyOperands(t *testing.T) {
	var equation = Equation{
		operands: []int{1, 4, 6, 10},
		total:    1 + 4 + 6 + 10,
	}

	assert.True(t, CanBeTrue(equation))
}

func TestRecognizeFalseANotSolvableEquation(t *testing.T) {
	var equation = Equation{
		operands: []int{10, 50},
		total:    8,
	}

	assert.False(t, CanBeTrue(equation))
}

func TestRecognizeTrueTheProductOfTwoOperands(t *testing.T) {
	var equation = Equation{
		operands: []int{5, 4},
		total:    5 * 4,
	}

	assert.True(t, CanBeTrue(equation))
}

func TestRecognizeTrueTheProductOfManyOperands(t *testing.T) {
	var equation = Equation{
		operands: []int{5, 4, 2, 9, 1},
		total:    1 * 2 * 4 * 5 * 9,
	}

	assert.True(t, CanBeTrue(equation))
}

func TestProvidedExamples(t *testing.T) {
	var testCases = []struct {
		equation     Equation
		shouldBeTrue bool
	}{
		{equation: Equation{operands: []int{10, 19}, total: 190}, shouldBeTrue: true},
		{equation: Equation{operands: []int{81, 40, 27}, total: 3267}, shouldBeTrue: true},
		{equation: Equation{operands: []int{17, 5}, total: 83}, shouldBeTrue: false},
		{equation: Equation{operands: []int{15, 6}, total: 156}, shouldBeTrue: false},
		{equation: Equation{operands: []int{6, 8, 6, 15}, total: 7290}, shouldBeTrue: false},
		{equation: Equation{operands: []int{16, 10, 13}, total: 161011}, shouldBeTrue: false},
		{equation: Equation{operands: []int{17, 8, 14}, total: 192}, shouldBeTrue: false},
		{equation: Equation{operands: []int{9, 7, 18, 13}, total: 21037}, shouldBeTrue: false},
		{equation: Equation{operands: []int{11, 6, 16, 20}, total: 292}, shouldBeTrue: true},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			assert.Equal(t, testCase.shouldBeTrue, CanBeTrue(testCase.equation))
		})
	}

}

func TestRecognizeFalseSpecialCase(t *testing.T) {
	// increasing the first operand by 1 the operation
	// become true but it is originally false

	var equation = Equation{operands: []int{3, 4}, total: 8}
	assert.False(t, CanBeTrue(equation))

	equation = Equation{operands: []int{3, 4}, total: 16}
	assert.False(t, CanBeTrue(equation))
}
