package day11

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestZeroEngravedStoneOnBlink(t *testing.T) {
	var stone = Stone(0)

	var leftStone, rightStone = stone.OnBlink()

	assert.Equal(t, Stone(1), leftStone)
	assert.Nil(t, rightStone)
}

func TestEvenNumberOfDigitsEngravedStoneOnBlink(t *testing.T) {
	var testCases = []struct {
		engravedNumber int64
		expectedLeft   int64
		expectedRight  int64
	}{
		{2024, 20, 24},
		{96, 9, 6},
		{83794560, 8379, 4560},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			var stone = Stone(testCase.engravedNumber)

			var leftStone, rightStone = stone.OnBlink()

			assert.Equal(t, Stone(testCase.expectedLeft), leftStone)
			assert.Equal(t, Stone(testCase.expectedRight), *rightStone)
		})
	}
}

func TestEvenNumberOfDigitsDoNotKeepExtraLeadingZeroesOnBlink(t *testing.T) {
	var stone = Stone(1001)

	var leftStone, rightStone = stone.OnBlink()

	assert.Equal(t, Stone(10), leftStone)
	assert.Equal(t, Stone(1), *rightStone)
}

func TestMultiplyBy2024OddNumberOfDigitsOnBlink(t *testing.T) {
	var testCases = []struct {
		stone Stone
		expectedLeft   Stone
	}{
		{1, 1 * 2024},
		{2, 2 * 2024},
		{89322, 89322 * 2024},
		{999, 2021976},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			var leftStone, rightStone = testCase.stone.OnBlink()

			assert.Equal(t, testCase.expectedLeft, leftStone)
			assert.Nil(t, rightStone)
		})
	}
}
