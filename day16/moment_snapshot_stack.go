package day16

import "slices"

type MomentSnapshot = struct {
	reindeer Reindeer
	cost     int
}

type SnapshotStack []MomentSnapshot

func NewInizializedSnapshotStack(maze MazeMap, reindeer Reindeer) SnapshotStack {
	var stack = SnapshotStack{}

	// we know Reindeer initial direction always RIGHT
	stack.AppendWithCostIfVisitable(maze, reindeer.Coordinate.NextFor(RIGHT), RIGHT, 1)
	stack.AppendWithCostIfVisitable(maze, reindeer.Coordinate.NextFor(UP), UP, 1001)
	stack.AppendWithCostIfVisitable(maze, reindeer.Coordinate.NextFor(DOWN), DOWN, 1001)
	stack.AppendWithCostIfVisitable(maze, reindeer.Coordinate.NextFor(LEFT), LEFT, 2001)

	return stack
}

func (this *SnapshotStack) PopFirstElement() MomentSnapshot {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *SnapshotStack) AppendWithCostIfVisitable(maze MazeMap, coordinate Coordinate, direction Direction, cost int) {
	if !maze.isVisitable(coordinate) {
		return
	}

	this.AppendSortedByCost(coordinate, direction, cost)
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
