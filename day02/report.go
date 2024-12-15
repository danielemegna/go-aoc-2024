package day02

import (
	"slices"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func (this Report) IsSafe() bool {
	return this.unsafeValueIndex() == nil
}

func (this Report) IsSafeWithTolerance() bool {
	var unsafeValueIndex = this.unsafeValueIndex()
	if unsafeValueIndex == nil {
		return true
	}

	var reportsToTry = []Report{
		{slices.Clone(this.levels[1:])},
		{append(
			slices.Clone(this.levels[:*unsafeValueIndex-1]),
			slices.Clone(this.levels[*unsafeValueIndex:])...,
		)},
		{append(
			slices.Clone(this.levels[:*unsafeValueIndex]),
			slices.Clone(this.levels[*unsafeValueIndex+1:])...,
		)},
	}

	for _, report := range reportsToTry {
		if report.IsSafe() {
			return true
		}
	}

	return false
}

func (this Report) unsafeValueIndex() *int {
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
			var unsafeValueIndex = i + 1
			return &unsafeValueIndex
		}
	}

	return nil
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
