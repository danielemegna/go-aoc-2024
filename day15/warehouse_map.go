package day15

type WarehouseMap [][]MapElement

func (this WarehouseMap) MoveRobot(direction Direction) {
	var startCoordinate = this.GetRobotPosition()
	this.shiftElementsIfPossible(startCoordinate, direction)
}

func (this WarehouseMap) shiftElementsIfPossible(startCoordinate Coordinate, direction Direction) bool {
	var destination = startCoordinate.NextFor(direction)

	if destination.isOutOfBoundFor(this) {
		return false
	}

	var destinationElement = this.ElementAt(destination)
	if destinationElement == WALL {
		return false
	}

	if destinationElement.IsSmallBox() {
		var shifted = this.shiftElementsIfPossible(destination, direction)
		if !shifted {
			return false
		}
	}

	if destinationElement.IsBigBox() {
		if direction.IsHorizontal() {
			var shifted = this.shiftElementsIfPossible(destination, direction)
			if !shifted {
				return false
			}
		} else if direction.IsVertical() {
			var firstElementShifted = this.checkElementsShiftPossible(destination, direction)
			if !firstElementShifted {
				return false
			}

			var secondElementToShift Direction = destinationElement.OtherBigBoxPart()
			var secondElementShifted = this.shiftElementsIfPossible(destination.NextFor(secondElementToShift), direction)
			if !secondElementShifted {
				return false
			}

			this.shiftElementsIfPossible(destination, direction)
		}
	}

	this.setValueAt(destination, this.ElementAt(startCoordinate))
	this.setValueAt(startCoordinate, EMPTY)
	return true
}

func (this WarehouseMap) checkElementsShiftPossible(startCoordinate Coordinate, direction Direction) bool {
	var destination = startCoordinate.NextFor(direction)

	if destination.isOutOfBoundFor(this) {
		return false
	}

	var destinationElement = this.ElementAt(destination)
	if destinationElement == WALL {
		return false
	}

	if destinationElement.IsSmallBox() {
		var shifted = this.checkElementsShiftPossible(destination, direction)
		if !shifted {
			return false
		}
	}

	if destinationElement.IsBigBox() {
		if direction.IsHorizontal() {
			var shifted = this.checkElementsShiftPossible(destination, direction)
			if !shifted {
				return false
			}
		} else if direction.IsVertical() {
			var firstElementShifted = this.checkElementsShiftPossible(destination, direction)
			if !firstElementShifted {
				return false
			}

			var secondElementToShift Direction = destinationElement.OtherBigBoxPart()
			var secondElementShifted = this.checkElementsShiftPossible(destination.NextFor(secondElementToShift), direction)
			if !secondElementShifted {
				return false
			}

			this.checkElementsShiftPossible(destination, direction)
		}
	}

	return true
}

func (this WarehouseMap) GetWidth() int                         { return len(this[0]) }
func (this WarehouseMap) GetHeight() int                        { return len(this) }
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
			if mapElement == LBOX {
				sum += (100 * (y + 1)) + x + 2
			}
		}
	}

	return sum
}

func (this WarehouseMap) Clone() WarehouseMap {
	var deepCopy = make(WarehouseMap, len(this))

	for rowIndex, row := range this {
		deepCopy[rowIndex] = make([]MapElement, len(row))
		copy(deepCopy[rowIndex], row)
	}
	return deepCopy
}
