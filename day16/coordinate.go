package day16

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (this Coordinate) NextFor(direction Direction) Coordinate {
	switch direction {
	case RIGHT:
		return Coordinate{this.X + 1, this.Y}
	case LEFT:
		return Coordinate{this.X - 1, this.Y}
	case UP:
		return Coordinate{this.X, this.Y - 1}
	case DOWN:
		return Coordinate{this.X, this.Y + 1}
	}

	panic(fmt.Sprintf("Unexpected Direction value: %v", direction))
}

func (this Coordinate) IsOutOfBoundFor(m MazeMap) bool {
	return this.X < 0 || this.Y < 0 || this.X >= m.GetWidth() || this.Y >= m.GetHeight()
}

func (this Coordinate) CloseCoordinates() []Coordinate {
	return []Coordinate{
		{this.X + 1, this.Y},
		{this.X, this.Y + 1},
		{this.X - 1, this.Y},
		{this.X, this.Y - 1},
	}
}
