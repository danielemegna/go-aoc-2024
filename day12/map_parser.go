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

	var closeCoordinates = currentCoordinate.closeCoordinates()

	for _, closeCoordinate := range closeCoordinates {
		if !rawMap.hasPlantIn(partialRegion.plant, closeCoordinate) {
			partialRegion.perimeter.Add(currentCoordinate, closeCoordinate)
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

func (rawMap RawGardenMap) hasPlantIn(plant rune, c Coordinate) bool {
	if c.isOutOfBoundsOf(rawMap) {
		return false
	}

	return rawMap[c.Y][c.X] == byte(plant)
}

func rawGardenMapFrom(input string) RawGardenMap {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
