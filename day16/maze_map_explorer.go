package day16

import (
	"slices"
)

type Target = Coordinate

type MazeMapExplorer struct {
	maze         MazeMap
	reindeer     Reindeer
	target       Target
	toVisitStack SnapshotStack
	visited      []Reindeer
}

func NewMazeMapExplorer(maze MazeMap, reindeer Reindeer, target Target) MazeMapExplorer {
	return MazeMapExplorer{
		maze:         maze,
		reindeer:     reindeer,
		target:       target,
		toVisitStack: []MomentSnapshot{},
		visited:      []Reindeer{},
	}
}

func (this MazeMapExplorer) FindLowestCostToReachTarget() int {
	if this.reindeer.ForehandCoordinate() == this.target {
		return 1
	}

	this.InizializeSnapshotStack()

	for len(this.toVisitStack) > 0 {
		var snapshot = this.toVisitStack.PopFirstElement()
		var snapshotReindeer, snapshotCost = snapshot.reindeer, snapshot.cost

		for _, nextDirection := range []Direction{RIGHT, LEFT, UP, DOWN} {
			var nextCoordinate = snapshotReindeer.Coordinate.NextFor(nextDirection)

			if this.maze.IsOutOfBound(nextCoordinate) {
				continue
			}

			if nextDirection == snapshotReindeer.Direction.Opposite() {
				// let's exclude to visit again
				continue
			}

			if this.maze.ElementAt(nextCoordinate) == WALL {
				continue
			}

			var updatedReindeer = Reindeer{nextCoordinate, nextDirection}
			if slices.Contains(this.visited, updatedReindeer) {
				continue
			}

			var nextCoordinateCost = snapshotCost + 1

			if nextDirection != snapshotReindeer.Direction {
				nextCoordinateCost += 1000
			}

			if nextCoordinate == this.target {
				return nextCoordinateCost
			}

			this.toVisitStack.AppendSortedByCost(nextCoordinate, nextDirection, nextCoordinateCost)
		}
		this.visited = append(this.visited, snapshotReindeer)
	}

	return 0
}

func (this *MazeMapExplorer) InizializeSnapshotStack() {
	var reindeerDirection = this.reindeer.Direction
	this.AppendToStackWithCostIfVisitable(reindeerDirection, 1)
	this.AppendToStackWithCostIfVisitable(reindeerDirection.Clockwise(), 1001)
	this.AppendToStackWithCostIfVisitable(reindeerDirection.CounterClockwise(), 1001)
	this.AppendToStackWithCostIfVisitable(reindeerDirection.Opposite(), 2001)
}

func (this *MazeMapExplorer) AppendToStackWithCostIfVisitable(direction Direction, cost int) {
	var nextCoordinate = this.reindeer.Coordinate.NextFor(direction)
	if !this.maze.isVisitable(nextCoordinate) {
		return
	}

	this.toVisitStack.AppendSortedByCost(nextCoordinate, direction, cost)
}
