package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var PROVIDED_EXAMPLE_SPACE = Space{
	size: SpaceSize{width: 11, height: 7},
	guards: []RobotGuard{
		{position: Position{x: 6, y: 0}, velocity: Velocity{1, 2}},
		{position: Position{x: 6, y: 0}, velocity: Velocity{3, -1}},
		{position: Position{x: 9, y: 0} /* any velocity */},
		{position: Position{x: 0, y: 2} /* any velocity */},
		{position: Position{x: 1, y: 3} /* any velocity */},
		{position: Position{x: 2, y: 3} /* any velocity */},
		{position: Position{x: 5, y: 4} /* any velocity */},
		{position: Position{x: 3, y: 5} /* any velocity */},
		{position: Position{x: 4, y: 5}, velocity: Velocity{1, 1}},
		{position: Position{x: 4, y: 5}, velocity: Velocity{-2, 3}},
		{position: Position{x: 1, y: 6} /* any velocity */},
		{position: Position{x: 6, y: 6} /* any velocity */},
	},
}

func TestGetNumerOfRobotsInASpaceWithNoRobots(t *testing.T) {
	var emptySpace = Space{
		size:   SpaceSize{width: 7, height: 7},
		guards: []RobotGuard{},
	}

	assert.Equal(t, 0, emptySpace.GetNumberOfRobotsInArea(NORTH_WEST))
	assert.Equal(t, 0, emptySpace.GetNumberOfRobotsInArea(NORTH_EAST))
	assert.Equal(t, 0, emptySpace.GetNumberOfRobotsInArea(SOUTH_WEST))
	assert.Equal(t, 0, emptySpace.GetNumberOfRobotsInArea(SOUTH_EAST))
	assert.Equal(t, 0, emptySpace.GetSafetyFactor())
}

func TestGetNumerOfRobotsInNorthEastArea(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(NORTH_EAST)
	assert.Equal(t, 3, robotsInArea)
}

func TestGetNumerOfRobotsInSouthEastArea(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(SOUTH_EAST)
	assert.Equal(t, 1, robotsInArea)
}

func TestGetNumerOfRobotsInNorthWestArea(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(NORTH_WEST)
	assert.Equal(t, 1, robotsInArea)
}

func TestGetNumerOfRobotsInSouthWestArea(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(SOUTH_WEST)
	assert.Equal(t, 4, robotsInArea)
}

func TestGetSafetyFactor(t *testing.T) {
	var safetyFactor = PROVIDED_EXAMPLE_SPACE.GetSafetyFactor()
	assert.Equal(t, 3*1*1*4, safetyFactor)
}

func TestCoordinatesWithoutGuard(t *testing.T) {
	assert.False(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(0,0))
	assert.False(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(0,1))
	assert.False(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(1,0))
	assert.False(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(9,9))
}

func TestCoordinatesWithGuard(t *testing.T) {
	assert.True(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(6,0))
	assert.True(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(9,0))
	assert.True(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(3,5))
	assert.True(t, PROVIDED_EXAMPLE_SPACE.HasGuardIn(6,6))
}

func TestGetNumerOfRobotsInNorthEastAreaWithGreaterCenterThickness(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInAreaWith(NORTH_EAST, 3)
	assert.Equal(t, 1, robotsInArea)
}

func TestGetNumerOfRobotsInSouthWestAreaWithGreaterCenterThickness(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInAreaWith(SOUTH_WEST, 3)
	assert.Equal(t, 1, robotsInArea)
}

func TestGetSafetyFactorWithGreaterCenterThickness(t *testing.T) {
	var safetyFactor = PROVIDED_EXAMPLE_SPACE.GetSafetyFactorWith(1)
	assert.Equal(t, 12, safetyFactor)
	safetyFactor = PROVIDED_EXAMPLE_SPACE.GetSafetyFactorWith(2)
	assert.Equal(t, 0, safetyFactor)
}

func TestGetNumberOfRobotsInCenterArea(t *testing.T) {
	assert.Equal(t, 1, PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(CENTER))
	assert.Equal(t, 4, PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInAreaWith(CENTER, 2))
	assert.Equal(t, 8, PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInAreaWith(CENTER, 3))
}
