package day12

import "fmt"

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

func NewBorderFor(coordinate Coordinate, closeCoordinate Coordinate) Border {
	var borderDirection = cardinalDirectionFor(coordinate, closeCoordinate)
	return Border{coordinate: coordinate, direction: borderDirection}
}

func cardinalDirectionFor(coordinate Coordinate, closeCoordinate Coordinate) CardinalDirection {
	if closeCoordinate.Y < coordinate.Y {
		return NORTH
	}
	if closeCoordinate.Y > coordinate.Y {
		return SOUTH
	}
	if closeCoordinate.X < coordinate.X {
		return WEST
	}
	if closeCoordinate.X > coordinate.X {
		return EAST
	}

	panic(fmt.Sprintf("Unexpected cardinalDirectionFor two identical coordinates: [%v]", closeCoordinate))
}

func (this Border) CompareLineOrderFn() func(Border, Border) int {
	if this.direction == NORTH || this.direction == SOUTH {
		return func(a Border, b Border) int {
			return a.coordinate.X - b.coordinate.X
		}
	}

	return func(a Border, b Border) int {
		return a.coordinate.Y - b.coordinate.Y
	}
}
