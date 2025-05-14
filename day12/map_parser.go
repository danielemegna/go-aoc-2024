package day12

import "strings"

func ParseGardenMap(fileContent string) GardenMap {
	var result = GardenMap{}

	var rows = rowsFrom(fileContent)
	for y, row := range rows {
		for x, char := range row {
			var plantPerimeter = plantPerimeterFor(x, y, rows)
			result = append(result, GardenRegion{plant: char, area: 1, perimeter: plantPerimeter})
		}
	}
	return result
}

func plantPerimeterFor(x int, y int, rows []string) int {
	var perimeter = 0
	var plantType = rows[y][x]

	if y > 0 {
		if rows[y-1][x] != plantType {
			perimeter++
		}
	} else {
		perimeter++
	}

	if x > 0 {
		if rows[y][x-1] != plantType {
			perimeter++
		}
	} else {
		perimeter++
	}
	if x < len(rows[y])-1 {
		if rows[y][x+1] != plantType {
			perimeter++
		}
	} else {
		perimeter++
	}
	if y < len(rows)-1 {
		if rows[y+1][x] != plantType {
			perimeter++
		}
	} else {
		perimeter++
	}
	return perimeter
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
