package day16

import (
	"math"
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
	var cost = 0

	for _, nextDirection := range []Direction{RIGHT, LEFT, UP, DOWN} {
		var nextCoordinate = reindeer.Coordinate.NextFor(nextDirection)

		if nextCoordinate.IsOutOfBoundFor(maze) {
			continue
		}

		if maze.ElementAt(nextCoordinate) == WALL {
			continue
		}

		cost += 1

		if nextDirection != reindeer.Direction {
			cost += 1000
		}

		if nextDirection == reindeer.Direction.Opposite() {
			cost += 1000
		}

		if nextCoordinate == target {
			return cost
		}

		stack.Append(MomentSnapshot{
			Reindeer{nextCoordinate, nextDirection},
			cost,
		})
	}
	completed = append(completed, reindeer)

	for len(stack) > 0 {
		var snapshot = stack.PopMinimumCost()
		var snapshotReindeer, snapshotCost = snapshot.reindeer, snapshot.cost

		for _, nextDirection := range []Direction{RIGHT, LEFT, UP, DOWN} {
			var nextCoordinate = snapshotReindeer.Coordinate.NextFor(nextDirection)

			if nextCoordinate.IsOutOfBoundFor(maze) {
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

			if nextDirection == snapshotReindeer.Direction.Opposite() {
				nextCoordinateCost += 1000
			}

			if nextCoordinate == target {
				return nextCoordinateCost
			}

			stack.Append(MomentSnapshot{updatedReindeer, nextCoordinateCost})
		}
		completed = append(completed, snapshotReindeer)
	}

	return 0
}

func (this *SnapshotStack) PopMinimumCost() MomentSnapshot {
	var stack = (*this)
	var minIndex int = -1
	var minCost int = math.MaxInt64
	for index, snapshot := range stack {
		if snapshot.cost < minCost {
			minCost = snapshot.cost
			minIndex = index
		}
	}

	var pop = stack[minIndex]
	(*this) = slices.Delete(stack, minIndex, minIndex+1)

	return pop
}

func (this *SnapshotStack) Append(snapshot MomentSnapshot) {
	(*this) = append((*this), snapshot)
}
