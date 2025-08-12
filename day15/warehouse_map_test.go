package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRobotToTheRightInEmptyMap(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{1, 2}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 2}))
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 2}))
}

func TestMoveRobotUpInEmptyMap(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 1}))
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 1}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 2}))
}

func TestRobotCannotMoveOutOfMapBounds(t *testing.T) {
	var warehouseMap = smallEmptyMap()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))

	warehouseMap.MoveRobot(LEFT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))
}

func TestRobotCannotMoveOnAWall(t *testing.T) {
	var warehouseMap = smallFullMap()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))
}

func TestRobotMovesCloseBoxWithHim(t *testing.T) {
	var warehouseMap = smallFullMap()
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{1, 2}))

	warehouseMap.MoveRobot(DOWN)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 2}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{1, 3}))
}

func TestRobotCannotMoveCloseBoxesBlockedByMapBounds(t *testing.T) {
	var warehouseMap = smallFullMap()
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{0, 1}))

	warehouseMap.MoveRobot(LEFT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{0, 1}))
}

func TestRobotShiftCloseBoxesWhenMoves(t *testing.T) {
	var warehouseMap = smallFullMap()
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))
}

func TestRobotCannotMoveCloseBoxesBlockedByWalls(t *testing.T) {
	var warehouseMap = smallFullMap()
	warehouseMap.MoveRobot(DOWN)
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 2}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{1, 3}))
	assert.Equal(t, WALL, warehouseMap.ElementAt(Coordinate{1, 4}))

	warehouseMap.MoveRobot(DOWN)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 2}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{1, 3}))
	assert.Equal(t, WALL, warehouseMap.ElementAt(Coordinate{1, 4}))
}

func TestRobotCannotMoveMultipleCloseBoxesBlockedByMapBounds(t *testing.T) {
	var warehouseMap = smallFullMap()
	warehouseMap.MoveRobot(RIGHT)
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))
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

func TestGetBoxesGPSCoordinatesSumOfSmallerProvidedExampleAfterAllRobotMoves(t *testing.T) {
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
	assert.Equal(t, Coordinate{0, 2}, smallEmptyMap().GetRobotPosition())
	assert.Equal(t, Coordinate{1, 1}, smallFullMap().GetRobotPosition())
	assert.Equal(t, Coordinate{1, 1}, smallerProvidedExampleMap().GetRobotPosition())
}

func TestGetMapSize(t *testing.T) {
	assert.Equal(t, 4, smallEmptyMap().MapWidth())
	assert.Equal(t, 4, smallEmptyMap().MapHeigth())
	assert.Equal(t, 5, smallFullMap().MapWidth())
	assert.Equal(t, 5, smallFullMap().MapHeigth())
	assert.Equal(t, 6, smallerProvidedExampleMap().MapWidth())
	assert.Equal(t, 6, smallerProvidedExampleMap().MapHeigth())
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

func smallerProvidedExampleMap() WarehouseMap {
	return WarehouseMap{
		{EMPTY, EMPTY, BOX, EMPTY, BOX, EMPTY},
		{WALL, ROBOT, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, WALL, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}
}
