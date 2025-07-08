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
