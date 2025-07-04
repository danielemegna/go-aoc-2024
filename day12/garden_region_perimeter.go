package day12

import (
	"slices"
)

type PartitionedSortedBordersLine map[CardinalDirection][]Border

type GardenRegionPerimeter struct {
	vertical   map[int]PartitionedSortedBordersLine
	horizontal map[int]PartitionedSortedBordersLine
}

func NewGardenRegionPerimeter() GardenRegionPerimeter {
	return GardenRegionPerimeter{
		vertical:   map[int]PartitionedSortedBordersLine{},
		horizontal: map[int]PartitionedSortedBordersLine{},
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
	if border.direction == NORTH || border.direction == SOUTH {
		if this.horizontal[border.coordinate.Y] == nil {
			this.horizontal[border.coordinate.Y] = map[CardinalDirection][]Border{}
		}
		this.horizontal[border.coordinate.Y].Add(border)
		return
	}

	if this.vertical[border.coordinate.X] == nil {
		this.vertical[border.coordinate.X] = map[CardinalDirection][]Border{}
	}
	this.vertical[border.coordinate.X].Add(border)
}

func (this PartitionedSortedBordersLine) Add(border Border) {
	this[border.direction] = append(this[border.direction], border)
	slices.SortFunc(this[border.direction], func(a Border, b Border) int {
		return a.CompareTo(b)
	})
}
