package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			position:  Coordinate{x: 1, y: 1},
			direction: East,
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
			position:  Coordinate{x: 1, y: 2},
			direction: West,
		},
		obstacles: []Coordinate{
			// WIP TODO obstacles!
/* 			{x: 1, y: 0},
			{x: 0, y: 1},
			{x: 3, y: 2}, */
		},
	}
	assert.Equal(t, expected, actual)
}

func TestParseProvidedExampleGuardMap(t *testing.T) {
	var actual = ParseGuardMap(PROVIDED_EXAMPLE_INPUT_LINES)

	var expected = GuardMap{
		size: 10,
		guard: Guard{
			position:  Coordinate{x: 4, y: 6},
			direction: North,
		},
		obstacles: []Coordinate{},
	}
	assert.Equal(t, expected, actual)
}
