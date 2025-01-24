package day06

import (
	"github.com/samber/lo"
	"strings"
)

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
	var possiblePlacesToPutExtraObstacle = lo.Without(normalWalkMap.guard.visitedPositions, guardMap.guard.position)
	for _, visitedPosition := range possiblePlacesToPutExtraObstacle {
		var guardWalkWithExtraObstacle = RunGuardWalkWithExtraObstacle(guardMap, visitedPosition)
		if guardWalkWithExtraObstacle.IsGuardOutOfBoundaries() {
			continue
		}
		possibleWaysToCreateLoopWithAnExtraObstacle++
	}

	return possibleWaysToCreateLoopWithAnExtraObstacle
}

func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
