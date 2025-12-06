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
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))

	assert.Equal(t, Coordinate{1, 1}, racetrackMap.StartCoordinate())
	assert.Equal(t, 8, racetrackMap.Length())
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{0,0}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{3,2}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{5,5}))
	assert.Equal(t, START, racetrackMap.ValueAt(Coordinate{1,1}))
	assert.Equal(t, 1, racetrackMap.ValueAt(Coordinate{2,1}))
	assert.Equal(t, 2, racetrackMap.ValueAt(Coordinate{3,1}))
	assert.Equal(t, 7, racetrackMap.ValueAt(Coordinate{2,3}))
	assert.Equal(t, 8, racetrackMap.ValueAt(Coordinate{2,4}))
}

func TestParseProvidedExampleRacetrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, Coordinate{1, 3}, racetrackMap.StartCoordinate())
	assert.Equal(t, 84, racetrackMap.Length())
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{0,0}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{5,4}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{10,7}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{8,10}))
	assert.Equal(t, WALL, racetrackMap.ValueAt(Coordinate{14,14}))
	assert.Equal(t, START, racetrackMap.ValueAt(Coordinate{1,3}))
	assert.Equal(t, 1, racetrackMap.ValueAt(Coordinate{1,2}))
	assert.Equal(t, 2, racetrackMap.ValueAt(Coordinate{1,1}))
	assert.Equal(t, 2, racetrackMap.ValueAt(Coordinate{1,1}))
	assert.Equal(t, 59, racetrackMap.ValueAt(Coordinate{8,9}))
	assert.Equal(t, 83, racetrackMap.ValueAt(Coordinate{4,7}))
	assert.Equal(t, 84, racetrackMap.ValueAt(Coordinate{5,7}))
}
