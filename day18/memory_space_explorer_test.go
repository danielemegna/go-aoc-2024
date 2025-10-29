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

	var actual = explorer.ShortestPathFromTopLeftToBottomRight()

	assert.Len(t, actual, 4 + 1)
	assert.ElementsMatch(t, []Coordinate{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, actual)
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

	var actual = explorer.ShortestPathFromTopLeftToBottomRight()

	assert.Len(t, actual, 10 + 1)
	assert.ElementsMatch(t, []Coordinate{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 4}, {2, 4}, {2, 3}, {3, 3}, {4, 3}, {4, 4},
	}, actual)
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

	var actual = explorer.ShortestPathFromTopLeftToBottomRight()

	assert.Len(t, actual, 22 + 1)
}
