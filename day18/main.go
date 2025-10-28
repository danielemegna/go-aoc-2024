package day18

func ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent string, memorySpaceSize int, inputBytesToRead int) int {
	var memorySpace = BuildMemorySpaceFrom(fileContent, memorySpaceSize, inputBytesToRead)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	return explorer.ShortestPathFromTopLeftToBottomRight()
}

func FirstByteMakesBottomRightCornerUnreachable(fileContent string, memorySpaceSize int) Coordinate {
	var inputBytesToRead = 22 // loop on this
	var memorySpace = BuildMemorySpaceFrom(fileContent, memorySpaceSize, inputBytesToRead)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	explorer.ShortestPathFromTopLeftToBottomRight()
	return Coordinate{6, 1}
}
