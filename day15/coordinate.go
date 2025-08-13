package day15

import "fmt"

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
		return Coordinate{this.x, this.y + 1}
	}

	panic(fmt.Sprintf("Unexpected Direction value: %v", direction))
}

func (this Coordinate) isOutOfBound(mapSize int) bool {
	return this.x < 0 || this.y < 0 || this.x >= mapSize || this.y >= mapSize
}
