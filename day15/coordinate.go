package day15

type Coordinate struct {
	x int
	y int
}

func (this Coordinate) NextFor(direction Direction) Coordinate {
	switch direction {
	case RIGHT:
		return Coordinate{this.x + 1, this.y}
	case LEFT:
		return Coordinate{this.x - 1, this.y}
	case UP:
		return Coordinate{this.x, this.y - 1}
	case DOWN:
		fallthrough
	default:
		return Coordinate{this.x, this.y + 1}
	}
}

func (this Coordinate) isOutOfBound(mapWidth int, mapHeigth int) bool {
	return this.x < 0 || this.y < 0 || this.x >= mapWidth || this.y >= mapHeigth
}
