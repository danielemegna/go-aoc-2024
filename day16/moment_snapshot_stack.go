package day16

import "slices"

type MomentSnapshot struct {
	reindeer       Reindeer
	cost           int
	parentSnapshot *MomentSnapshot
}

func (this MomentSnapshot) GetPathCoordinates() []Coordinate {
	var pathCoordinates = []Coordinate{}

	var current = &this
	for current != nil {
		pathCoordinates = append(pathCoordinates, current.reindeer.Coordinate)
		current = current.parentSnapshot
	}
	return pathCoordinates
}

type SnapshotStack []MomentSnapshot

func (this *SnapshotStack) PopFirstElement() MomentSnapshot {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *SnapshotStack) AppendSortedByCost(newSnapshot MomentSnapshot) {
	var cost = newSnapshot.cost
	var stack = (*this)
	for index, existingSnapshot := range stack {
		if cost <= existingSnapshot.cost {
			(*this) = slices.Insert(stack, index, newSnapshot)
			return
		}
	}

	(*this) = append(stack, newSnapshot)
}
