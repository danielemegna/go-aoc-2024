package day04

import (
	"github.com/samber/lo"
	"strings"
)

func CountXMasOccurrences(fileContent string) int {
	var rows = rowsFrom(fileContent)

	var charactersMap CharactersMap = lo.Map(rows, func(row string, _ int) []string {
		return strings.Split(row, "")
	})

	var result = 0
	for y := 0; y < len(charactersMap); y++ {
		for x := 0; x < len(charactersMap); x++ {
			result += charactersMap.XMasOccurrencesAt(Coordinate{x, y})
		}
	}

	return result
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
