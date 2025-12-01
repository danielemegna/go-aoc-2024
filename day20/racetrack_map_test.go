package day20

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	var _ = RacetrackMap{
		{START, 1, 	  2,    3},
		{WALL,  WALL, 5,    4},
		{WALL,  7,    6,    WALL},
		{WALL,  8,    WALL, WALL},
	}

	assert.True(t, true)
}
