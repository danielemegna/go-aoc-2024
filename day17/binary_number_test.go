package day17

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestBinaryFromInt(t *testing.T) {
	var testCases = []struct {
		number   int
		expected BinaryNumber
	}{
		{0, BinaryNumber{false}},
		{1, BinaryNumber{true}},
		{2, BinaryNumber{true, false}},
		{64, BinaryNumber{true, false, false, false, false, false, false}},
		{87, BinaryNumber{true, false, true, false, true, true, true}},
	}

	for index, testCase := range testCases {
		t.Run("TestCase #"+strconv.Itoa(index+1), func(t *testing.T) {
			var result = BinaryNumberFromInt(testCase.number)
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestBinaryToInt(t *testing.T) {
	var testCases = []struct {
		binary   BinaryNumber
		expected int
	}{
		{BinaryNumber{false}, 0},
		{BinaryNumber{true}, 1},
		{BinaryNumber{true, false}, 2},
		{BinaryNumber{true, false, false, false, false, false, false}, 64},
		{BinaryNumber{true, false, true, false, true, true, true}, 87},
		{BinaryNumber{true, false, true, false, true, false, true, true}, 171},
	}

	for index, testCase := range testCases {
		t.Run("TestCase #"+strconv.Itoa(index+1), func(t *testing.T) {
			var result = testCase.binary.ToInt()
			assert.Equal(t, testCase.expected, result)
		})
	}

}

func TestBothWaysConversion(t *testing.T) {
	var numbers = []int{0, 1, 3, 9, 12, 34, 53, 64, 72, 128, 240, 255, 500, 1025, 2022}
	for _, number := range numbers {
		t.Run(fmt.Sprintf("Both ways conversion %d", number), func(t *testing.T) {
			var binary = BinaryNumberFromInt(number)
			assert.Equal(t, binary.ToInt(), number)
		})
	}
}

func TestBinaryPadLeft(t *testing.T) {
	var binary = BinaryNumber{true, false}
	var result = binary.PadLeft(8)
	assert.Equal(t, BinaryNumber{false, false, false, false, false, false, true, false}, result)
}

func TestBinaryPadLeftDoNotTouchBigger(t *testing.T) {
	var binary = BinaryNumber{true, false, true, false, true, false, true, true}

	assert.Equal(t, binary, binary.PadLeft(8))
	assert.Equal(t, binary, binary.PadLeft(4))
}
