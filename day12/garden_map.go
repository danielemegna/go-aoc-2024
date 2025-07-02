package day12

import "fmt"

type GardenMap []GardenRegion

type GardenRegion struct {
	plant     rune
	area      int
	perimeter GardenRegionPerimeter
}

func NewGardenRegion(plant rune) GardenRegion {
	return GardenRegion{
		plant:     plant,
		area:      0,
		perimeter: NewGardenRegionPerimeter(),
	}
}

func (this GardenRegion) AddPerimeterPiece(coordinate Coordinate, closeCoordinate Coordinate) {
	var borderDirection = cardinalDirectionFor(coordinate, closeCoordinate)
	this.perimeter.AddBorder(coordinate, borderDirection)
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
