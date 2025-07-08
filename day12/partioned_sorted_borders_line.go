package day12

import "slices"

type PartitionedSortedBordersLine map[CardinalDirection][]Border

func (this PartitionedSortedBordersLine) Add(border Border) {
	this[border.direction] = append(this[border.direction], border)
	slices.SortFunc(this[border.direction], border.CompareLineOrderFn())
}

func (this PartitionedSortedBordersLine) NumberOfDifferentSides() int {
	if len(this) == 0 {
		return 0
	}

	var sides = 0
	for _, bordersWithSameDirection := range this {
		if len(bordersWithSameDirection) == 0 {
			continue
		}

		var borderCompareFn = bordersWithSameDirection[0].CompareLineOrderFn()

		sides++
		for i := range len(bordersWithSameDirection) - 1 {
			var current = bordersWithSameDirection[i]
			var next = bordersWithSameDirection[i+1]
			var areCloseBorderOfSameSide = borderCompareFn(next, current) == 1
			if !areCloseBorderOfSameSide {
				sides++
			}
		}
	}

	return sides
}
