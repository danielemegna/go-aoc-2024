package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSafeSmallMapShortestPath(t *testing.T) {
	var memorySpace = MemorySpace{
		{SAFE, SAFE, SAFE},
		{SAFE, SAFE, SAFE},
		{SAFE, SAFE, SAFE},
	}
	var explorer = NewMemorySpaceExplorer(memorySpace)

	var length, pathCoordinates = explorer.ShortestPathFromTopLeftToBottomRight()

	assert.Equal(t, 4, length)
	assert.ElementsMatch(t, []Coordinate{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, pathCoordinates)
}

func TestSmallMapWithSomeCorrupted(t *testing.T) {
	var memorySpace = MemorySpace{
		{SAFE, CORRUPTED, SAFE, SAFE, SAFE},
		{SAFE, CORRUPTED, SAFE, SAFE, SAFE},
		{SAFE, CORRUPTED, SAFE, SAFE, SAFE},
		{SAFE, CORRUPTED, SAFE, SAFE, SAFE},
		{SAFE, SAFE, SAFE, CORRUPTED, SAFE},
	}
	var explorer = NewMemorySpaceExplorer(memorySpace)

	var length, pathCoordinates = explorer.ShortestPathFromTopLeftToBottomRight()

	assert.Equal(t, 10, length)
	assert.ElementsMatch(t, []Coordinate{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 4}, {2, 4}, {2, 3}, {3, 3}, {4, 3}, {4, 4},
	}, pathCoordinates)
}

func TestProvidedExampleMemorySpace(t *testing.T) {
	var memorySpace = MemorySpace{
		{SAFE, SAFE, SAFE, CORRUPTED, SAFE, SAFE, SAFE},
		{SAFE, SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED, SAFE},
		{SAFE, SAFE, SAFE, SAFE, CORRUPTED, SAFE, SAFE},
		{SAFE, SAFE, SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED},
		{SAFE, SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED, SAFE},
		{SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED, SAFE, SAFE},
		{CORRUPTED, SAFE, CORRUPTED, SAFE, SAFE, SAFE, SAFE},
	}
	var explorer = NewMemorySpaceExplorer(memorySpace)

	var length, pathCoordinates = explorer.ShortestPathFromTopLeftToBottomRight()

	assert.Equal(t, 22, length)
	assert.Len(t, pathCoordinates, 23)
}
