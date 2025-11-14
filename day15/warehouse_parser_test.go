package day15

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestWarehouseParser_Parse(t *testing.T) {
	// Simple input for testing
	input := `###
#@#
#O#
###

<>v^`

	parser := NewWarehouseParser()
	warehouseMap, moves := parser.Parse(input)
	
	// Check the map dimensions
	assert.Equal(t, 4, len(warehouseMap.Grid))
	assert.Equal(t, 3, len(warehouseMap.Grid[0]))
	
	// Check robot position
	assert.Equal(t, Coordinate{Row: 1, Col: 1}, warehouseMap.RobotPos)
	
	// Check box position
	boxPositions := warehouseMap.GetBoxPositions()
	assert.Len(t, boxPositions, 1)
	assert.Equal(t, Coordinate{Row: 2, Col: 1}, boxPositions[0])
	
	// Check moves
	assert.Equal(t, []rune{'<', '>', 'v', '^'}, moves)
}

func TestWarehouseParser_ParseWithProvidedExample(t *testing.T) {
	var input = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	
	parser := NewWarehouseParser()
	warehouseMap, moves := parser.Parse(input)
	
	// Check dimensions
	assert.Equal(t, 8, len(warehouseMap.Grid))
	assert.Equal(t, 8, len(warehouseMap.Grid[0]))
	
	// Check robot position
	assert.Equal(t, Coordinate{Row: 2, Col: 2}, warehouseMap.RobotPos)
	
	// Check number of boxes
	boxPositions := warehouseMap.GetBoxPositions()
	assert.Len(t, boxPositions, 6)
	
	// Check the moves
	expectedMoves := []rune{'<', '^', '^', '>', '>', '>', 'v', 'v', '<', 'v', '>', '>', 'v', '<', '<'}
	assert.Equal(t, expectedMoves, moves)
}