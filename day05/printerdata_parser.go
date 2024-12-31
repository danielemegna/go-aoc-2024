package day05

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func ParsePrinterData(lines []string) (PageOrderingRules, []PagesToProduceInTheUpdate) {
	var lineIndex = 0

	var rules = PageOrderingRules{}
	for true {
		var line = lines[lineIndex]
		if line == "" {
			lineIndex++
			break
		}

		var pageOrderingRule = parsePageOrderingRule(line)
		rules = append(rules, pageOrderingRule)
		lineIndex++
	}

	var updates = []PagesToProduceInTheUpdate{}
	for lineIndex < len(lines) {
		var line = lines[lineIndex]
		var pagesToProduceInTheUpdate = parsePagesToProduceInTheUpdate(line)
		updates = append(updates, pagesToProduceInTheUpdate)
		lineIndex++
	}

	return rules, updates
}

func parsePageOrderingRule(s string) PageOrderingRule {
	var beforeString, afterString, _ = strings.Cut(s, "|")
	return PageOrderingRule{
		before: toInt(beforeString),
		after:  toInt(afterString),
	}
}

func parsePagesToProduceInTheUpdate(s string) PagesToProduceInTheUpdate {
	var numberStrings = strings.Split(s, ",")
	return lo.Map(numberStrings, func(s string, _ int) int {
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
