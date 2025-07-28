package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var SPACE_SIZE = SpaceSize{width: 11, height: 7}

func TestChangePositionAfterOneSecondTick(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 2, y: 2},
		velocity: Velocity{horizontal: 1, vertical: 1},
	}

	robot.AfterOneSecond(SPACE_SIZE)

	var expectedNewPosition = Position{x: 2 + 1, y: 2 + 1}
	assert.Equal(t, expectedNewPosition, robot.position)
}

func TestTeleportAtTheEdgeOfTheSpaceVertically(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 5, y: 6},
		velocity: Velocity{horizontal: 1, vertical: 1},
	}

	robot.AfterOneSecond(SPACE_SIZE)

	var expectedNewPosition = Position{x: 5 + 1, y: 0}
	assert.Equal(t, expectedNewPosition, robot.position)
}

func TestTeleportAtTheEdgeOfTheSpaceHorizontally(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 9, y: 2},
		velocity: Velocity{horizontal: 3, vertical: 3},
	}

	robot.AfterOneSecond(SPACE_SIZE)

	var expectedNewPosition = Position{x: 1, y: 2 + 3}
	assert.Equal(t, expectedNewPosition, robot.position)
}

func TestTeleportAtTheEdgeOfTheSpaceBackward(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 1, y: 1},
		velocity: Velocity{horizontal: -3, vertical: -3},
	}

	robot.AfterOneSecond(SPACE_SIZE)

	var expectedNewPosition = Position{x: 9, y: 5}
	assert.Equal(t, expectedNewPosition, robot.position)
}

func TestProvidedExampleRobotChangePosition(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 2, y: 4},
		velocity: Velocity{horizontal: 2, vertical: -3},
	}

	robot.AfterOneSecond(SPACE_SIZE)

	var expectedNewPosition = Position{x: 4, y: 1}
	assert.Equal(t, expectedNewPosition, robot.position)
}

func TestProvidedExampleRobotChangePositionAfter5Seconds(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 2, y: 4},
		velocity: Velocity{horizontal: 2, vertical: -3},
	}

	robot.AfterSeconds(5, SPACE_SIZE)

	var expectedNewPosition = Position{x: 1, y: 3}
	assert.Equal(t, expectedNewPosition, robot.position)
}

func TestRobotWithHypervelocityChangePosition(t *testing.T) {
	var robot = RobotGuard{
		position: Position{x: 4, y: 4},
		velocity: Velocity{horizontal: 11 * 10, vertical: 7 * 10},
	}

	robot.AfterOneSecond(SPACE_SIZE)

	var expectedNewPosition = Position{x: 4, y: 4}
	assert.Equal(t, expectedNewPosition, robot.position)
}
