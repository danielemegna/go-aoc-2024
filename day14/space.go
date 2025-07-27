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

func (this Space) getAreaLimitsFor(area SpaceArea) (int, int, int, int) {
	var width, height = this.size.width, this.size.height
	switch area {
	case NORTH_EAST:
		return (width / 2) + 1, width - 1, 0, (height / 2) - 1
	case SOUTH_EAST:
		return (width / 2) + 1, width - 1, (height / 2) + 1, height - 1
	default:
		return -1, -1, -1, -1
	}
}
