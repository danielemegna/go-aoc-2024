package day12

import (
	"fmt"
	"slices"
)

type GardenRegionPerimeter struct {
	borders []Border
	sides   int
}

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
	return len(this.borders)
}

func (this GardenRegionPerimeter) NumberOfSides() int {
	return this.sides
}

func (this *GardenRegionPerimeter) Add(outsideCoordinate Coordinate, insideCoordinate Coordinate) {
	var direction = cardinalDirectionFor(outsideCoordinate, insideCoordinate)
	var border = Border{direction: direction, coordinate: insideCoordinate}

	if !this.isSideAlreadyCounted(border) {
		this.sides++
	}

	this.borders = append(this.borders, border)
}

func (this GardenRegionPerimeter) isSideAlreadyCounted(border Border) bool {
	switch border.direction {
	case EAST, WEST:
		if (this.contains(
			Border{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X, Y: border.coordinate.Y - 1},
			})) {
			return true
		}
		if (this.contains(
			Border{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X, Y: border.coordinate.Y + 1},
			})) {
			return true
		}
		return false
	case NORTH, SOUTH:
		if (this.contains(
			Border{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X - 1, Y: border.coordinate.Y},
			})) {
			return true
		}
		if (this.contains(
			Border{
				direction:  border.direction,
				coordinate: Coordinate{X: border.coordinate.X + 1, Y: border.coordinate.Y},
			})) {
			return true
		}
		return false
	}

	panic(fmt.Sprintf("Unexpected isSideAlreadyCounted with border direction: [%v]", border))
}

func (this GardenRegionPerimeter) contains(border Border) bool {
	return slices.Contains(this.borders, border)
}

func cardinalDirectionFor(outsideCoordinate Coordinate, insideCoordinate Coordinate) CardinalDirection {
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
