package day12

import (
	"slices"

	"github.com/samber/lo"
)

type GardenRegionPerimeter struct {
	vertical   map[int][]Border
	horizontal map[int][]Border
}

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

		for _, group := range groups {
			var previousY = -99
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

		for _, group := range groups {
			var previousX = -99
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

func (this *GardenRegionPerimeter) AddBorder(border Border) {
	switch border.direction {
	case NORTH:
		fallthrough
	case SOUTH:
		this.horizontal[border.coordinate.Y] = append(this.horizontal[border.coordinate.Y], border)
	case EAST:
		fallthrough
	case WEST:
		this.vertical[border.coordinate.X] = append(this.vertical[border.coordinate.X], border)
	}
}
