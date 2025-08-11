package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRobotToTheRightInEmptyMap(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, ROBOT, warehouseMap[2][0])
	assert.Equal(t, EMPTY, warehouseMap[2][1])

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, EMPTY, warehouseMap[2][0])
	assert.Equal(t, ROBOT, warehouseMap[2][1])
}

func TestMoveRobotUpInEmptyMap(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, EMPTY, warehouseMap[1][0])
	assert.Equal(t, ROBOT, warehouseMap[2][0])

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap[1][0])
	assert.Equal(t, EMPTY, warehouseMap[2][0])
}

func TestRobotCannotMoveOutOfMapBounds(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, ROBOT, warehouseMap[2][0])

	warehouseMap.MoveRobot(LEFT)

	assert.Equal(t, ROBOT, warehouseMap[2][0])
}

func TestRobotCannotMoveOnAWall(t *testing.T) {
	var warehouseMap = smallFullMap()
	assert.Equal(t, ROBOT, warehouseMap[1][1])

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap[1][1])
}

func TestRobotMovesCloseBoxWithHim(t *testing.T) {
	var warehouseMap = smallFullMap()
	assert.Equal(t, ROBOT, warehouseMap[1][1])
	assert.Equal(t, BOX, warehouseMap[2][1])

	warehouseMap.MoveRobot(DOWN)

	assert.Equal(t, ROBOT, warehouseMap[2][1])
	assert.Equal(t, BOX, warehouseMap[3][1])
}

func TestGetBoxesGPSCoordinatesSumOfAnEmptyMap(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, 0, warehouseMap.GetBoxesGPSCoordinatesSum())
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
	var warehouseMap = smallFullMap()
	var expected = ((100 * 2) + 1) +
		((100 * 2) + 3) + ((100 * 2) + 4) +
		((100 * 3) + 2)
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

func TestGetRobotPosition(t *testing.T) {
	var x, y = smallEmptyMap().GetRobotPosition()
	assert.Equal(t, 0, x)
	assert.Equal(t, 2, y)
	x, y = smallFullMap().GetRobotPosition()
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, y)
}

func smallEmptyMap() WarehouseMap {
	return WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{ROBOT, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
	}
}

func smallFullMap() WarehouseMap {
	return WarehouseMap{
		{EMPTY, WALL, EMPTY, EMPTY, EMPTY},
		{BOX, ROBOT, BOX, BOX, EMPTY},
		{EMPTY, BOX, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, WALL, EMPTY, EMPTY, EMPTY},
	}
}
