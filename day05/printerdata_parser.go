package day05

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func ParsePrinterData(inputLines []string) ([]PageOrderingRule, []PagesToProduceInTheUpdate) {
	var lineIndex = 0

	var pageOrderingRule = []PageOrderingRule{
		{before: 47, after: 53},
		{before: 53, after: 13},
	}
	for true {
		var line = inputLines[lineIndex]
		if line == "" {
			lineIndex++
			break
		}

		// TODO parse the page ordering rule here
		lineIndex++
	}

	var pagesToProduceInTheUpdate = []PagesToProduceInTheUpdate{}
	for lineIndex < len(inputLines) {
		var numberStrings = strings.Split(inputLines[lineIndex], ",")
		var numbers = toIntSlice(numberStrings)
		pagesToProduceInTheUpdate = append(pagesToProduceInTheUpdate, numbers)
		lineIndex++
	}

	return pageOrderingRule, pagesToProduceInTheUpdate
}

func toIntSlice(slice []string) []int {
	return lo.Map(slice, func(s string, _ int) int {
		var value, _ = strconv.Atoi(s)
		return value
	})
}
