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
	var explorer = MazeMapExplorer{
		maze:         maze,
		reindeer:     reindeer,
		target:       target,
		toVisitStack: []MomentSnapshot{},
		visited:      []Reindeer{},
	}
	explorer.inizializeSnapshotStack()
	return explorer
}

func (this MazeMapExplorer) FindLowestCostToReachTarget() int {
	for len(this.toVisitStack) > 0 {
		var snapshot = this.toVisitStack.PopFirstElement()
		this.reindeer = snapshot.reindeer

		if this.reindeer.Coordinate == this.target {
			return snapshot.cost
		}

		var reindeerDirection = this.reindeer.Direction
		var snapshotCost = snapshot.cost
		this.appendToStackWithCostIfVisitable(reindeerDirection, snapshotCost+1, &snapshot)
		this.appendToStackWithCostIfVisitable(reindeerDirection.Clockwise(), snapshotCost+1001, &snapshot)
		this.appendToStackWithCostIfVisitable(reindeerDirection.CounterClockwise(), snapshotCost+1001, &snapshot)

		this.updateVisited()
	}

	return -1
}

func (this MazeMapExplorer) CoordinatesCountOfBestPaths() int {
	for len(this.toVisitStack) > 0 {
		var snapshot = this.toVisitStack.PopFirstElement()
		this.reindeer = snapshot.reindeer

		if this.reindeer.Coordinate == this.target {
			return snapshot.GetPathLength()
		}

		var reindeerDirection = this.reindeer.Direction
		var snapshotCost = snapshot.cost
		this.appendToStackWithCostIfVisitable(reindeerDirection, snapshotCost+1, &snapshot)
		this.appendToStackWithCostIfVisitable(reindeerDirection.Clockwise(), snapshotCost+1001, &snapshot)
		this.appendToStackWithCostIfVisitable(reindeerDirection.CounterClockwise(), snapshotCost+1001, &snapshot)

		this.updateVisited()
	}

	return -1
}

func (this *MazeMapExplorer) inizializeSnapshotStack() {
	var reindeerDirection = this.reindeer.Direction
	this.appendToStackWithCostIfVisitable(reindeerDirection, 1, nil)
	this.appendToStackWithCostIfVisitable(reindeerDirection.Clockwise(), 1001, nil)
	this.appendToStackWithCostIfVisitable(reindeerDirection.CounterClockwise(), 1001, nil)
	this.appendToStackWithCostIfVisitable(reindeerDirection.Opposite(), 2001, nil)
}

func (this *MazeMapExplorer) appendToStackWithCostIfVisitable(direction Direction, cost int, parentSnapshot *MomentSnapshot) {
	var nextCoordinate = this.reindeer.Coordinate.NextFor(direction)
	if !this.maze.isVisitable(nextCoordinate) {
		return
	}

	var updatedReindeer = Reindeer{nextCoordinate, direction}
	if this.alreadyVisited(updatedReindeer) {
		return
	}

	var momentSnapshot = MomentSnapshot{reindeer: updatedReindeer, cost: cost, parentSnapshot: parentSnapshot}
	this.toVisitStack.AppendSortedByCost(momentSnapshot)
}

func (this *MazeMapExplorer) alreadyVisited(reindeer Reindeer) bool {
	return slices.Contains(this.visited, reindeer)
}

func (this *MazeMapExplorer) updateVisited() {
	this.visited = append(this.visited, this.reindeer)
}
