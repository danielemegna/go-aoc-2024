package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPossibileCheatsInSmallTrack(t *testing.T) {
	var racetrackMap = RacetrackMap{
		{START, 1,    2,    3},
		{WALL,  WALL, WALL, 4},
		{WALL,  7,    6,    5},
		{WALL,  8,    WALL, WALL},
	}

	var expected = []Cheat{
		{startPosition: Coordinate{1, 0}, endPosition: Coordinate{1, 2}, savingInPicoseconds: 6 - 2},
		{startPosition: Coordinate{2, 0}, endPosition: Coordinate{2, 2}, savingInPicoseconds: 4 - 2},
	}
	assert.Equal(t, expected, PossibleCheatsIn(racetrackMap))
}
