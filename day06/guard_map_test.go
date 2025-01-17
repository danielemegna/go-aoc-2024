package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseProducePanicWithoutAnyGuard(t *testing.T) {
	var mapRows = []string{
		"..",
		"..",
	}

	assert.PanicsWithValue(t, "Cannot find any guard in map!", func() {
		ParseGuardMap(mapRows)
	})
}

func TestParseSimpleGuardMapWithoutObstacles(t *testing.T) {
	var mapRows = []string{
		"...",
		".>.",
		"...",
	}

	var actual = ParseGuardMap(mapRows)

	var expected = GuardMap{
		size: 3,
		guard: Guard{
			position:         Coordinate{x: 1, y: 1},
			direction:        East,
			visitedPositions: []Coordinate{{x: 1, y: 1}},
		},
		obstacles: []Coordinate{},
	}
	assert.Equal(t, expected, actual)
}

func TestParseGuardMapWithSomeObstaclesMap(t *testing.T) {
	var mapRows = []string{
		".#..",
		"#...",
		".<.#",
		"....",
	}

	var actual = ParseGuardMap(mapRows)

	var expected = GuardMap{
		size: 4,
		guard: Guard{
			position:         Coordinate{x: 1, y: 2},
			direction:        West,
			visitedPositions: []Coordinate{{x: 1, y: 2}},
		},
		obstacles: []Coordinate{
			{x: 1, y: 0},
			{x: 0, y: 1},
			{x: 3, y: 2},
		},
	}
	assert.Equal(t, expected, actual)
}

func TestParseProvidedExampleGuardMap(t *testing.T) {
	var actual = ParseGuardMap(PROVIDED_EXAMPLE_INPUT_LINES)

	var expected = GuardMap{
		size: 10,
		guard: Guard{
			position:         Coordinate{x: 4, y: 6},
			direction:        North,
			visitedPositions: []Coordinate{{x: 4, y: 6}},
		},
		obstacles: []Coordinate{
			{x: 4, y: 0},
			{x: 9, y: 1},
			{x: 2, y: 3},
			{x: 7, y: 4},
			{x: 1, y: 6},
			{x: 8, y: 7},
			{x: 0, y: 8},
			{x: 6, y: 9},
		},
	}
	assert.Equal(t, expected, actual)
}

func TestGuardWalkStopOutOfMapBoundariesOnEast(t *testing.T) {
	var guardMap = ParseGuardMap([]string{
		"....",
		".>..",
		"....",
		"....",
	})

	var newMap = guardMap.GuardWalk()

	var expectedGuard = Guard{
		position:         Coordinate{x: 4, y: 1},
		direction:        East,
		visitedPositions: []Coordinate{{x: 1, y: 1}, {x: 2, y: 1}, {x: 3, y: 1}},
	}
	assert.Equal(t, expectedGuard, newMap.guard)
	assert.NotEqual(t, newMap.guard, guardMap.guard) // original map guard should not change
}

func TestGuardWalkStopReachingAnObstacleGoingEast(t *testing.T) {
	var guardMap = ParseGuardMap([]string{
		".....",
		">..#.",
		".....",
		".....",
		".....",
	})

	var newMap = guardMap.GuardWalk()

	var expectedGuard = Guard{
		position:         Coordinate{x: 2, y: 1},
		direction:        East,
		visitedPositions: []Coordinate{{x: 0, y: 1}, {x: 1, y: 1}, {x: 2, y: 1}},
	}
	assert.Equal(t, expectedGuard, newMap.guard)
}

func TestGuardWalkStopOutOfMapBoundariesOnNorth(t *testing.T) {
	var guardMap = ParseGuardMap([]string{
		"....",
		"....",
		"....",
		"..^.",
	})

	var newMap = guardMap.GuardWalk()

	var expectedGuard = Guard{
		position:         Coordinate{x: 2, y: -1},
		direction:        North,
		visitedPositions: []Coordinate{
			{x: 2, y: 3}, {x: 2, y: 2},
			{x: 2, y: 1}, {x: 2, y: 0},
		},
	}
	assert.Equal(t, expectedGuard, newMap.guard)
}

func TestGuardWalkStopReachingAnObstacleGoingSouth(t *testing.T) {
	var guardMap = ParseGuardMap([]string{
		"..v.",
		"....",
		"....",
		"..#.",
	})

	var newMap = guardMap.GuardWalk()

	var expectedGuard = Guard{
		position:         Coordinate{x: 2, y: 2},
		direction:        South,
		visitedPositions: []Coordinate{
			{x: 2, y: 0}, {x: 2, y: 1},
			{x: 2, y: 2},
		},
	}
	assert.Equal(t, expectedGuard, newMap.guard)
}
