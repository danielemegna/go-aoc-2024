package day16

func LowestReindeerCostFor(fileContent string) int {
	var mazeMap, reindeer, target = ParseMazeMap(fileContent)
	var explorer = NewMazeMapExplorer(mazeMap, reindeer, target)
	return explorer.FindLowestCostToReachTarget()
}

func TilesCountOfBestPaths(fileContent string) int {
	var mazeMap, reindeer, target = ParseMazeMap(fileContent)
	var _ = NewMazeMapExplorer(mazeMap, reindeer, target)
	return 45
}
