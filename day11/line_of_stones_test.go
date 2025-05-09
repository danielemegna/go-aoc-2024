package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineOfZeroStonesOnBlink(t *testing.T) {
	var stones LineOfStones = LineOfStones{0: 3}

	var actual = stones.OnBlink()

	var expected = LineOfStones{1: 3}
	assert.Equal(t, expected, actual)
}

func TestLineOfOddDigitsNumbersStonesOnBlink(t *testing.T) {
	var stones = LineOfStones{1: 1, 9: 1, 146: 1}

	var actual = stones.OnBlink()

	var expected = LineOfStones{
		2024:       1,
		9 * 2024:   1,
		146 * 2024: 1,
	}
	assert.Equal(t, expected, actual)
}

func TestLineOfEvenDigitsNumbersStonesOnBlink(t *testing.T) {
	var stones = LineOfStones{10: 1, 46: 1, 2823: 1}

	var actual = stones.OnBlink()

	var expected = LineOfStones{
		1: 1, 0: 1,
		4: 1, 6: 1,
		28: 1, 23: 1,
	}
	assert.Equal(t, expected, actual)
}

func TestPreserveDuplicatedStonesOnBlink(t *testing.T) {
	var stones = LineOfStones{10: 2, 287: 3, 9871: 4}

	var actual = stones.OnBlink()

	var expected = LineOfStones{
		1: 2, 0: 2,
		287 * 2024: 3,
		98: 4, 71: 4,
	}
	assert.Equal(t, expected, actual)
}

func TestSizeIsSumOfOccurrences(t *testing.T) {
	var stones = LineOfStones{1: 6, 9: 4, 146: 2}
	var actual = stones.Size()
	assert.Equal(t, 6+4+2, actual)
}

func TestAddNonPresentStones(t *testing.T) {
	var stones = LineOfStones{}
	stones.Add(Stone(1), 6)
	assert.Equal(t, LineOfStones{1: 6}, stones)
	stones.Add(Stone(146), 2)
	assert.Equal(t, LineOfStones{1: 6, 146: 2}, stones)
	stones.Add(Stone(19), 2)
	assert.Equal(t, LineOfStones{1: 6, 146: 2, 19: 2}, stones)
}

func TestAddAlreadyPresentStonesIncreaseOccurrences(t *testing.T) {
	var stones = LineOfStones{1: 6, 9: 4, 146: 2}
	stones.Add(Stone(9), 2)
	assert.Equal(t, LineOfStones{1: 6, 9: 4 + 2, 146: 2}, stones)
	stones.Add(Stone(146), 3)
	assert.Equal(t, LineOfStones{1: 6, 9: 4 + 2, 146: 2 + 3}, stones)
}
