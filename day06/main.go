package day06

import "strings"

func DistinctPositionsVisitedByGuardCount(fileContent string) int {
	var mapRows = linesFrom(fileContent)
	var guardMap = ParseGuardMap(mapRows)

	guardMap = RunGuardWalk(guardMap)

	return len(guardMap.guard.visitedPositions)
}

func PossibleWaysToCreateLoopWithAnExtraObstacle(fileContent string) int {
	var mapRows = linesFrom(fileContent)
	var guardMap = ParseGuardMap(mapRows)

	var normalWalkMap = RunGuardWalk(guardMap)
	var possibleWaysToCreateLoopWithAnExtraObstacle = 0
	for _, visitedPosition := range normalWalkMap.guard.visitedPositions {
		var guardWalkWithExtraObstacle = RunGuardWalkWithExtraObstacle(guardMap, visitedPosition)
		if(guardWalkWithExtraObstacle.IsGuardOutOfBoundaries()) {
			continue
		}
		possibleWaysToCreateLoopWithAnExtraObstacle++
	}

	return 6
}


func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
