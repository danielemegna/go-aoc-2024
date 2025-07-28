package day14

func SafetyFactorAfter100Seconds(fileContent string, spaceSize SpaceSize) int {
	var robots = ParseRobotLines(fileContent)
	var space = Space{size: spaceSize, guards: robots}

	space.AfterSeconds(100)

	return space.GetNumberOfRobotsInArea(NORTH_EAST) *
		space.GetNumberOfRobotsInArea(NORTH_WEST) *
		space.GetNumberOfRobotsInArea(SOUTH_EAST) *
		space.GetNumberOfRobotsInArea(SOUTH_WEST)
}
