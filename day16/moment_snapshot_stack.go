package day16

import "slices"

type MomentSnapshot struct {
	reindeer       Reindeer
	cost           int
	parentSnapshot *MomentSnapshot
}

func (this MomentSnapshot) GetPathLength() int {
	var s = &this
	var length = 0
	for s.parentSnapshot != nil {
		length++
		s = s.parentSnapshot
	}
	return length + 1
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
