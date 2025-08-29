package day16

func LowestReindeerCostFor(fileContent string) int {
	var mazeMap, reindeer, target = ParseMazeMap(fileContent)
	return FindLowestCostToReachTarget(mazeMap, reindeer, target)
}
