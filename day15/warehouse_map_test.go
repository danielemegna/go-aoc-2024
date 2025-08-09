package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBoxesGPSCoordinatesSumOfAnEmptyMap(t *testing.T) {
	var emptyMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, ROBOT, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
	}
	assert.Equal(t, 0, emptyMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfAMapWithASingleBoxAtTopLeftEdge(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{BOX, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, ROBOT, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
	}
	assert.Equal(t, ((100 * 1) + 1), warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfAMapWithASingleMiddleBox(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, ROBOT, BOX, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
	}
	assert.Equal(t, ((100 * 3) + 3), warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfAMapWithSomeBoxes(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, BOX},
		{BOX, EMPTY, EMPTY, EMPTY},
		{EMPTY, ROBOT, BOX, BOX},
		{BOX, EMPTY, EMPTY, EMPTY},
	}
	var expected = ((100 * 1) + 4) +
		((100 * 2) + 1) +
		((100 * 3) + 3) + ((100 * 3) + 4) +
		((100 * 4) + 1)
	assert.Equal(t, expected, warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfSmallerProvidedExample(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, BOX, BOX},
		{WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, BOX},
		{EMPTY, WALL, BOX, ROBOT, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
	}
	assert.Equal(t, 2028, warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfLargerProvidedExample(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, BOX, EMPTY, BOX, EMPTY, BOX, BOX, BOX},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{BOX, BOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{BOX, BOX, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{BOX, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, BOX},
		{BOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, BOX, BOX},
		{BOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, BOX, BOX},
		{BOX, BOX, EMPTY, EMPTY, EMPTY, EMPTY, BOX, BOX},
	}
	assert.Equal(t, 10092, warehouseMap.GetBoxesGPSCoordinatesSum())
}
