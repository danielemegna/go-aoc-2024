package day18

import "strings"

func ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent string, memorySpaceSize int, inputBytesToRead int) int {
	var memorySpace = NewSafeMemorySpace(memorySpaceSize)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	var lines = linesFrom(fileContent)

	for i := 0; i < inputBytesToRead; i++ {
		var corruptedCoordinate = ParseCoordinateFrom(lines[i])
		memorySpace.Corrupt(corruptedCoordinate)
	}

	return explorer.ShortestPathFromTopLeftToBottomRight()
}

func FirstByteMakesBottomRightCornerUnreachable(fileContent string, memorySpaceSize int) Coordinate {
	var memorySpace = NewSafeMemorySpace(memorySpaceSize)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	var lines = linesFrom(fileContent)

	for i := 0; ; i++ {
		var corruptedCoordinate = ParseCoordinateFrom(lines[i])
		memorySpace.Corrupt(corruptedCoordinate)
		var shortestPathLength = explorer.ShortestPathFromTopLeftToBottomRight()
		if shortestPathLength == -1 {
			return corruptedCoordinate
		}
	}
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
