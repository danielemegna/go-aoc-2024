package day14

import "fmt"

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
	CENTER
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
	return this.GetSafetyFactorWith(1)
}

func (this Space) GetSafetyFactorWith(centerThickness int) int {
	return this.GetNumberOfRobotsInAreaWith(NORTH_EAST, centerThickness) *
		this.GetNumberOfRobotsInAreaWith(NORTH_WEST, centerThickness) *
		this.GetNumberOfRobotsInAreaWith(SOUTH_EAST, centerThickness) *
		this.GetNumberOfRobotsInAreaWith(SOUTH_WEST, centerThickness)
}

func (this Space) GetNumberOfRobotsInArea(area SpaceArea) int {
	return this.GetNumberOfRobotsInAreaWith(area, 1)
}

func (this Space) GetNumberOfRobotsInAreaWith(area SpaceArea, centerThickness int) int {
	var minX, maxX, minY, maxY = this.getAreaLimitsFor(area, centerThickness)

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

func (this Space) Print() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("========= Print Space")

	for y := 0; y < this.size.height; y++ {
		for x := 0; x < this.size.width; x++ {
			var hasGuard = this.HasGuardIn(x, y)
			if hasGuard {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("========= End Print Space")
}

func (this Space) getAreaLimitsFor(area SpaceArea, centerThickness int) (int, int, int, int) {
	var width, height = this.size.width, this.size.height

	var minX, maxX int
	switch area {
	case NORTH_WEST, SOUTH_WEST:
		minX, maxX = (0), ((width / 2) - centerThickness)
	case NORTH_EAST, SOUTH_EAST:
		minX, maxX = ((width / 2) + centerThickness), (width - 1)
	case CENTER:
		minX, maxX = ((width / 2) - centerThickness), ((width / 2) + centerThickness)
	}

	var minY, maxY int
	switch area {
	case NORTH_WEST, NORTH_EAST:
		minY, maxY = (0), ((height / 2) - centerThickness)
	case SOUTH_WEST, SOUTH_EAST:
		minY, maxY = ((height / 2) + centerThickness), (height - 1)
	case CENTER:
		minY, maxY = ((height / 2) - centerThickness), ((height / 2) + centerThickness)
	}

	return minX, maxX, minY, maxY
}
