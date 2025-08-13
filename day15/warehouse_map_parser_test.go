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
