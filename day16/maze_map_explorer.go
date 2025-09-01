package day16

import (
	"slices"
)

type Target = Coordinate

func FindLowestCostToReachTarget(maze MazeMap, reindeer Reindeer, target Target) int {
	if reindeer.ForehandCoordinate() == target {
		return 1
	}

	var stack = NewInizializedSnapshotStack(maze, reindeer)
	var fullyVisited = []Reindeer{reindeer}

	for len(stack) > 0 {
		var snapshot = stack.PopFirstElement()
		var snapshotReindeer, snapshotCost = snapshot.reindeer, snapshot.cost

		for _, nextDirection := range []Direction{RIGHT, LEFT, UP, DOWN} {
			var nextCoordinate = snapshotReindeer.Coordinate.NextFor(nextDirection)

			if maze.IsOutOfBound(nextCoordinate) {
				continue
			}

			if nextDirection == snapshotReindeer.Direction.Opposite() {
				continue
			}

			if maze.ElementAt(nextCoordinate) == WALL {
				continue
			}

			var updatedReindeer = Reindeer{nextCoordinate, nextDirection}
			if slices.Contains(fullyVisited, updatedReindeer) {
				continue
			}

			var nextCoordinateCost = snapshotCost + 1

			if nextDirection != snapshotReindeer.Direction {
				nextCoordinateCost += 1000
			}

			if nextCoordinate == target {
				return nextCoordinateCost
			}

			stack.AppendSortedByCost(nextCoordinate, nextDirection, nextCoordinateCost)
		}
		fullyVisited = append(fullyVisited, snapshotReindeer)
	}

	return 0
}
