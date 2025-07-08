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
	for _, lineOfPartionedBorders := range this.vertical {
		for _, borders := range lineOfPartionedBorders {
			result += len(borders)
		}
	}
	for _, lineOfPartionedBorders := range this.horizontal {
		for _, borders := range lineOfPartionedBorders {
			result += len(borders)
		}
	}
	return result
}

func (this GardenRegionPerimeter) NumberOfSides() int {
	var sides = 0

	for _, lineOfPartitionedBorders := range this.vertical {
		sides+= lineOfPartitionedBorders.NumberOfDifferentSides()
	}

	for _, lineOfPartitionedBorders := range this.horizontal {
		sides+= lineOfPartitionedBorders.NumberOfDifferentSides()
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

func (this PartitionedSortedBordersLine) NumberOfDifferentSides() int {
	if(len(this) == 0) {
		return 0
	}

	var sides = 0
	for _, bordersWithSameDirection := range this {
		for i := 1; i < len(bordersWithSameDirection); i++ {
			var current = bordersWithSameDirection[i]
			var previous = bordersWithSameDirection[i-1]
			var borderCompareFn = current.CompareLineOrderFn()
			if borderCompareFn(current, previous) != 1 {
				sides++
			}
		}
		sides++
	}
	return sides
}
