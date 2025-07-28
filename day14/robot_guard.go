package day14

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

func (this *RobotGuard) AfterOneSecond(spaceSize SpaceSize) {
	this.AfterSeconds(1, spaceSize)
}

func (this *RobotGuard) AfterSeconds(seconds int, spaceSize SpaceSize) {
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
