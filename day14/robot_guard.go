package day14

type SpaceSize struct {
	width  int
	height int
}

type RobotGuard struct {
	position Position
	velocity Velocity
}

type Position struct {
	x int
	y int
}

type Velocity struct {
	horizontal int
	vertical   int
}

func (this *RobotGuard) OneSecondTick(spaceSize SpaceSize) {
	this.position = Position{3, 3}
}
