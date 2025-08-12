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
	var start = this.GetRobotPosition()

	var destination = start.NextFor(direction)
	if destination.isOutOfBound() {
		return
	}

	var destinationElement = this.ElementAt(destination)
	if destinationElement == WALL {
		return
	}

	if destinationElement == BOX {
		var shifted = this.shiftBoxesIfPossible(destination, direction)
		if(!shifted) {
			return
		}
	}

	this.setValueAt(start, EMPTY)
	this.setValueAt(destination, ROBOT)
}

func (this WarehouseMap) ElementAt(c Coordinate) MapElement {
	return this[c.y][c.x]
}

func (this WarehouseMap) setValueAt(c Coordinate, e MapElement) {
	this[c.y][c.x] = e
}

func (this WarehouseMap) shiftBoxesIfPossible(boxCoordinate Coordinate, direction Direction) bool {
	var destination = boxCoordinate.NextFor(direction)

	if destination.isOutOfBound() {
		return false
	}
	// NEXT: return false if destination is a WALL
	// NEXT: call recursively if another box in destination

	this.setValueAt(boxCoordinate, EMPTY)
	this.setValueAt(destination, BOX)

	return true
}

func (this WarehouseMap) GetRobotPosition() Coordinate {
	for y, row := range this {
		for x, mapElement := range row {
			if mapElement == ROBOT {
				return Coordinate{x, y}
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
