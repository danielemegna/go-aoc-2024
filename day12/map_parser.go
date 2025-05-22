package day12

import (
	"slices"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type RawGardenMap = []string

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
	var closeCoordinates = closeCoordinatesFor(currentCoordinate, rawMap)
	var outOfMapBorders = 4 - len(closeCoordinates)
	currentCoordinatePerimeter += outOfMapBorders

	for _, closeCoordinate := range closeCoordinates {
		if !samePlantIn(currentCoordinate, closeCoordinate, rawMap) {
			currentCoordinatePerimeter++
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

	partialRegion.perimeter += currentCoordinatePerimeter
	return partialRegion
}

func samePlantIn(c1 Coordinate, c2 Coordinate, rawMap RawGardenMap) bool {
	return rawMap[c1.Y][c1.X] == rawMap[c2.Y][c2.X]
}

func closeCoordinatesFor(c Coordinate, rawMap RawGardenMap) []Coordinate {
	// check performance initializing with the four coordinates and applying a filter later
	var closeCoordinates = []Coordinate{}
	if c.X > 0 {
		closeCoordinates = append(closeCoordinates, Coordinate{c.X - 1, c.Y})
	}
	if c.Y > 0 {
		closeCoordinates = append(closeCoordinates, Coordinate{c.X, c.Y - 1})
	}
	if c.X < len(rawMap[c.Y])-1 {
		closeCoordinates = append(closeCoordinates, Coordinate{c.X + 1, c.Y})
	}
	if c.Y < len(rawMap)-1 {
		closeCoordinates = append(closeCoordinates, Coordinate{c.X, c.Y + 1})
	}
	return closeCoordinates
}

func rawGardenMapFrom(input string) RawGardenMap {
	var rows = strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
