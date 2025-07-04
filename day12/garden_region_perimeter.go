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
	this.InitLinesIfNeeded(border.coordinate)
	if border.direction == NORTH || border.direction == SOUTH {
		this.horizontal[border.coordinate.Y].Add(border)
		return
	}

	this.vertical[border.coordinate.X].Add(border)
}

func (this *GardenRegionPerimeter) InitLinesIfNeeded(coordinate Coordinate) {
	if this.horizontal[coordinate.Y] == nil {
		this.horizontal[coordinate.Y] = map[CardinalDirection][]Border{}
	}
	if this.vertical[coordinate.X] == nil {
		this.vertical[coordinate.X] = map[CardinalDirection][]Border{}
	}
}

func (this PartitionedSortedBordersLine) Add(border Border) {
	this[border.direction] = append(this[border.direction], border)
	slices.SortFunc(this[border.direction], border.CompareLineOrderFn())
}
