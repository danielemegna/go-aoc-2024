package day16

import "strings"

func LowestReindeerCostFor(fileContent string) int {
	var mazeMap, reindeer, target = parseMazeMap(fileContent)
	return FindLowestCostToReachTarget(mazeMap, reindeer, target)
}

func parseMazeMap(fileContent string) (MazeMap, Reindeer, Target) {
	var mazeMap = MazeMap{}
	var reindeer Reindeer
	var target Target

	var lines = linesFrom(fileContent)
	for y := 1; y < len(lines)-1; y++ {
		var line = lines[y]
		var mazeMapRow = []MapElement{}
		for x := 1; x < len(line)-1; x++ {
			var char = line[x]
			if char == '#' {
				mazeMapRow = append(mazeMapRow, WALL)
				continue
			}

			var currentCoordinate = Coordinate{x - 1, y - 1}
			if char == 'S' {
				reindeer = Reindeer{Coordinate: currentCoordinate, Direction: RIGHT}
			}
			if char == 'E' {
				target = currentCoordinate
			}
			mazeMapRow = append(mazeMapRow, EMPTY)
		}
		mazeMap = append(mazeMap, mazeMapRow)
	}

	return mazeMap, reindeer, target
}

func linesFrom(s string) []string {
	var rows = strings.Split(s, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
