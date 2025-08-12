package day15

type WarehouseMap [][]MapElement

type MapElement int

const (
	EMPTY MapElement = iota
	BOX
	ROBOT
	WALL
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (this WarehouseMap) MoveRobot(direction Direction) {
	var start = this.GetRobotPosition()

	var destination = start.NextFor(direction)
	if destination.isOutOfBound(this.MapWidth(), this.MapHeigth()) {
		return
	}

	var destinationElement = this.ElementAt(destination)
	if destinationElement == WALL {
		return
	}

	if destinationElement == BOX {
		var shifted = this.shiftBoxesIfPossible(destination, direction)
		if !shifted {
			return
		}
	}

	this.setValueAt(start, EMPTY)
	this.setValueAt(destination, ROBOT)
}

func (this WarehouseMap) shiftBoxesIfPossible(boxCoordinate Coordinate, direction Direction) bool {
	var destination = boxCoordinate.NextFor(direction)

	if destination.isOutOfBound(this.MapWidth(), this.MapHeigth()) {
		return false
	}

	var destinationElement = this.ElementAt(destination)
	if destinationElement == WALL {
		return false
	}

	if destinationElement == BOX {
		var shifted = this.shiftBoxesIfPossible(destination, direction)
		if !shifted {
			return false
		}
	}

	this.setValueAt(boxCoordinate, EMPTY)
	this.setValueAt(destination, BOX)

	return true
}

func (this WarehouseMap) MapHeigth() int                        { return len(this) }
func (this WarehouseMap) MapWidth() int                         { return len(this[0]) } // unsupported empty maps
func (this WarehouseMap) ElementAt(c Coordinate) MapElement     { return this[c.y][c.x] }
func (this WarehouseMap) setValueAt(c Coordinate, e MapElement) { this[c.y][c.x] = e }

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
