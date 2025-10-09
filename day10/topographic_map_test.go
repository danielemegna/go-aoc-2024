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
