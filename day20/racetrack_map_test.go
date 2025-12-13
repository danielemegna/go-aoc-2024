package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var SMALL_RACETRACK_INPUT_LINES = []string{
	"######",
	"#S...#",
	"####.#",
	"##...#",
	"##E###",
	"######",
}

var ANOTHER_RACETRACK_INPUT_LINES = []string{
	"#######",
	"#....S#",
	"#.#####",
	"#.....#",
	"#####.#",
	"###E..#",
	"#######",
}

func TestParseSmallRacetrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(SMALL_RACETRACK_INPUT_LINES))

	assert.Equal(t, Coordinate{1, 1}, racetrackMap.RacetrackStart().Coordinate)
	assert.NotNil(t, racetrackMap.RacetrackStart().Next)
	assert.Equal(t, 0, racetrackMap.RacetrackStart().DistanceFromStart)
	assert.Equal(t, 1, racetrackMap.RacetrackStart().Next.DistanceFromStart)
	assert.Equal(t, 5, racetrackMap.RacetrackStart().Next.Next.Next.Next.Next.DistanceFromStart)
	assert.Equal(t, 8, calcRacetrackLength(racetrackMap))
}

func TestParseAnotherRacetrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(ANOTHER_RACETRACK_INPUT_LINES))

	assert.Equal(t, Coordinate{5, 1}, racetrackMap.RacetrackStart().Coordinate)
	assert.NotNil(t, racetrackMap.RacetrackStart().Next)
	assert.Equal(t, 0, racetrackMap.RacetrackStart().DistanceFromStart)
	assert.Equal(t, 4, racetrackMap.RacetrackStart().Next.Next.Next.Next.DistanceFromStart)
	assert.Equal(t, 14, calcRacetrackLength(racetrackMap))
}

func TestParseProvidedExampleRacetrack(t *testing.T) {
	var racetrackMap = ParseRacetrack(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, Coordinate{1, 3}, racetrackMap.RacetrackStart().Coordinate)
	assert.NotNil(t, racetrackMap.RacetrackStart().Next)
	assert.Equal(t, 0, racetrackMap.RacetrackStart().DistanceFromStart)
	assert.Equal(t, 6, racetrackMap.RacetrackStart().Next.Next.Next.Next.Next.Next.DistanceFromStart)
	assert.Equal(t, 84, calcRacetrackLength(racetrackMap))
}

func calcRacetrackLength(this Racetrack) int {
	var length = 0
	var element = &this.start
	for element.Next != nil {
		element = element.Next
		length++
	}
	return length
}
