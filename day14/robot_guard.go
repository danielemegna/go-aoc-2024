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
	var newX = this.position.x + this.velocity.horizontal + spaceSize.width
	var newY = this.position.y + this.velocity.vertical + spaceSize.height
	this.position = Position{x: newX % spaceSize.width, y: newY % spaceSize.height}
}
