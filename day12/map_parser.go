package day12

import (
	"slices"
	"strings"
)


type RawGardenMap []string

func ParseGardenMap(fileContent string) GardenMap {
	var visitedCoordinates = []Coordinate{}
	var result = GardenMap{}

	var rawMap = rawGardenMapFrom(fileContent)
	for y, row := range rawMap {
		for x, char := range row {
			var coordinate = Coordinate{x, y}
			if slices.Contains(visitedCoordinates, coordinate) {
				continue
			}

			var partialRegion = GardenRegion{plant: char, area: 0, perimeter: 0}
			var gardenRegion = completeRegionVisit(partialRegion, coordinate, rawMap, &visitedCoordinates)
			result = append(result, gardenRegion)
		}
	}
	return result
}

func completeRegionVisit(
	partialRegion GardenRegion,
	currentCoordinate Coordinate,
	rawMap RawGardenMap,
	visitedCoordinates *[]Coordinate,
) GardenRegion {
	*visitedCoordinates = append(*visitedCoordinates, currentCoordinate)
	partialRegion.area++

	var currentCoordinatePerimeter = 0
	for _, closeCoordinate := range currentCoordinate.closeCoordinates() {
		if !rawMap.samePlantIn(currentCoordinate, closeCoordinate) {
			currentCoordinatePerimeter++
			continue
		}

		if !slices.Contains(*visitedCoordinates, closeCoordinate) {
			partialRegion = completeRegionVisit(
				partialRegion,
				closeCoordinate,
				rawMap,
				visitedCoordinates,
			)
		}
	}

	partialRegion.perimeter += currentCoordinatePerimeter
	return partialRegion
}

func (rawMap RawGardenMap) samePlantIn(c1 Coordinate, c2 Coordinate) bool {
	if c1.isOutOfBoundsOf(rawMap) || c2.isOutOfBoundsOf(rawMap) {
		return false
	}

	return rawMap[c1.Y][c1.X] == rawMap[c2.Y][c2.X]
}


func rawGardenMapFrom(input string) RawGardenMap {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
