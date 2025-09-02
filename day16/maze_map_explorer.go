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
	// we know Reindeer initial direction always RIGHT
	var reindeerCoordinate = this.reindeer.Coordinate
	this.AppendToStackWithCostIfVisitable(reindeerCoordinate.NextFor(RIGHT), RIGHT, 1)
	this.AppendToStackWithCostIfVisitable(reindeerCoordinate.NextFor(UP), UP, 1001)
	this.AppendToStackWithCostIfVisitable(reindeerCoordinate.NextFor(DOWN), DOWN, 1001)
	this.AppendToStackWithCostIfVisitable(reindeerCoordinate.NextFor(LEFT), LEFT, 2001)
}

func (this *MazeMapExplorer) AppendToStackWithCostIfVisitable(coordinate Coordinate, direction Direction, cost int) {
	if !this.maze.isVisitable(coordinate) {
		return
	}

	this.toVisitStack.AppendSortedByCost(coordinate, direction, cost)
}
