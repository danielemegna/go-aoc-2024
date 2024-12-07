package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequentialIncreasingReportIsValid(t *testing.T) {
	var report = Report{[]int{1, 2, 3}}
	assert.True(t, report.IsValid())
}

func TestNotConstantIncreasingReportIsInvalid(t *testing.T) {
	var report = Report{[]int{1, 3, 2}}
	assert.False(t, report.IsValid())
}

func TestConstantReportIsNotValid(t *testing.T) {
	var report = Report{[]int{2, 2, 2}}
	assert.False(t, report.IsValid())
}

func TestReportWithTwoConstantValueInTheMiddleIsNotValid(t *testing.T) {
	var report = Report{[]int{1, 2, 3, 3, 4}}
	assert.False(t, report.IsValid())
}

func TestSequentialDecreasingReportIsValid(t *testing.T) {
	var report = Report{[]int{4, 3, 2}}
	assert.True(t, report.IsValid())
}

