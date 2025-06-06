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
		return false
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
	if outsideCoordinate.Y > insideCoordinate.Y {
		return NORTH
	}
	if outsideCoordinate.X < insideCoordinate.X {
		return WEST
	}
	if outsideCoordinate.X > insideCoordinate.X {
		return EAST
	}

	panic("Unexpected cardinalDirectionFor two identical coordinates")
}
