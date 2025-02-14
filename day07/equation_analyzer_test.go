package day07

import (
	"github.com/stretchr/testify/assert"
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
