package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func TotalDistanceBetweenListsElements(input string) int {
	var rows = getRowsOf(input)

	var firstList []int
	var secondList []int
	for _, row := range rows {
		var rowParts = strings.Split(row, "   ")
		firstListNumber, _ := strconv.Atoi(rowParts[0])
		secondListNumber, _ := strconv.Atoi(rowParts[1])
		firstList = append(firstList, firstListNumber)
		secondList = append(secondList, secondListNumber)
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	var differences = 0
	for i := 0; i < len(firstList); i++ {
		differences += int(math.Abs(float64(firstList[i] - secondList[i])))
	}

	return int(differences)
}

func getRowsOf(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
