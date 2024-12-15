package day02

import "strings"

func SafeReportsCount(fileContent string) int {
	var safeReportsCount = 0
	var rows = rowsFrom(fileContent)

	for _, row := range rows {
		var report = BuildReportFrom(row)	
		if(report.IsSafe()) {
			safeReportsCount++
		}
	}

	return safeReportsCount
}

func SafeReportsCountWithTollerance(fileContent string) int {
	return 4
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
