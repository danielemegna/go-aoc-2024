package day19

import "strings"

func CountPossibleDesigns(fileContent string) int {
	var fileContentLines = linesFrom(fileContent)
	var availableTowelPatterns = InitAvailableTowelPatternsFrom(fileContentLines[0])

	var possibleDesigns = 0
	for _, towelStringPattern := range fileContentLines[2:] {
		var towelPattern = TowelPattern(towelStringPattern)
		if towelPattern.IsDesignPossibleWith(availableTowelPatterns) {
			possibleDesigns++
		}
	}

	return possibleDesigns
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
