package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSmallerProvidedExampleWarehouseMap(t *testing.T) {
	var warehouseMap, moves = ParseWarehouseMapAndMoves(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)

	var expectedMap = smallerProvidedExampleMap()
	var expectedMoves = []Direction{
		LEFT, UP, UP, RIGHT, RIGHT, RIGHT, DOWN, DOWN,
		LEFT, DOWN, RIGHT, RIGHT, DOWN, LEFT, LEFT,
	}
	assert.Equal(t, expectedMap, warehouseMap)
	assert.Equal(t, expectedMoves, moves)
}

func TestParseLargerProvidedExampleWarehouseMap(t *testing.T) {
	var warehouseMap, moves = ParseWarehouseMapAndMoves(LARGER_PROVIDED_EXAMPLE_INPUT_LINES)

	assert.Equal(t, 8, warehouseMap.MapSize())
	assert.Equal(t, Coordinate{3, 3}, warehouseMap.GetRobotPosition())
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 0}))
	assert.Equal(t, WALL, warehouseMap.ElementAt(Coordinate{1, 4}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{7, 7}))
	assert.Len(t, moves, 700)
	assert.Equal(t, LEFT, moves[0])
	assert.Equal(t, DOWN, moves[1])
	assert.Equal(t, UP, moves[699])
}
