package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoTrailheadsInMap(t *testing.T) {
	var topographicMap = TopographicMap{
		{0, 1, 2, 1},
		{1, 1, 3, 1},
		{2, 1, 4, 5},
		{3, 1, 1, 1},
	}

	var trailheads = topographicMap.FindTrailheads()

	assert.Empty(t, trailheads)
}

func TestSingleTrailheadInMap(t *testing.T) {
	var topographicMap = TopographicMap{
		{0, 1, 2, 3},
		{9, 1, 1, 4},
		{1, 1, 1, 5},
		{9, 8, 7, 6},
	}

	var trailheads = topographicMap.FindTrailheads()

	assert.Equal(t, []Trailhead{{
		startingPosition:             Coordinate{0, 0},
		reachableNineHeightPositions: []Coordinate{{0, 3}},
	}}, trailheads)
	assert.Equal(t, 1, trailheads[0].GetScore())
}

func TestSingleTrailheadWithMultipleReachedNineHeightPositions(t *testing.T) {
	var topographicMap = TopographicMap{
		{0, 1, 2, 3},
		{6, 1, 1, 4},
		{1, 9, 1, 5},
		{9, 8, 7, 6},
	}

	var trailheads = topographicMap.FindTrailheads()

	assert.Equal(t, []Trailhead{{
		startingPosition:             Coordinate{0, 0},
		reachableNineHeightPositions: []Coordinate{{0, 3}, {1, 2}},
	}}, trailheads)
	assert.Equal(t, 2, trailheads[0].GetScore())
}

func TestSingleTrailheadStartingBottomRight(t *testing.T) {
	var topographicMap = TopographicMap{
		{1, 1, 4, 3},
		{1, 6, 5, 2},
		{1, 7, 8, 1},
		{1, 1, 9, 0},
	}

	var trailheads = topographicMap.FindTrailheads()

	assert.Equal(t, []Trailhead{{
		startingPosition:             Coordinate{3, 3},
		reachableNineHeightPositions: []Coordinate{{2, 3}},
	}}, trailheads)
	assert.Equal(t, 1, trailheads[0].GetScore())
}

func TestTwoTrailheadInMap(t *testing.T) {
	var topographicMap = TopographicMap{
		{1, 0, 9, 1, 9, 1, 1},
		{2, 9, 1, 1, 8, 1, 1},
		{3, 1, 1, 1, 7, 1, 1},
		{4, 5, 6, 7, 6, 5, 4},
		{1, 1, 1, 8, 1, 1, 3},
		{1, 1, 1, 9, 1, 9, 2},
		{1, 1, 1, 1, 9, 0, 1},
	}

	var trailheads = topographicMap.FindTrailheads()

	assert.Len(t, trailheads, 2)
	var firstTrailhead = trailheads[0]
	assert.Equal(t, Coordinate{1, 0}, firstTrailhead.startingPosition)
	assert.ElementsMatch(t, []Coordinate{{3, 5}}, firstTrailhead.reachableNineHeightPositions)
	assert.Equal(t, 1, firstTrailhead.GetScore())
	var secondTrailhead = trailheads[1]
	assert.Equal(t, Coordinate{5, 6}, secondTrailhead.startingPosition)
	assert.ElementsMatch(t, []Coordinate{{3, 5}, {4, 0}}, secondTrailhead.reachableNineHeightPositions)
	assert.Equal(t, 2, secondTrailhead.GetScore())
}

func TestSingleTrailheadWithFourReachedNineHeightPositions(t *testing.T) {
	t.Skip("WIP remove duplicate in reached nine height positions")
	var topographicMap = TopographicMap{
		{1, 1, 9, 0, 9, 1, 9},
		{1, 1, 1, 1, 1, 9, 8},
		{1, 1, 1, 2, 1, 1, 7},
		{6, 5, 4, 3, 4, 5, 6},
		{7, 6, 5, 1, 9, 8, 7},
		{8, 7, 6, 1, 1, 1, 1},
		{9, 8, 7, 1, 1, 1, 1},
	}

	var trailheads = topographicMap.FindTrailheads()

	assert.Len(t, trailheads, 1)
	assert.Equal(t, 4, trailheads[0].GetScore())
}
