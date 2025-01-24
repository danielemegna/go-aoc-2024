package day06

import "slices"

func RunGuardWalk(guardMap GuardMap) GuardMap {
	var obstacleHits = map[Coordinate][]Direction{}
	var appendObstacleHit = func(c Coordinate, d Direction) {
		var directions, _ = obstacleHits[c]
		obstacleHits[c] = append(directions, d)
	}
	var containsObstacleHit = func(c Coordinate, d Direction) bool {
		var directions, keyFound = obstacleHits[c]
		return keyFound && slices.Contains(directions, d)
	}

	for {
		guardMap = guardMap.GuardWalk()
		if guardMap.IsGuardOutOfBoundaries() {
			break
		}

		var guardPosition = guardMap.guard.position
		var guardDirection = guardMap.guard.direction
		if containsObstacleHit(guardPosition, guardDirection) {
			break
		}

		appendObstacleHit(guardPosition, guardDirection)
		guardMap.TurnGuardClockwise()
	}

	return guardMap
}

func RunGuardWalkWithExtraObstacle(guardMap GuardMap, obstacleCoordinate Coordinate) GuardMap {
	guardMap.AddObstacle(obstacleCoordinate)
	return RunGuardWalk(guardMap)
}
