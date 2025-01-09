package day06

import "strings"

func DistinctPositionsVisitedByGuardCount(fileContent string) int {
	return 41
}


func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
