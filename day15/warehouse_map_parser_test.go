package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSmallerProvidedExampleWarehouseMap(t *testing.T) {
	var warehouseMap, moves = ParseWarehouseMapAndMoves(simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES))

	var expectedMap = expectedSmallerProvidedExampleParsedMap()
	var expectedMoves = expectedSmallerProvidedExampleParsedMoves()
	assert.Equal(t, expectedMap, warehouseMap)
	assert.Equal(t, expectedMoves, moves)
}

func TestParseLargerProvidedExampleWarehouseMap(t *testing.T) {
	var warehouseMap, moves = ParseWarehouseMapAndMoves(simulateFileContent(LARGER_PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 8, warehouseMap.GetWidth())
	assert.Equal(t, 8, warehouseMap.GetHeight())
	assert.Equal(t, Coordinate{3, 3}, warehouseMap.GetRobotPosition())
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 0}))
	assert.Equal(t, WALL, warehouseMap.ElementAt(Coordinate{1, 4}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{7, 7}))
	assert.Len(t, moves, 700)
	assert.Equal(t, LEFT, moves[0])
	assert.Equal(t, DOWN, moves[1])
	assert.Equal(t, UP, moves[699])
}

func TestParseLargerProvidedExampleWarehouseMapInDoubleWide(t *testing.T) {
	var warehouseMap, moves = ParseWarehouseMapAndMovesInDoubleWide(simulateFileContent(LARGER_PROVIDED_EXAMPLE_INPUT_LINES))

	var expectedMap = expectedLargerProvidedExampleParsedMapInDoubleWide()
	assert.Equal(t, expectedMap, warehouseMap)
	assert.Len(t, moves, 700)
	assert.Equal(t, LEFT, moves[0])
	assert.Equal(t, RIGHT, moves[46])
	assert.Equal(t, UP, moves[699])
}

func expectedSmallerProvidedExampleParsedMap() WarehouseMap {
	return WarehouseMap{
		{EMPTY, EMPTY, BOX, EMPTY, BOX, EMPTY},
		{WALL, ROBOT, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, WALL, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}
}

func expectedSmallerProvidedExampleParsedMoves() []Direction {
	return []Direction{
		LEFT, UP, UP, RIGHT, RIGHT, RIGHT, DOWN, DOWN,
		LEFT, DOWN, RIGHT, RIGHT, DOWN, LEFT, LEFT,
	}
}

func expectedLargerProvidedExampleParsedMapInDoubleWide() WarehouseMap {
	/*
	####################
	##....[]....[]..[]##
	##............[]..##
	##..[][]....[]..[]##
	##....[]@.....[]..##
	##[]##....[]......##
	##[]....[]....[]..##
	##..[][]..[]..[][]##
	##........[]......##
	####################
	*/
	return WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, LBOX, RBOX},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, LBOX, RBOX, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, LBOX, RBOX},
		{EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY},
		{LBOX, RBOX, WALL, WALL, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, LBOX, RBOX, LBOX, RBOX, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, LBOX, RBOX, LBOX, RBOX},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}
}
