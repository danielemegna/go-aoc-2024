package day18

func ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent string, memorySpaceSize int) int {
	var memorySpace = BuildMemorySpaceFrom(fileContent, memorySpaceSize)
	var explorer = NewMemorySpaceExplorer(memorySpace)
	return explorer.ShortestPathFromTopLeftToBottomRight()
}
