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
	return 0
}
