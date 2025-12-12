package day20

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcManhattanDistance(t *testing.T) {
	var testCases = []struct {
		from     Coordinate
		to       Coordinate
		expected int
	}{
		{from: Coordinate{0, 0}, to: Coordinate{0, 0}, expected: 0},
		{from: Coordinate{0, 0}, to: Coordinate{0, 1}, expected: 1},
		{from: Coordinate{1, 0}, to: Coordinate{0, 0}, expected: 1},
		{from: Coordinate{0, 0}, to: Coordinate{1, 1}, expected: 2},
		{from: Coordinate{12, 6}, to: Coordinate{0, 0}, expected: 18},
		{from: Coordinate{10, 7}, to: Coordinate{20, 6}, expected: 11},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, CalcManhattanDistance(testCase.from, testCase.to))
			assert.Equal(t, testCase.expected, CalcManhattanDistance(testCase.to, testCase.from))
		})
	}
}
