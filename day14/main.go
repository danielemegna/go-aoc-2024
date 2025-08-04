package day14

func SafetyFactorAfter100Seconds(fileContent string, spaceSize SpaceSize) int {
	var robots = ParseRobotLines(fileContent)
	var space = Space{size: spaceSize, guards: robots}

	space.AfterSeconds(100)

	return space.GetSafetyFactor()
}

func SecondsToWaitToSeeTheChristmasTreeEasterEgg(fileContent string, spaceSize SpaceSize) int {
	var robots = ParseRobotLines(fileContent)
	var space = Space{size: spaceSize, guards: robots}
	var patternRepeatLimit = 10403

	var elapsedSeconds = 0
	for elapsedSeconds < patternRepeatLimit {
		space.AfterSeconds(1)
		elapsedSeconds++

		// NEXT: find a better peculiarity to recognize the tree
		var safetyFactor = space.GetSafetyFactorWith(36)
		if(safetyFactor == 75) {
			// run `go test` with `-v` too see this print
			space.Print()
			break
		}
	}

	return elapsedSeconds
}
