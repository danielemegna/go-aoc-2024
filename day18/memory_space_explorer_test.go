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

	assert.Equal(t, 4, actual)
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

	assert.Equal(t, 10, actual)
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

	assert.Equal(t, 22, actual)
}
