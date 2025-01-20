package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
