package day08

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) IsOutOfBounds(mapSize int) bool {
	return c.X < 0 || c.X >= mapSize || c.Y < 0 || c.Y >= mapSize
}
