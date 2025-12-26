package day21

import (
	"strconv"
	"strings"
)

func SumOfCodesComplexity(fileContent string) int {
	var lines = linesFrom(fileContent)
	var stackOfKeypads = StackOfKeypads{totalNumberOfKeypads: 4}

	var result = 0
	for _, line := range lines {
		var codeToCompose = ParseNumericCodeSequence(line)
		var numericCode, _ = strconv.Atoi(line[:len(line)-1])
		var sequenceLength = stackOfKeypads.LengthOfShortestSequenceToPressOnFirstKeypadFor(codeToCompose)
		result += numericCode * sequenceLength
	}

	return result
}

func linesFrom(s string) []string {
	var trimmed = strings.TrimSpace(s)
	return strings.Split(trimmed, "\n")
}
