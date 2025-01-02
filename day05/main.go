package day05

import "strings"

func SumOfCorrectlyOrderedUpdatesMiddlePageNumbers(fileContent string) int {
	var lines = linesFrom(fileContent)
	var rules, updates = ParsePrinterData(lines)

	var validUpdatesMiddlePageNumbersSum = 0
	for _, update := range updates {
		if(update.IsValidFor(rules)) {
				validUpdatesMiddlePageNumbersSum += update.GetMiddlePageNumber()
		}
		
	}

	return validUpdatesMiddlePageNumbersSum
}

func SumOfIncorrectlyOrderedFixedUpdatesMiddlePageNumbers(fileContent string) int {
	var lines = linesFrom(fileContent)
	var _, _ = ParsePrinterData(lines)

	return 123
}


func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
