package day01

import (
	"strconv"
	"strings"
)

func TotalDistanceBetweenListsElements(input string) int {
	var firstList, secondList = parseListsFrom(input)

	firstList.Sort()
	secondList.Sort()

	var totalDifferences = 0
	for i := 0; i < firstList.Length(); i++ {
		totalDifferences += absoluteDifference(firstList.At(i), secondList.At(i))
	}

	return totalDifferences
}

func SimilarityScoreFor(input string) int {
	var firstList, secondList = parseListsFrom(input)
	return firstList.SimilarityScoreWith(secondList)
}

func parseListsFrom(input string) (NumberList, NumberList) {
	var rows = rowsFrom(input)
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

	return NumberList{firstList}, NumberList{secondList}
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
