package day14

type SpaceSize struct {
	width  int
	height int
}

type SpaceArea int

const (
	NORTH_WEST SpaceArea = iota
	NORTH_EAST
	SOUTH_WEST
	SOUTH_EAST
)

type Space struct {
	size   SpaceSize
	guards []RobotGuard
}

func (this Space) AfterSeconds(seconds int) {
	for i := range this.guards {
		this.guards[i].AfterSeconds(seconds, this.size)
	}
}

func (this Space) GetSafetyFactor() int {
	return this.GetNumberOfRobotsInArea(NORTH_EAST) *
		this.GetNumberOfRobotsInArea(NORTH_WEST) *
		this.GetNumberOfRobotsInArea(SOUTH_EAST) *
		this.GetNumberOfRobotsInArea(SOUTH_WEST)
}

func (this Space) GetNumberOfRobotsInArea(area SpaceArea) int {
	var minX, maxX, minY, maxY = this.getAreaLimitsFor(area)

	var count = 0
	for _, guard := range this.guards {
		if guard.position.x >= minX && guard.position.x <= maxX &&
			guard.position.y >= minY && guard.position.y <= maxY {
			count++
		}
	}

	return count
}

func (this Space) HasGuardIn(x int, y int) bool {
	// NEXT: improve this keeping an up-to-date boolean matrix of presences
	for _, guard := range this.guards {
		if guard.position.x == x && guard.position.y == y {
			return true
		}
	}

	return false
}

func (this Space) getAreaLimitsFor(area SpaceArea) (int, int, int, int) {
	var width, height = this.size.width, this.size.height

	var minX, maxX int
	switch area {
	case NORTH_WEST, SOUTH_WEST:
		minX, maxX = (0), ((width / 2) - 1)
	case NORTH_EAST, SOUTH_EAST:
		minX, maxX = ((width / 2) + 1), (width - 1)
	}

	var minY, maxY int
	switch area {
	case NORTH_WEST, NORTH_EAST:
		minY, maxY = (0), ((height / 2) - 1)
	case SOUTH_WEST, SOUTH_EAST:
		minY, maxY = ((height / 2) + 1), (height - 1)
	}

	return minX, maxX, minY, maxY
}
