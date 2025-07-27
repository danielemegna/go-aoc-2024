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
		{position: Position{x: 6, y: 6} /* any velocity */},
	},
}

func TestGetNumerOfRobotsInASpaceWithNoRobots(t *testing.T) {
	var space = Space{
		size:   SpaceSize{width: 7, height: 7},
		guards: []RobotGuard{},
	}

	assert.Equal(t, 0, space.GetNumberOfRobotsInArea(NORTH_WEST))
	assert.Equal(t, 0, space.GetNumberOfRobotsInArea(NORTH_EAST))
	assert.Equal(t, 0, space.GetNumberOfRobotsInArea(SOUTH_WEST))
	assert.Equal(t, 0, space.GetNumberOfRobotsInArea(SOUTH_EAST))
}

func TestGetNumerOfRobotsInNorthEastArea(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(NORTH_EAST)
	assert.Equal(t, 3, robotsInArea)
}

func TestGetNumerOfRobotsInSouthEastArea(t *testing.T) {
	var robotsInArea = PROVIDED_EXAMPLE_SPACE.GetNumberOfRobotsInArea(SOUTH_EAST)
	assert.Equal(t, 1, robotsInArea)
}
