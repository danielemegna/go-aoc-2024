package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperation(t *testing.T) {
	var testCases = []struct {
		name      string
		numerator int
		operand   int
		expected  int
	}{
		{"easy case, no rounding", 8, 2, 2},
		{"easy case, bigger number", 2024, 1, 1012},
		{"round floor, 1 as operand", 253, 1, 126},
		{"round floor, 2 as operand", 126, 2, 31},
		{"round floor, more complex result", 31, 3, 3},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var result = AdvOperation(testCase.numerator, testCase.operand)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
