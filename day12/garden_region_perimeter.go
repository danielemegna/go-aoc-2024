package day12

import "slices"

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
	this.borders = append(this.borders, border)

	if !this.isSideAlreadyCounted(border) {
		this.sides++
	}
}

func (this GardenRegionPerimeter) isSideAlreadyCounted(border Border) bool {
	switch border.direction {
	case EAST:
		fallthrough
	case WEST:
		fallthrough
	case NORTH:
		fallthrough
	case SOUTH:
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
		fallthrough
	default:
		return false
	}
}

func (this GardenRegionPerimeter) contains(border Border) bool {
	return slices.Contains(this.borders, border)
}

func cardinalDirectionFor(outsideCoordinate Coordinate, insideCoordinate Coordinate) CardinalDirection {
	if outsideCoordinate.Y < insideCoordinate.Y {
		return SOUTH
	}
	return NORTH
}
