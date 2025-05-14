package day12

import (
	"slices"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func ParseGardenMap(fileContent string) GardenMap {
	var visitedCoordinates = []Coordinate{}
	var result = GardenMap{}

	var rows = rowsFrom(fileContent)
	for y, row := range rows {
		for x, char := range row {
			var coordinate = Coordinate{x, y}
			if slices.Contains(visitedCoordinates, coordinate) {
				continue
			}

			var partialRegion = GardenRegion{plant: char, area: 0, perimeter: 0}
			var gardenRegion = completeRegionVisit(partialRegion, coordinate, rows, &visitedCoordinates)
			result = append(result, gardenRegion)
		}
	}
	return result
}

func completeRegionVisit(
	partialRegion GardenRegion,
	currentCoordinate Coordinate,
	rows []string,
	visitedCoordinates *[]Coordinate,
) GardenRegion {
	if slices.Contains(*visitedCoordinates, currentCoordinate) {
		return partialRegion
	}

	var x, y = currentCoordinate.X, currentCoordinate.Y
	var plantTypeInByte = rows[y][x]

	*visitedCoordinates = append(*visitedCoordinates, currentCoordinate)
	partialRegion.area++

	var currentCoordinatePerimeter = 0
	if y > 0 {
		if rows[y-1][x] != plantTypeInByte {
			currentCoordinatePerimeter++
		} else {
			partialRegion = completeRegionVisit(
				partialRegion,
				Coordinate{x, y - 1},
				rows,
				visitedCoordinates,
			)
		}
	} else {
		currentCoordinatePerimeter++
	}

	if x > 0 {
		if rows[y][x-1] != plantTypeInByte {
			currentCoordinatePerimeter++
		} else {
			partialRegion = completeRegionVisit(
				partialRegion,
				Coordinate{x - 1, y},
				rows,
				visitedCoordinates,
			)
		}
	} else {
		currentCoordinatePerimeter++
	}
	if x < len(rows[y])-1 {
		if rows[y][x+1] != plantTypeInByte {
			currentCoordinatePerimeter++
		} else {
			partialRegion = completeRegionVisit(
				partialRegion,
				Coordinate{x + 1, y},
				rows,
				visitedCoordinates,
			)
		}
	} else {
		currentCoordinatePerimeter++
	}
	if y < len(rows)-1 {
		if rows[y+1][x] != plantTypeInByte {
			currentCoordinatePerimeter++
		} else {
			partialRegion = completeRegionVisit(
				partialRegion,
				Coordinate{x, y + 1},
				rows,
				visitedCoordinates,
			)
		}
	} else {
		currentCoordinatePerimeter++
	}

	partialRegion.perimeter += currentCoordinatePerimeter
	return partialRegion
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
