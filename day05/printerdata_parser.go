package day05

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func ParsePrinterData(inputLines []string) ([]PageOrderingRule, []PagesToProduceInTheUpdate) {
	var lineIndex = 0

	var pageOrderingRule = []PageOrderingRule{}
	for true {
		var line = inputLines[lineIndex]
		if line == "" {
			lineIndex++
			break
		}

		var beforeString, afterString, _ = strings.Cut(line, "|")
		pageOrderingRule = append(pageOrderingRule, PageOrderingRule{
			before: toInt(beforeString),
			after:  toInt(afterString),
		})
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
		return toInt(s)
	})
}

func toInt(s string) int {
	var value, err = strconv.Atoi(s)
	if err == nil {
		return value
	}

	panic(err)
}
