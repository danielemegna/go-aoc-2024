package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWarehouse_Move_Up(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	warehouse, _ := parseWarehouse(fileContent)

	// Initial state
	assert.Equal(t, Coordinate{X: 2, Y: 2}, warehouse.Robot)
	assert.True(t, warehouse.Boxes[Coordinate{X: 3, Y: 1}])

	// Move up
	warehouse.Move('^')
	assert.Equal(t, Coordinate{X: 2, Y: 1}, warehouse.Robot)
	assert.False(t, warehouse.Boxes[Coordinate{X: 2, Y: 1}]) // This box is not here
}

func TestWarehouse_Move_Down(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	warehouse, _ := parseWarehouse(fileContent)
	warehouse.Robot = Coordinate{X: 2, Y: 3}

	warehouse.Move('v')
	assert.Equal(t, Coordinate{X: 2, Y: 4}, warehouse.Robot)
}

func TestWarehouse_Move_Left_BlockedByWall(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	warehouse, _ := parseWarehouse(fileContent)

	warehouse.Move('<')
	assert.Equal(t, Coordinate{X: 2, Y: 2}, warehouse.Robot)
}

func TestWarehouse_Move_Right_PushBox(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	warehouse, _ := parseWarehouse(fileContent)
	warehouse.Robot = Coordinate{X: 2, Y: 1}
	warehouse.Boxes[Coordinate{X: 3, Y: 1}] = true

	warehouse.Move('>')
	assert.Equal(t, Coordinate{X: 3, Y: 1}, warehouse.Robot)
	assert.True(t, warehouse.Boxes[Coordinate{X: 4, Y: 1}])
	assert.False(t, warehouse.Boxes[Coordinate{X: 3, Y: 1}])
}
