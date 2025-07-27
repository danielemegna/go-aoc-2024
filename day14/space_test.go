package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
