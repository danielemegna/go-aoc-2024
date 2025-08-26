package day16

import (
	"fmt"
)

type VisitedPair = struct {
	Coordinate
	Reindeer
}

func FindLowestCostToReachTarget(maze MazeMap, reindeer Reindeer, target Coordinate) int {
	var visitedWithCost = map[VisitedPair]int{}

	var closeCoordinates = reindeer.Coordinate.CloseCoordinates()

	for _, currentCoordinate := range closeCoordinates {
		if currentCoordinate.IsOutOfBoundFor(maze) {
			continue
		}

		if maze.ElementAt(currentCoordinate) == WALL {
			continue
		}

		var cost = 1
		if reindeer.ForehandCoordinate() != currentCoordinate {
			cost += 1000
		}

		if currentCoordinate == reindeer.OppositeCoordinate() {
			cost += 1000
		}

		if currentCoordinate == target {
			return cost
		}

		visitedWithCost[VisitedPair{currentCoordinate, reindeer}] = cost
	}

	fmt.Printf("%+v\n", visitedWithCost)
	return 0
}
