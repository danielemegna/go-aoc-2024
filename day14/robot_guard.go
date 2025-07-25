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
	this.MoreSecondsTick(spaceSize, 1)
}

func (this *RobotGuard) MoreSecondsTick(spaceSize SpaceSize, seconds int) {
	var newX = this.position.x + (this.velocity.horizontal * seconds)
	var newY = this.position.y + (this.velocity.vertical * seconds)

	for newX < 0 {
		newX += spaceSize.width
	}
	for newY < 0 {
		newY += spaceSize.height
	}

	this.position = Position{x: newX % spaceSize.width, y: newY % spaceSize.height}
}
