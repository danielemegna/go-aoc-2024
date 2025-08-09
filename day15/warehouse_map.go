package day15

type MapElement int

const (
	EMPTY MapElement = iota
	BOX
	ROBOT
	WALL
)

type WarehouseMap [][]MapElement

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
