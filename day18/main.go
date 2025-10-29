package day18

import (
	"slices"
	"strings"
)

func ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent string, memorySpaceSize int, inputBytesToRead int) int {
	var memorySpace = NewSafeMemorySpace(memorySpaceSize)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	var lines = linesFrom(fileContent)

	for i := 0; i < inputBytesToRead; i++ {
		var corruptedCoordinate = ParseCoordinateFrom(lines[i])
		memorySpace.Corrupt(corruptedCoordinate)
	}

	var pathCoordinates = explorer.ShortestPathFromTopLeftToBottomRight()
	return len(pathCoordinates) - 1
}

func FirstByteMakesBottomRightCornerUnreachable(fileContent string, memorySpaceSize int) Coordinate {
	var memorySpace = NewSafeMemorySpace(memorySpaceSize)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	var lines = linesFrom(fileContent)

	var latestFoundPathCoordinates = explorer.ShortestPathFromTopLeftToBottomRight()
	for i := 0; ; i++ {
		var corruptedCoordinate = ParseCoordinateFrom(lines[i])
		memorySpace.Corrupt(corruptedCoordinate)
		if !slices.Contains(latestFoundPathCoordinates, corruptedCoordinate) {
			continue
		}
		latestFoundPathCoordinates = explorer.ShortestPathFromTopLeftToBottomRight()
		if len(latestFoundPathCoordinates) == 0 {
			return corruptedCoordinate
		}
	}
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
