package day06

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunSimpleGuardWalkWithoutLoop(t *testing.T) {
	var guardMap = ParseGuardMap([]string{
		"#...",
		"...#",
		"^...",
		"....",
	})

	var newGuardMap = RunGuardWalk(guardMap)

	assert.Len(t, newGuardMap.guard.visitedPositions, 6)
	assert.True(t, newGuardMap.IsGuardOutOfBoundaries())
}

func TestRunGuardWalkWithLoop(t *testing.T) {
	var guardMap = ParseGuardMap([]string{
		".#...",
		"....#",
		".....",
		"#^...",
		"...#.",
	})

	var newGuardMap = RunGuardWalk(guardMap)

	assert.Len(t, newGuardMap.guard.visitedPositions, 8)
	assert.False(t, newGuardMap.IsGuardOutOfBoundaries())
}

func TestRunGuardWalkOfProvidedExample(t *testing.T) {
	var guardMap = ParseGuardMap(PROVIDED_EXAMPLE_INPUT_LINES)

	var newGuardMap = RunGuardWalk(guardMap)

	assert.Len(t, newGuardMap.guard.visitedPositions, 41)
	assert.True(t, newGuardMap.IsGuardOutOfBoundaries())
}

func TestRunGuardWalkOfProvidedExampleAddingAnExtraObstacleCreatingALoop(t *testing.T) {
	var testCasesCreatingALoop = []struct {
		obstacleCoordinate       Coordinate
		expectedVisitedPositions int
	}{
		{obstacleCoordinate: Coordinate{x: 3, y: 6}, expectedVisitedPositions: 18},
		{obstacleCoordinate: Coordinate{x: 7, y: 7}, expectedVisitedPositions: 38},
		{obstacleCoordinate: Coordinate{x: 1, y: 8}, expectedVisitedPositions: 33},
	}

	var guardMap = ParseGuardMap(PROVIDED_EXAMPLE_INPUT_LINES)
	for index, testCase := range testCasesCreatingALoop {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {

			var newGuardMap = RunGuardWalkWithExtraObstacle(guardMap, testCase.obstacleCoordinate)

			assert.Len(t, newGuardMap.guard.visitedPositions, testCase.expectedVisitedPositions)
			assert.False(t, newGuardMap.IsGuardOutOfBoundaries())
		})
	}
}
