package day15

type MapElement int

const (
	EMPTY MapElement = iota
	BOX
	ROBOT
	WALL
)

type WarehouseMap [][]MapElement

func (this WarehouseMap) MoveRobot(direction Direction) {
	// TODO: extract Coordinate concept?
	var startingX, startingY = this.GetRobotPosition()

	// TODO: remove destination calc duplication
	var destinationX, destinationY int
	switch direction {
	case RIGHT:
		destinationX, destinationY = startingX+1, startingY
	case LEFT:
		destinationX, destinationY = startingX-1, startingY
	case UP:
		destinationX, destinationY = startingX, startingY-1
	case DOWN:
		destinationX, destinationY = startingX, startingY+1
	}

	if destinationX < 0 || destinationY < 0 {
		return
	}

	if this[destinationY][destinationX] == WALL {
		return
	}

	if this[destinationY][destinationX] == BOX {
		this.shiftBoxesIfPossible(destinationX, destinationY, direction)
		// NEXT: exit if shiftBoxesIfPossible returns false
	}

	this[startingY][startingX] = EMPTY
	this[destinationY][destinationX] = ROBOT
}

func (this WarehouseMap) shiftBoxesIfPossible(boxX int, boxY int, direction Direction) bool {
	// TODO: remove destination calc duplication
	var destinationX, destinationY int
	switch direction {
	case RIGHT:
		destinationX, destinationY = boxX+1, boxY
	case LEFT:
		destinationX, destinationY = boxX-1, boxY
	case UP:
		destinationX, destinationY = boxX, boxY-1
	case DOWN:
		destinationX, destinationY = boxX, boxY+1
	}

	// NEXT: return false if destination WALL or out of bound
	// NEXT: call recursively if another box in destination

	this[boxY][boxX] = EMPTY
	this[destinationY][destinationX] = BOX

	return true
}

func (this WarehouseMap) GetRobotPosition() (int, int) {
	for y, row := range this {
		for x, mapElement := range row {
			if mapElement == ROBOT {
				return x, y
			}
		}
	}

	panic("Cannot find the robot in the map")
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (this WarehouseMap) GetBoxesGPSCoordinatesSum() int {
	var sum = 0
	for y, row := range this {
		for x, mapElement := range row {
			if mapElement == BOX {
				sum += (100 * (y + 1)) + x + 1
			}
		}
	}

	return sum
}
