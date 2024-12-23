package day04

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

func CountXMasOccurrences(fileContent string) int {
	var rows = rowsFrom(fileContent)

	var matrix CharactersMap = lo.Map(rows, func(row string, _ int) []string {
		return strings.Split(row, "")
	})

	fmt.Println(matrix)
	fmt.Println(matrix[0])
	fmt.Println(matrix[1][1])

	return 18
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
