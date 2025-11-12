package day19

import "strings"

func CountPossibleDesigns(fileContent string) int {
	var fileContentLines = linesFrom(fileContent)
	var availableTowelPatterns = AvailableTowelPatternsFrom(fileContentLines[0])

	var possibleDesigns = 0
	for _, towelStringPattern := range fileContentLines[2:] {
		var towelPattern = TowelPattern(towelStringPattern)
		if availableTowelPatterns.IsDesignPossible(towelPattern) {
			possibleDesigns++
		}
	}

	return possibleDesigns
}

func SumOfPossibleDesignsCombinations(fileContent string) int {
	var fileContentLines = linesFrom(fileContent)
	var availableTowelPatterns = AvailableTowelPatternsFrom(fileContentLines[0])

	var possibleDesigns = 0
	for _, towelStringPattern := range fileContentLines[2:] {
		var towelPattern = TowelPattern(towelStringPattern)
		possibleDesigns += availableTowelPatterns.PossibleDesignCombinationsFor(towelPattern)
	}

	return possibleDesigns
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
