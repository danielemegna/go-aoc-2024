package day16

import "slices"

type Target = Coordinate

type MazeMapExplorer struct {
	maze         MazeMap
	reindeer     Reindeer
	target       Target
	toVisitStack SnapshotStack
	visited      []Reindeer
}

func NewMazeMapExplorer(mazeMap MazeMap, startReindeer Reindeer, target Target) MazeMapExplorer {
	var explorer = MazeMapExplorer{
		maze:     mazeMap,
		reindeer: startReindeer,
		target:   target,
		toVisitStack: []MomentSnapshot{
			{reindeer: startReindeer, cost: 0, parentSnapshot: nil},
		},
		visited: []Reindeer{},
	}
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
	var coordinatesOfBestPaths = map[Coordinate]struct{}{}
	var bestPathsCost = 999999

	for len(this.toVisitStack) > 0 {
		var snapshot = this.toVisitStack.PopFirstElement()
		if snapshot.cost > bestPathsCost {
			break
		}

		if snapshot.reindeer.Coordinate == this.target {
			bestPathsCost = snapshot.cost
			for _, c := range snapshot.GetPathCoordinates() {
				coordinatesOfBestPaths[c] = struct{}{}
			}
			continue
		}

		var reindeerDirection = snapshot.reindeer.Direction
		var snapshotCost = snapshot.cost
		this.reindeer = snapshot.reindeer
		this.appendToStackWithCostIfVisitable(reindeerDirection, snapshotCost+1, &snapshot)
		this.appendToStackWithCostIfVisitable(reindeerDirection.Clockwise(), snapshotCost+1001, &snapshot)
		this.appendToStackWithCostIfVisitable(reindeerDirection.CounterClockwise(), snapshotCost+1001, &snapshot)

		this.updateVisited()
	}

	return len(coordinatesOfBestPaths)
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
