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
	*visitedCoordinates = append(*visitedCoordinates, currentCoordinate)
	partialRegion.area++

	var currentCoordinatePerimeter = 0
	var closeCoordinates = closePlantsFor(currentCoordinate, rows)
	currentCoordinatePerimeter += 4 - len(closeCoordinates)

	for _, closeCoordinate := range closeCoordinates {
		if !samePlantIn(currentCoordinate, closeCoordinate, rows) {
			currentCoordinatePerimeter++
			continue
		}

		if slices.Contains(*visitedCoordinates, closeCoordinate) {
			continue
		}

		partialRegion = completeRegionVisit(
			partialRegion,
			closeCoordinate,
			rows,
			visitedCoordinates,
		)
	}

	partialRegion.perimeter += currentCoordinatePerimeter
	return partialRegion
}

func samePlantIn(c1 Coordinate, c2 Coordinate, rows []string) bool {
	return rows[c1.Y][c1.X] == rows[c2.Y][c2.X]
}

func closePlantsFor(c Coordinate, rows []string) []Coordinate {
	var x, y = c.X, c.Y
	var closeCoordinates = []Coordinate{}
	if x > 0 {
		closeCoordinates = append(closeCoordinates, Coordinate{x - 1, y})
	}
	if y > 0 {
		closeCoordinates = append(closeCoordinates, Coordinate{x, y - 1})
	}
	if x < len(rows[y])-1 {
		closeCoordinates = append(closeCoordinates, Coordinate{x + 1, y})
	}
	if y < len(rows)-1 {
		closeCoordinates = append(closeCoordinates, Coordinate{x, y + 1})
	}
	return closeCoordinates
}

func rowsFrom(input string) []string {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
