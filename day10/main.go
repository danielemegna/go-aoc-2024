package day10

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func SumOfTrailheadScores(fileContent string) int {
	var topographicMap = parseTopographicMap(fileContent)
	var trailheads = topographicMap.FindTrailheads()
	return lo.SumBy(trailheads, func(a Trailhead) int {
		return a.GetScore()
	})
}

func parseTopographicMap(fileContent string) TopographicMap {
	var fileRows = rowsFrom(fileContent)
	return lo.Map(fileRows, func(row string, _ int) []int {
		return lo.Map([]rune(row), func(char rune, _ int) int {
			var charToInt, _ = strconv.Atoi(string(char))
			return charToInt
		})
	})
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
