package day16

import (
	"slices"
)

type MomentSnapshot = struct {
	reindeer Reindeer
	cost     int
}
type SnapshotStack []MomentSnapshot

type Target = Coordinate

func FindLowestCostToReachTarget(maze MazeMap, reindeer Reindeer, target Target) int {
	var stack = SnapshotStack{}
	var completed = []Reindeer{}

	for _, nextDirection := range []Direction{RIGHT, LEFT, UP, DOWN} {
		var nextCoordinate = reindeer.Coordinate.NextFor(nextDirection)

		if nextCoordinate.IsOutOfBoundFor(maze) {
			continue
		}

		if maze.ElementAt(nextCoordinate) == WALL {
			continue
		}

		var cost = 1

		if nextDirection != reindeer.Direction {
			cost += 1000
		}

		if nextDirection == reindeer.Direction.Opposite() {
			cost += 1000
		}

		if nextCoordinate == target {
			return cost
		}

		stack.AppendSortedByCost(MomentSnapshot{
			Reindeer{nextCoordinate, nextDirection},
			cost,
		})
	}
	completed = append(completed, reindeer)

	for len(stack) > 0 {
		var snapshot = stack.PopFirstElement()
		var snapshotReindeer, snapshotCost = snapshot.reindeer, snapshot.cost

		for _, nextDirection := range []Direction{RIGHT, LEFT, UP, DOWN} {
			var nextCoordinate = snapshotReindeer.Coordinate.NextFor(nextDirection)

			if nextCoordinate.IsOutOfBoundFor(maze) {
				continue
			}

			if nextDirection == snapshotReindeer.Direction.Opposite() {
				continue
			}

			if maze.ElementAt(nextCoordinate) == WALL {
				continue
			}

			var updatedReindeer = Reindeer{nextCoordinate, nextDirection}
			if slices.Contains(completed, updatedReindeer) {
				continue
			}

			var nextCoordinateCost = snapshotCost + 1

			if nextDirection != snapshotReindeer.Direction {
				nextCoordinateCost += 1000
			}

			if nextCoordinate == target {
				return nextCoordinateCost
			}

			stack.AppendSortedByCost(MomentSnapshot{updatedReindeer, nextCoordinateCost})
		}
		completed = append(completed, snapshotReindeer)
	}

	return 0
}

func (this *SnapshotStack) PopFirstElement() MomentSnapshot {
	var stack = (*this)
	var pop = stack[0]
	(*this) = stack[1:]
	return pop
}

func (this *SnapshotStack) AppendSortedByCost(newSnapshot MomentSnapshot) {
	var stack = (*this)
	for index, snapshot := range stack {
		if newSnapshot.cost <= snapshot.cost {
			(*this) = slices.Insert(stack, index, newSnapshot)
			return
		}
	}

	(*this) = append(stack, newSnapshot)
}
