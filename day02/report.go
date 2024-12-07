package day02

import (
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func (this Report) IsValid() bool {
	var isAValidIncreasingReport = true
	var isAValidDecreasingReport = true

	for i := 0; i < len(this.levels)-1; i++ {
		var levelDifferenceWithNext = this.levels[i+1] - this.levels[i]

		if isAValidIncreasingReport {
			isAValidIncreasingReport = isAValidIncreasingDiffence(levelDifferenceWithNext)
		}

		if isAValidDecreasingReport {
			isAValidDecreasingReport = isAValidDecreasingDiffence(levelDifferenceWithNext)
		}

		if !isAValidIncreasingReport && !isAValidDecreasingReport {
			return false
		}
	}

	return true
}

func BuildReportFrom(value string) Report {
	var reportLevels []int
	var levelsAsStrings = strings.Split(value, " ")

	for _, levelAsString := range levelsAsStrings {
		var levelValue, _ = strconv.Atoi(levelAsString)
		reportLevels = append(reportLevels, levelValue)
	}

	return Report{reportLevels}
}

func isAValidIncreasingDiffence(levelsDifference int) bool {
	return levelsDifference > 0 && levelsDifference <= 3
}

func isAValidDecreasingDiffence(levelsDifference int) bool {
	return levelsDifference < 0 && levelsDifference >= -3
}
