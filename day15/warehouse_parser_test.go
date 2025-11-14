package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWarehouse(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	warehouse, moves := parseWarehouse(fileContent)

	assert.Equal(t, 8, warehouse.Width)
	assert.Equal(t, 8, warehouse.Height)
	assert.Equal(t, "<^^>>>vv<v>>v<<", moves)
	assert.Equal(t, Coordinate{X: 2, Y: 2}, warehouse.Robot)
	assert.True(t, warehouse.Walls[Coordinate{X: 0, Y: 0}])
	assert.True(t, warehouse.Boxes[Coordinate{X: 3, Y: 1}])
}
