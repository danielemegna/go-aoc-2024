package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPossibileCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))

	var expected = []Cheat{
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 6 - 2},
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 4 - 2},
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{4, 2}, savingInPicoseconds: 4 - 4},
		{startPosition: Coordinate{4, 2}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 6 - 6},
		{startPosition: Coordinate{3, 3}, endPosition: Coordinate{2, 4}, savingInPicoseconds: 8 - 8},
	}
	assert.ElementsMatch(t, expected, PossibleCheatsIn(racetrackMap))
}

// TEST: map with cheat starting from track start position ?
// TEST: cheat with L-angle move saving some picoseconds (>0)
