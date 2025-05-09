package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineOfZeroStonesOnBlink(t *testing.T) {
	var stones LineOfStones = LineOfStones{0: 3}

	var actual = stones.OnBlink()

	var expected = LineOfStones{1:3}
	assert.Equal(t, expected, actual)
}

func TestLineOfOddDigitsNumbersStonesOnBlink(t *testing.T) {
	var stones = LineOfStones{1:1, 9:1, 146:1}

	var actual = stones.OnBlink()

	var expected = LineOfStones{2024:1, 9 * 2024:1, 146 * 2024:1}
	assert.Equal(t, expected, actual)
}

func TestLineOfEvenDigitsNumbersStonesOnBlink(t *testing.T) {
	var stones = LineOfStones{10:1, 46:1, 2823:1}

	var actual = stones.OnBlink()

	var expected = LineOfStones{
		1:1, 0:1,
		4:1, 6:1,
		28:1, 23:1,
	}
	assert.Equal(t, expected, actual)
}
