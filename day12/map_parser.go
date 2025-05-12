package day12

import "strings"

func ParseGardenMap(fileContent string) GardenMap {
	var result = GardenMap{}

	var rows = rowsFrom(fileContent)
	for y, row := range rows {
		for x, char := range row {
			var plantPerimeter = plantPerimeterFor(x, y, rows)
			var gardenRegion, isPresent = result[char]
			if !isPresent {
				result[char] = GardenRegion{area: 1, perimeter: plantPerimeter}
				continue
			}

			gardenRegion.area++
			gardenRegion.perimeter += plantPerimeter
			result[char] = gardenRegion
		}
	}
	return result
}

func plantPerimeterFor(x, y int, rows []string) int {
	return 4
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
