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

	var length, _ = explorer.ShortestPathFromTopLeftToBottomRight()
	return length
}

func FirstByteMakesBottomRightCornerUnreachable(fileContent string, memorySpaceSize int) Coordinate {
	var memorySpace = NewSafeMemorySpace(memorySpaceSize)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	var lines = linesFrom(fileContent)

	for i := 0; ; i++ {
		var corruptedCoordinate = ParseCoordinateFrom(lines[i])
		memorySpace.Corrupt(corruptedCoordinate)
		// here we could check if the new corrupted coordinate
		// is present in the shortest found path
		// and skip to the next if not
		var shortestPathLength, _ = explorer.ShortestPathFromTopLeftToBottomRight()
		if shortestPathLength == -1 {
			return corruptedCoordinate
		}
	}
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
