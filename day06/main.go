package day06

import "strings"

func DistinctPositionsVisitedByGuardCount(fileContent string) int {
	var mapRows = linesFrom(fileContent)
	var guardMap = ParseGuardMap(mapRows)

	guardMap = RunGuardWalk(guardMap)

	return len(guardMap.guard.visitedPositions)
}

func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
