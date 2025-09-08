package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCostOneStepForehand(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY},
			{EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT},
		Coordinate{1, 0},
	)

	var cost = explorer.FindLowestCostToReachTarget()

	assert.Equal(t, 1, cost)
}

func TestCostTwoStepForehand(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT},
		Coordinate{2, 0},
	)

	var cost = explorer.FindLowestCostToReachTarget()

	assert.Equal(t, 2, cost)
}

func TestCostTwoStepAndATurn(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT},
		Coordinate{2, 1},
	)

	var cost = explorer.FindLowestCostToReachTarget()

	assert.Equal(t, 1003, cost)
}

func TestCostTwoStepDown(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY},
			{EMPTY, EMPTY},
			{EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT},
		Coordinate{0, 2},
	)

	var cost = explorer.FindLowestCostToReachTarget()

	assert.Equal(t, 1002, cost)
}

func TestFourStepsAndOneTurn(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT},
		Coordinate{2, 2},
	)

	var cost = explorer.FindLowestCostToReachTarget()

	assert.Equal(t, 1004, cost)
}

func TestCostWithEmptySmallMap(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 3}, Direction: RIGHT},
		Coordinate{3, 0},
	)

	var cost = explorer.FindLowestCostToReachTarget()

	assert.Equal(t, 1006, cost)
}

func TestBestPathsCoordinatesCountTwoStepForehand(t *testing.T) {
	var explorer = NewMazeMapExplorer(
		MazeMap{
			{EMPTY, EMPTY, EMPTY},
			{EMPTY, EMPTY, EMPTY},
		},
		Reindeer{Coordinate: Coordinate{0, 0}, Direction: RIGHT},
		Coordinate{2, 0},
	)

	var count = explorer.CoordinatesCountOfBestPaths()

	assert.Equal(t, 3, count)
}
