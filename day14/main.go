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

	var elapsedSeconds = 0
	for {
		space.AfterSeconds(1)
		elapsedSeconds++

		var numberOfRobotsInCenterArea = space.GetNumberOfRobotsInAreaWith(CENTER, 20)
		if numberOfRobotsInCenterArea > 300 {
			// run `go test` with `-v` too see this print
			space.Print()
			return elapsedSeconds
		}
	}
}
