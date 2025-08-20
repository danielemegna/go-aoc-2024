package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRobotToTheRightInEmptyMap(t *testing.T) {
	var warehouseMap = aSmallEmptyMap()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{1, 2}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 2}))
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 2}))
}

func TestMoveRobotUpInEmptyMap(t *testing.T) {
	var warehouseMap = aSmallEmptyMap()
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 1}))
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 1}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{0, 2}))
}

func TestRobotCannotMoveOutOfMapBounds(t *testing.T) {
	var warehouseMap = aSmallEmptyMap()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))

	warehouseMap.MoveRobot(LEFT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{0, 2}))
}

func TestRobotCannotMoveOnAWall(t *testing.T) {
	var warehouseMap = aSmallFullMap()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))
}

func TestRobotMovesCloseBoxWithHim(t *testing.T) {
	var warehouseMap = aSmallFullMap()
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{1, 2}))

	warehouseMap.MoveRobot(DOWN)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 2}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{1, 3}))
}

func TestRobotMovesToTheRightCloseBigBoxWithHim(t *testing.T) {
	t.Skip("WIP")
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, ROBOT, LBOX, RBOX, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, LBOX, warehouseMap.ElementAt(Coordinate{4, 1}))
	assert.Equal(t, RBOX, warehouseMap.ElementAt(Coordinate{5, 1}))
}

func TestRobotMovesToTheLeftCloseBigBoxWithHim(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, LBOX, RBOX, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}

	warehouseMap.MoveRobot(LEFT)

	assert.Equal(t, LBOX, warehouseMap.ElementAt(Coordinate{0, 1}))
	assert.Equal(t, RBOX, warehouseMap.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
}

func TestRobotMovesUpCloseBigBoxWithHim(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, LBOX, warehouseMap.ElementAt(Coordinate{2, 0}))
	assert.Equal(t, RBOX, warehouseMap.ElementAt(Coordinate{3, 0}))
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
}

func TestRobotMovesDownCloseBigBoxWithHim(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}

	warehouseMap.MoveRobot(DOWN)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, LBOX, warehouseMap.ElementAt(Coordinate{1, 2}))
	assert.Equal(t, RBOX, warehouseMap.ElementAt(Coordinate{2, 2}))
}

func TestRobotCannotMoveCloseBoxesBlockedByMapBounds(t *testing.T) {
	var warehouseMap = aSmallFullMap()
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{0, 1}))

	warehouseMap.MoveRobot(LEFT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{0, 1}))
}

func TestRobotShiftCloseBoxesWhenMoves(t *testing.T) {
	var warehouseMap = aSmallFullMap()
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))
}

func TestRobotCannotMoveCloseBoxesBlockedByWalls(t *testing.T) {
	var warehouseMap = aSmallFullMap()
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
	var warehouseMap = aSmallFullMap()
	warehouseMap.MoveRobot(RIGHT)
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{3, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))
}

func TestRobotCannotMoveCloseBigBoxesPartiallyBlockedByWalls(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 3}))
	assert.Equal(t, LBOX, warehouseMap.ElementAt(Coordinate{1, 2}))
	assert.Equal(t, RBOX, warehouseMap.ElementAt(Coordinate{2, 2}))
}

func TestRobotCannotMoveMultipleCloseBigBoxesPartiallyBlockedByWalls(t *testing.T) {
	var warehouseMap = WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, WALL, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
	}

	warehouseMap.MoveRobot(UP)

	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{2, 3}))
	assert.Equal(t, LBOX, warehouseMap.ElementAt(Coordinate{2, 2}))
	assert.Equal(t, RBOX, warehouseMap.ElementAt(Coordinate{3, 2}))
}

func TestMakeAllRobotMovesInSmallerProvidedExampleMap(t *testing.T) {
	var warehouseMap = expectedSmallerProvidedExampleParsedMap()

	for _, direction := range expectedSmallerProvidedExampleParsedMoves() {
		warehouseMap.MoveRobot(direction)
	}

	var expected = expectedSmallerProvidedExampleMapAfterAllRobotMoves()
	assert.Equal(t, expected, warehouseMap)
}

func TestGetBoxesGPSCoordinatesSumOfAnEmptyMap(t *testing.T) {
	var warehouseMap = aSmallEmptyMap()
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
	var warehouseMap = aSmallFullMap()
	var expected = ((100 * 2) + 1) +
		((100 * 2) + 3) + ((100 * 2) + 4) +
		((100 * 3) + 2)
	assert.Equal(t, expected, warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfSmallerProvidedExampleAfterAllRobotMoves(t *testing.T) {
	var warehouseMap = expectedSmallerProvidedExampleMapAfterAllRobotMoves()
	assert.Equal(t, 2028, warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetBoxesGPSCoordinatesSumOfLargerProvidedExampleAfterAllRobotMoves(t *testing.T) {
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

func TestGetBoxesGPSCoordinatesSumOfLargerProvidedExampleInDoubleWideAfterAllRobotMoves(t *testing.T) {
	var warehouseMap = expectedLargerProvidedExampleMapInDoubleWideAfterAllRobotMoves()
	assert.Equal(t, 9021, warehouseMap.GetBoxesGPSCoordinatesSum())
}

func TestGetRobotPosition(t *testing.T) {
	assert.Equal(t, Coordinate{0, 2}, aSmallEmptyMap().GetRobotPosition())
	assert.Equal(t, Coordinate{1, 1}, aSmallFullMap().GetRobotPosition())
	assert.Equal(t, Coordinate{1, 1}, expectedSmallerProvidedExampleParsedMap().GetRobotPosition())
}

func TestGetMapSize(t *testing.T) {
	assert.Equal(t, 4, aSmallEmptyMap().MapSize())
	assert.Equal(t, 5, aSmallFullMap().MapSize())
	assert.Equal(t, 6, expectedSmallerProvidedExampleParsedMap().MapSize())
	assert.Equal(t, 4, aSmallEmptyMap().GetWidth())
	assert.Equal(t, 4, aSmallEmptyMap().GetHeight())
	assert.Equal(t, 5, aSmallFullMap().GetWidth())
	assert.Equal(t, 5, aSmallFullMap().GetHeight())
	assert.Equal(t, 6, expectedSmallerProvidedExampleParsedMap().GetWidth())
	assert.Equal(t, 6, expectedSmallerProvidedExampleParsedMap().GetHeight())
	assert.Equal(t, 16, expectedLargerProvidedExampleMapInDoubleWideAfterAllRobotMoves().GetWidth())
	assert.Equal(t, 8, expectedLargerProvidedExampleMapInDoubleWideAfterAllRobotMoves().GetHeight())
}

func TestClone(t *testing.T) {
	var warehouseMap = aSmallFullMap()
	var clone = warehouseMap.Clone()
	assert.Equal(t, ROBOT, warehouseMap.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{4, 1}))
	assert.Equal(t, ROBOT, clone.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, EMPTY, clone.ElementAt(Coordinate{4, 1}))

	warehouseMap.MoveRobot(RIGHT)

	assert.Equal(t, EMPTY, warehouseMap.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, BOX, warehouseMap.ElementAt(Coordinate{4, 1}))
	assert.Equal(t, ROBOT, clone.ElementAt(Coordinate{1, 1}))
	assert.Equal(t, EMPTY, clone.ElementAt(Coordinate{4, 1}))
}

func aSmallEmptyMap() WarehouseMap {
	return WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{ROBOT, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
	}
}

func aSmallFullMap() WarehouseMap {
	return WarehouseMap{
		{EMPTY, WALL, EMPTY, EMPTY, EMPTY},
		{BOX, ROBOT, BOX, BOX, EMPTY},
		{EMPTY, BOX, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, WALL, EMPTY, EMPTY, EMPTY},
	}
}

func expectedSmallerProvidedExampleMapAfterAllRobotMoves() WarehouseMap {
	return WarehouseMap{
		{EMPTY, EMPTY, EMPTY, EMPTY, BOX, BOX},
		{WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, BOX},
		{EMPTY, WALL, BOX, ROBOT, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, BOX, EMPTY, EMPTY},
	}
}

func expectedLargerProvidedExampleMapInDoubleWideAfterAllRobotMoves() WarehouseMap {
	/*
		####################
		##[].......[].[][]##
		##[]...........[].##
		##[]........[][][]##
		##[]......[]....[]##
		##..##......[]....##
		##..[]............##
		##..@......[].[][]##
		##......[][]..[]..##
		####################
	*/
	return WarehouseMap{
		{LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, LBOX, RBOX, LBOX, RBOX},
		{LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY},
		{LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, LBOX, RBOX, LBOX, RBOX},
		{LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX},
		{EMPTY, EMPTY, WALL, WALL, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, ROBOT, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, EMPTY, LBOX, RBOX, LBOX, RBOX},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, LBOX, RBOX, LBOX, RBOX, EMPTY, EMPTY, LBOX, RBOX, EMPTY, EMPTY},
	}
}
