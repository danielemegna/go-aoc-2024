package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequentialIncreasingReportIsSafe(t *testing.T) {
	var report = Report{[]int{1, 2, 3}}
	assert.True(t, report.IsSafe())
}

func TestSequentialDecreasingReportIsSafe(t *testing.T) {
	var report = Report{[]int{4, 3, 2}}
	assert.True(t, report.IsSafe())
}

func TestNotIncreasingOrDecreasingReportIsNotSafe(t *testing.T) {
	var report = Report{[]int{1, 3, 2}}
	assert.False(t, report.IsSafe())
}

func TestConstantLevelValuesReportIsNotSafe(t *testing.T) {
	var report = Report{[]int{2, 2, 2}}
	assert.False(t, report.IsSafe())
}

func TestReportWithConstantLevelValuesInTheMiddleIsNotSafe(t *testing.T) {
	var report = Report{[]int{1, 2, 3, 3, 4}}
	assert.False(t, report.IsSafe())
}

func TestReportWithAtMostThreeLevelDifferencesIsSafe(t *testing.T) {
	var report = Report{[]int{1, 3, 6, 7, 9}}
	assert.True(t, report.IsSafe())
}

func TestIncreasingReportWithMoreThanThreeLevelDifferenceIsNotSafe(t *testing.T) {
	var report = Report{[]int{1, 2, 7, 8, 9}}
	assert.False(t, report.IsSafe())
}

func TestDecreasingReportWithMoreThanThreeLevelDifferenceIsNotSafe(t *testing.T) {
	var report = Report{[]int{9, 7, 6, 2, 1}}
	assert.False(t, report.IsSafe())
}

func TestSafeReportToleratingSingleBadLevelInTheMiddle(t *testing.T) {
	var reports = []Report{
		{[]int{1, 3, 2, 4, 5}},
		{[]int{8, 6, 4, 4, 1}},
		{[]int{1, 2, 3, 2, 4}},
	}

	for _, report := range reports {
		assert.False(t, report.IsSafe(), report)
		assert.True(t, report.IsSafeWithTolerance(), report)
	}
}

func TestSafeReportToleratingSingleBadLevelAtTheEnd(t *testing.T) {
	var reports = []Report{
		{[]int{1, 3, 6, 1}},
		{[]int{1, 3, 6, 10}},
		{[]int{8, 7, 6, 6}},
		{[]int{8, 7, 6, 2}},
	}

	for _, report := range reports {
		assert.False(t, report.IsSafe(), report)
		assert.True(t, report.IsSafeWithTolerance(), report)
	}
}

func TestUnsafeReportRegardlessToleratingSingleBadLevel(t *testing.T) {
	var reports = []Report{
		{[]int{1, 2, 7, 8, 9}},
		{[]int{9, 7, 6, 2, 1}},
		{[]int{1, 2, 3, 2, 3}},
	}

	for _, report := range reports {
		assert.False(t, report.IsSafe(), report)
		assert.False(t, report.IsSafeWithTolerance(), report)
	}
}

func TestBuildReportFromString(t *testing.T) {
	assert.Equal(t, Report{[]int{1, 3, 2, 4, 5}}, BuildReportFrom("1 3 2 4 5"))
	assert.Equal(t, Report{[]int{1, 3, 6, 7, 9}}, BuildReportFrom("1 3 6 7 9"))
	assert.Equal(t, Report{[]int{7, 6, 4, 2, 1}}, BuildReportFrom("7 6 4 2 1"))
}
