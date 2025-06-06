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

			var newGardenRegion = NewGardenRegion(char)
			var gardenRegion = completeRegionVisit(newGardenRegion, coordinate, rawMap, &visitedCoordinates)
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

	for _, closeCoordinate := range currentCoordinate.closeCoordinates() {
		if !rawMap.samePlantIn(currentCoordinate, closeCoordinate) {
			partialRegion.perimeter.Add(closeCoordinate, currentCoordinate)
			continue
		}

		if slices.Contains(*visitedCoordinates, closeCoordinate) {
			continue
		}

		partialRegion = completeRegionVisit(
			partialRegion,
			closeCoordinate,
			rawMap,
			visitedCoordinates,
		)
	}

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
