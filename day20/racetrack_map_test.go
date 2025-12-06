package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SMALL_RACETRACK_INPUT_LINES = []string{
	"######",
	"#S...#",
	"####.#",
	"##...#",
	"##E###",
	"######",
}

func TestParseSmallRacetrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(SMALL_RACETRACK_INPUT_LINES)

	assert.Equal(t, Coordinate{1, 1}, racetrackMap.StartCoordinate())
	assert.Equal(t, 8, racetrackMap.Length())
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{0,0}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{3,2}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{5,5}))
	assert.Equal(t, 0, racetrackMap.ValueAt(Coordinate{1,1})) // START has distance 0
	assert.Equal(t, 1, racetrackMap.ValueAt(Coordinate{2,1}))
	assert.Equal(t, 2, racetrackMap.ValueAt(Coordinate{3,1}))
}
