package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSingleRobot(t *testing.T) {
	var actual = ParseRobotGuard("p=9,3 v=2,3")

	var expected = RobotGuard{
		position: Position{x: 9, y: 3},
		velocity: Velocity{horizontal: 2, vertical: 3},
	}
	assert.Equal(t, expected, actual)
}

func TestParseMultipleRobots(t *testing.T) {
	var actual = ParseRobotLines(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Len(t, actual, 12)
	assert.Equal(t, Position{0, 4}, actual[0].position)
	assert.Equal(t, Velocity{-1, -3}, actual[1].velocity)
	assert.Equal(t, 10, actual[2].position.x)
	assert.Equal(t, 0, actual[3].position.y)
	assert.Equal(t, 1, actual[4].velocity.horizontal)
	assert.Equal(t, -2, actual[5].velocity.vertical)
	assert.Equal(t, Velocity{-3,-3}, actual[11].velocity)
}
