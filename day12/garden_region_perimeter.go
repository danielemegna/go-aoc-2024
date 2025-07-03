package day12

import (
	"slices"
)

type GardenRegionPerimeter struct {
	vertical   map[int]map[CardinalDirection][]Border
	horizontal map[int]map[CardinalDirection][]Border
}

func NewGardenRegionPerimeter() GardenRegionPerimeter {
	return GardenRegionPerimeter{
		vertical:   map[int]map[CardinalDirection][]Border{},
		horizontal: map[int]map[CardinalDirection][]Border{},
	}
}

func (this GardenRegionPerimeter) Length() int {
	var result = 0
	for _, directions := range this.vertical {
		for _, borders := range directions {
			result += len(borders)
		}
	}
	for _, directions := range this.horizontal {
		for _, borders := range directions {
			result += len(borders)
		}
	}
	return result
}

func (this GardenRegionPerimeter) NumberOfSides() int {
	var sides = 0

	for _, directions := range this.vertical {
		for _, borders := range directions {
			var previousY = -99
			for _, b := range borders {
				if b.coordinate.Y != previousY+1 {
					sides++
				}
				previousY = b.coordinate.Y
			}
		}

	}

	for _, directions := range this.horizontal {
		for _, borders := range directions {
			var previousX = -99
			for _, b := range borders {
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
		if this.horizontal[border.coordinate.Y] == nil {
			this.horizontal[border.coordinate.Y] = map[CardinalDirection][]Border{}
		}
		this.horizontal[border.coordinate.Y][border.direction] = append(
			this.horizontal[border.coordinate.Y][border.direction], border,
		)
		slices.SortFunc(this.horizontal[border.coordinate.Y][border.direction], func(a Border, b Border) int {
			return a.coordinate.X - b.coordinate.X
		})
	case EAST:
		fallthrough
	case WEST:
		if this.vertical[border.coordinate.X] == nil {
			this.vertical[border.coordinate.X] = map[CardinalDirection][]Border{}
		}
		this.vertical[border.coordinate.X][border.direction] = append(
			this.vertical[border.coordinate.X][border.direction], border,
		)
		slices.SortFunc(this.vertical[border.coordinate.X][border.direction], func(a Border, b Border) int {
			return a.coordinate.Y - b.coordinate.Y
		})
	}
}
