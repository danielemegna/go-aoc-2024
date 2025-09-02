package day16

import "slices"

type MomentSnapshot = struct {
	reindeer Reindeer
	cost     int
}

type SnapshotStack []MomentSnapshot

func (this *SnapshotStack) PopFirstElement() MomentSnapshot {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *SnapshotStack) AppendSortedByCost(coordinate Coordinate, direction Direction, cost int) {
	var newSnapshot = MomentSnapshot{
		reindeer: Reindeer{coordinate, direction},
		cost:     cost,
	}
	var stack = (*this)
	for index, existingSnapshot := range stack {
		if cost <= existingSnapshot.cost {
			(*this) = slices.Insert(stack, index, newSnapshot)
			return
		}
	}

	(*this) = append(stack, newSnapshot)
}
