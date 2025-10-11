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
