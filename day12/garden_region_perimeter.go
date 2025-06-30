package day12

import (
	"fmt"
	"github.com/samber/lo"
	"slices"
)

type GardenRegionPerimeter struct {
	vertical   map[int][]Border
	horizontal map[int][]Border
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

func NewGardenRegionPerimeter() GardenRegionPerimeter {
	return GardenRegionPerimeter{
		vertical:   map[int][]Border{},
		horizontal: map[int][]Border{},
	}
}

func (this GardenRegionPerimeter) Length() int {
	var result = 0
	for _, borders := range this.vertical {
		result += len(borders)
	}
	for _, borders := range this.horizontal {
		result += len(borders)
	}
	return result
}

func (this GardenRegionPerimeter) NumberOfSides() int {
	var sides = 0

	for _, verticalRow := range this.vertical {

		slices.SortFunc(verticalRow, func(a Border, b Border) int {
			return a.coordinate.Y - b.coordinate.Y
		})

		var groups = lo.PartitionBy(verticalRow, func(b Border) CardinalDirection {
			return b.direction
		})

		var previousY = -99
		for _, group := range groups {
			for _, b := range group {
				if b.coordinate.Y != previousY+1 {
					sides++
				}
				previousY = b.coordinate.Y
			}
		}

	}

	for _, horizontalRow := range this.horizontal {

		slices.SortFunc(horizontalRow, func(a Border, b Border) int {
			return a.coordinate.X - b.coordinate.X
		})

		var groups = lo.PartitionBy(horizontalRow, func(b Border) CardinalDirection {
			return b.direction
		})

		var previousX = -99
		for _, group := range groups {
			for _, b := range group {
				if b.coordinate.X != previousX+1 {
					sides++
				}
				previousX = b.coordinate.X
			}
		}

	}

	return sides
}

func (this *GardenRegionPerimeter) Add(insideCoordinate Coordinate, outsideCoordinate Coordinate) {
	var borderDirection = cardinalDirectionFor(insideCoordinate, outsideCoordinate)
	var newBorder = Border{direction: borderDirection, coordinate: insideCoordinate}

	switch borderDirection {
	case NORTH:
		fallthrough
	case SOUTH:
		this.horizontal[insideCoordinate.Y] = append(this.horizontal[insideCoordinate.Y], newBorder)
	case EAST:
		fallthrough
	case WEST:
		this.vertical[insideCoordinate.X] = append(this.vertical[insideCoordinate.X], newBorder)
	}
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
