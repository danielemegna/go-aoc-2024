package day18

func ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent string, memorySpaceSize int, inputBytesToRead int) int {
	var memorySpace = BuildMemorySpaceFrom(fileContent, memorySpaceSize, inputBytesToRead)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	return explorer.ShortestPathFromTopLeftToBottomRight()
}
