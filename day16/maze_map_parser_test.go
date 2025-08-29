package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLittleExample(t *testing.T) {
	var mazeMap, reindeer, target = ParseMazeMap(simulateFileContent([]string{
		"######",
		"#.E..#",
		"#.#.##",
		"#..S.#",
		"#.####",
		"######",
	}))

	var expectedMap = MazeMap{
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, WALL, EMPTY, WALL},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, WALL, WALL, WALL},
	}
	assert.Equal(t, expectedMap, mazeMap)
	assert.Equal(t, Reindeer{Coordinate: Coordinate{2, 2}, Direction: RIGHT}, reindeer)
	assert.Equal(t, Coordinate{1, 0}, target)
}

func TestParseFirstProvidedExample(t *testing.T) {
	var mazeMap, reindeer, target = ParseMazeMap(simulateFileContent(FIRST_PROVIDED_EXAMPLE_INPUT_LINES))

	var expectedMap = MazeMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, WALL, EMPTY, WALL, WALL, WALL, EMPTY, WALL, EMPTY, WALL, WALL, WALL, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY, WALL, EMPTY, EMPTY, EMPTY, WALL, EMPTY},
		{EMPTY, WALL, WALL, WALL, EMPTY, WALL, WALL, WALL, WALL, WALL, EMPTY, WALL, EMPTY},
		{EMPTY, WALL, EMPTY, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY},
		{EMPTY, WALL, EMPTY, WALL, WALL, WALL, WALL, WALL, EMPTY, WALL, WALL, WALL, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY},
		{WALL, WALL, EMPTY, WALL, EMPTY, WALL, WALL, WALL, WALL, WALL, EMPTY, WALL, EMPTY},
		{EMPTY, EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY, WALL, EMPTY},
		{EMPTY, WALL, EMPTY, WALL, EMPTY, WALL, WALL, WALL, EMPTY, WALL, EMPTY, WALL, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY, WALL, EMPTY, WALL, EMPTY},
		{EMPTY, WALL, WALL, WALL, EMPTY, WALL, EMPTY, WALL, EMPTY, WALL, EMPTY, WALL, EMPTY},
		{EMPTY, EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY},
	}
	assert.Equal(t, expectedMap, mazeMap)
	assert.Equal(t, Reindeer{Coordinate: Coordinate{0, 12}, Direction: RIGHT}, reindeer)
	assert.Equal(t, Coordinate{12, 0}, target)
}
