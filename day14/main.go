package day14

func SafetyFactorAfter100Seconds(fileContent string, spaceSize SpaceSize) int {
	var robots = ParseRobotLines(fileContent)
	var space = Space{size: spaceSize, guards: robots}

	space.AfterSeconds(100)

	return space.GetSafetyFactor()
}
