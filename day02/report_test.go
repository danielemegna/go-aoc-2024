package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReportSafety(t *testing.T) {
	tests := []struct {
		name         string
		levels       []int
		expectedSafe bool
	}{
		{
			name:         "safe: sequential increasing report",
			levels:       []int{1, 2, 3},
			expectedSafe: true,
		},
		{
			name:         "safe: sequential decreasing report",
			levels:       []int{3, 2, 1},
			expectedSafe: true,
		},
		{
			name:         "safe: at most 3 level of distance increasing",
			levels:       []int{1, 3, 6, 7, 9},
			expectedSafe: true,
		},
		{
			name:         "safe: at most 3 level of distance decreasing",
			levels:       []int{9, 7, 4, 1},
			expectedSafe: true,
		},
		{
			name:         "unsafe: repeated level",
			levels:       []int{1, 2, 2, 3},
			expectedSafe: false,
		},
		{
			name:         "unsafe: distance greater than 4, increasing",
			levels:       []int{1, 2, 6, 7},
			expectedSafe: false,
		},
		{
			name:         "unsafe: distance greater than 4, decreasing",
			levels:       []int{7, 6, 2, 1},
			expectedSafe: false,
		},
		{
			name:         "unsafe: non monotonic",
			levels:       []int{1, 2, 3, 2},
			expectedSafe: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := Report{test.levels}
			assert.Equal(t, test.expectedSafe, report.IsSafe(), "IsSafe? ", test.levels)
		})
	}
}

func Test_ReportRecoverability(t *testing.T) {
	tests := []struct {
		name                      string
		levels                    []int
		expectedSafeWithTolerance bool
	}{
		{
			name:                      "single bad value at beginning: too large, sequence increasing",
			levels:                    []int{99, 1, 3, 5},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value at beginning: too large, sequence decreasing",
			levels:                    []int{99, 7, 6, 5},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value at beginning: too small, sequence decreasing",
			levels:                    []int{1, 6, 7, 8},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value at beginning: repeated value",
			levels:                    []int{1, 1, 3, 5},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value at beginning: decreasing in increasing sequence",
			levels:                    []int{8, 6, 7, 8},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value at beginning: increasing in decreasing sequence",
			levels:                    []int{2, 3, 2, 1},
			expectedSafeWithTolerance: true,
		},

		{
			name:                      "single bad value in the middle: too small",
			levels:                    []int{5, 1, 7, 8},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value in the middle: too large",
			levels:                    []int{5, 99, 7, 8},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value in the middle: repeated value",
			levels:                    []int{5, 7, 7, 8},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value in the middle: decreasing in increasing sequence",
			levels:                    []int{5, 4, 6, 7},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value in the end: too small",
			levels:                    []int{7, 8, 9, 1},
			expectedSafeWithTolerance: true,
		},
		{
			name:                      "single bad value in the end: too large",
			levels:                    []int{1, 2, 3, 99},
			expectedSafeWithTolerance: true,
		},

		{
			name:                      "unrecoverable despite single bad value",
			levels:                    []int{1, 2, 99, 2},
			expectedSafeWithTolerance: false,
		},

		{
			name:                      "unrecoverable because of 2 bad values",
			levels:                    []int{1, 2, 99, 3, 99},
			expectedSafeWithTolerance: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			report := Report{test.levels}
			assert.False(t, report.IsSafe(), "IsSafe? ", report)
			assert.Equal(t, test.expectedSafeWithTolerance, report.IsSafeWithTolerance(),
				"SafeWithTolerange? ", test.levels)
		})
	}
}

func TestBuildReportFromString(t *testing.T) {
	assert.Equal(t, Report{[]int{1, 3, 2, 4, 5}}, BuildReportFrom("1 3 2 4 5"))
	assert.Equal(t, Report{[]int{1, 3, 6, 7, 9}}, BuildReportFrom("1 3 6 7 9"))
	assert.Equal(t, Report{[]int{7, 6, 4, 2, 1}}, BuildReportFrom("7 6 4 2 1"))
}
