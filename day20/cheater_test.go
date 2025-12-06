package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPossibileCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(SMALL_RACETRACK_INPUT_LINES)

	var expected = []Cheat{
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 6 - 2},
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 4 - 2},
	}
	assert.Equal(t, expected, PossibleCheatsIn(racetrackMap))
}
