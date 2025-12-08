package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPossibileCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))

	var expected = []Cheat{
		{startPosition: Coordinate{2, 1}, endPosition: Coordinate{2, 3}, savingInPicoseconds: 8 - 2 - 2}, // 8 start, 2 end, 2 cheat duration
		{startPosition: Coordinate{3, 1}, endPosition: Coordinate{3, 3}, savingInPicoseconds: 7 - 3 - 2}, // 7 start, 3 end, 2 cheat duration
	}
	assert.ElementsMatch(t, expected, PossibleCheatsIn(racetrackMap))
}

// TEST: map with cheat starting from track start position ?
// TEST: cheat with L-angle move saving some picoseconds ? (>0)
