package day06

func RunGuardWalk(guardMap GuardMap) GuardMap {
	for {
		guardMap = guardMap.GuardWalk()
		if guardMap.IsGuardOutOfBoundaries() {
			break
		}

		guardMap.TurnGuardClockwise()
	}

	return guardMap

}
