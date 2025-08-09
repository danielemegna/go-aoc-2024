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
	var startingX = 0
	var startingY = 2

	var destinationX, destinationY int
	switch direction {
		case RIGHT:
			destinationX, destinationY = startingX+1, startingY
		case LEFT:
			destinationX, destinationY = startingX-1, startingY
		case UP:
			destinationX, destinationY = startingX, startingY-1
		//case DOWN:
	}

	if(destinationX < 0 || destinationY < 0) {
		return
	}

	this[startingY][startingX] = EMPTY
	this[destinationY][destinationX] = ROBOT
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
