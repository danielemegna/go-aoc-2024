package day12

import (
	"fmt"
	"slices"
)

type GardenRegionPerimeter []Border

type Border struct {
	direction  CardinalDirection
	coordinate Coordinate
}

type CardinalDirection int

const (
	NORTH CardinalDirection = iota
	SOUTH
	EAST
	WEST
)

func (this GardenRegionPerimeter) Length() int {
	return len(this)
}

func (this GardenRegionPerimeter) NumberOfSides() int {
	var computed = []Border{}
	var sides = 0

	// we have to sort this before ! at least two times !!
	for _, border := range this {
		if slices.Contains(computed, border) {
			continue;
		}

		computed = append(computed, border)

		var maybeCloses = []Border{
			{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X - 1, Y: border.coordinate.Y},
			},
			{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X + 1, Y: border.coordinate.Y},
			},
			{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X, Y: border.coordinate.Y - 1},
			},
			{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X, Y: border.coordinate.Y + 1},
			},
		}

		if slices.ContainsFunc(computed, func(b Border) bool {
			return slices.Contains(maybeCloses, b)
		}) {
			continue
		}

		sides++

		for _, maybeClose := range maybeCloses {
			if slices.Contains(this, maybeClose) {
				computed = append(computed, maybeClose)
			}
		}

	}

	return sides
}

func (this *GardenRegionPerimeter) Add(insideCoordinate Coordinate, outsideCoordinate Coordinate) {
	var borderDirection = cardinalDirectionFor(insideCoordinate, outsideCoordinate)
	var newBorder = Border{direction: borderDirection, coordinate: insideCoordinate}
	*this = append(*this, newBorder)
}

func (this GardenRegionPerimeter) contains(border Border) bool {
	return slices.Contains(this, border)
}

func cardinalDirectionFor(insideCoordinate Coordinate, outsideCoordinate Coordinate) CardinalDirection {
	if outsideCoordinate.Y < insideCoordinate.Y {
		return SOUTH
	}
	if outsideCoordinate.Y > insideCoordinate.Y {
		return NORTH
	}
	if outsideCoordinate.X < insideCoordinate.X {
		return WEST
	}
	if outsideCoordinate.X > insideCoordinate.X {
		return EAST
	}

	panic(fmt.Sprintf("Unexpected cardinalDirectionFor two identical coordinates: [%v]", outsideCoordinate))
}
