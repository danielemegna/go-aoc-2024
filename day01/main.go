package day01

import (
	"sort"
	"strconv"
	"strings"
)

func TotalDistanceBetweenListsElements(input string) int {
	var rows = rowsFrom(input)
	var firstList, secondList = parseListsFrom(rows)

	sort.Ints(firstList)
	sort.Ints(secondList)

	var totalDifferences = 0
	for i := 0; i < len(firstList); i++ {
		totalDifferences += absoluteDifference(firstList[i], secondList[i])
	}

	return totalDifferences
}

func SimilarityScoreFor(input string) int {
	return -1
}

func parseListsFrom(rows []string) ([]int, []int) {
	var rowsCount = len(rows)
	var firstList = make([]int, 0, rowsCount)
	var secondList = make([]int, 0, rowsCount)
	for _, row := range rows {
		var rowParts = strings.Split(row, "   ")
		firstListNumber, _ := strconv.Atoi(rowParts[0])
		secondListNumber, _ := strconv.Atoi(rowParts[1])
		firstList = append(firstList, firstListNumber)
		secondList = append(secondList, secondListNumber)
	}

	return firstList, secondList
}

func absoluteDifference(n1 int, n2 int) int {
	var difference = n1 - n2
	if difference > 0 {
		return difference
	}

	return -difference
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
