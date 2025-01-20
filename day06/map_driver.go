package day06

import "slices"

func RunGuardWalk(guardMap GuardMap) GuardMap {
	var obstacleHits = []Coordinate{}

	for {
		guardMap = guardMap.GuardWalk()
		if guardMap.IsGuardOutOfBoundaries() {
			break
		}

		var guardPosition = guardMap.guard.position
		if slices.Contains(obstacleHits, guardPosition) {
			break
		}

		obstacleHits = append(obstacleHits, guardPosition)
		guardMap.TurnGuardClockwise()
	}

	return guardMap

}
