package day20

import "math"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) CloseCoordinates() []Coordinate {
	return []Coordinate{
		{c.X - 1, c.Y},
		{c.X, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X, c.Y + 1},
	}
}

func (c Coordinate) IsOutOfBounds(trackSize int) bool {
	return c.X < 0 || c.X >= trackSize || c.Y < 0 || c.Y >= trackSize
}

func CalcManhattanDistance(from Coordinate, to Coordinate) int {
	return int(math.Abs(float64(to.X-from.X)) + math.Abs(float64(to.Y-from.Y)))
}
