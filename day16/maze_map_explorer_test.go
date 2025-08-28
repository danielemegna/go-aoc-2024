package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCostOneStepForehand(t *testing.T) {
	var reindeer = Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT}
	var target = Coordinate{1, 0}
	var maze = MazeMap{
		{EMPTY, EMPTY},
		{EMPTY, EMPTY},
	}

	var cost = FindLowestCostToReachTarget(maze, reindeer, target)

	assert.Equal(t, 1, cost)
}

func TestCostTwoStepForehand(t *testing.T) {
	var reindeer = Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT}
	var target = Coordinate{2, 0}
	var maze = MazeMap{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}

	var cost = FindLowestCostToReachTarget(maze, reindeer, target)

	assert.Equal(t, 2, cost)
}

func TestCostTwoStepAndATurn(t *testing.T) {
	var reindeer = Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT}
	var target = Coordinate{2, 1}
	var maze = MazeMap{
		{EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY},
	}

	var cost = FindLowestCostToReachTarget(maze, reindeer, target)

	assert.Equal(t, 1003, cost)
}

func TestCostWithEmptySmallMap(t *testing.T) {
	var reindeer = Reindeer{Coordinate: Coordinate{0, 3}, Direction: RIGHT}
	var target = Coordinate{3, 0}
	var maze = MazeMap{
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY},
	}

	var cost = FindLowestCostToReachTarget(maze, reindeer, target)

	assert.Equal(t, 1006, cost)
}
