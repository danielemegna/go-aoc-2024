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

	var actual = memorySpace.ShortestPathLengthFromTopLeftToBottomRightCorners()

	assert.Equal(t, 4, actual)
}


func TestProvidedExampleMemorySpace(t *testing.T) {
	t.Skip("WIP")
	var memorySpace = MemorySpace{
		{SAFE, SAFE, SAFE, CORRUPTED, SAFE, SAFE, SAFE},
		{SAFE, SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED, SAFE},
		{SAFE, SAFE, SAFE, SAFE, CORRUPTED, SAFE, SAFE},
		{SAFE, SAFE, SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED},
		{SAFE, SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED, SAFE},
		{SAFE, CORRUPTED, SAFE, SAFE, CORRUPTED, SAFE, SAFE},
		{CORRUPTED, SAFE, CORRUPTED, SAFE, SAFE, SAFE, SAFE},
	}

	var actual = memorySpace.ShortestPathLengthFromTopLeftToBottomRightCorners()

	assert.Equal(t, 22, actual)
}
