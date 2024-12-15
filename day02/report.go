package day02

import (
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func (this Report) IsSafe() bool {
	var isASafeIncreasingReport = true
	var isASafeDecreasingReport = true

	for i := 0; i < len(this.levels)-1; i++ {
		var levelDifferenceWithNext = this.levels[i+1] - this.levels[i]

		if isASafeIncreasingReport {
			isASafeIncreasingReport = isASafeIncreasingDiffence(levelDifferenceWithNext)
		}

		if isASafeDecreasingReport {
			isASafeDecreasingReport = isASafeDecreasingDiffence(levelDifferenceWithNext)
		}

		if !isASafeIncreasingReport && !isASafeDecreasingReport {
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

func isASafeIncreasingDiffence(levelsDifference int) bool {
	return levelsDifference > 0 && levelsDifference <= 3
}

func isASafeDecreasingDiffence(levelsDifference int) bool {
	return levelsDifference < 0 && levelsDifference >= -3
}
