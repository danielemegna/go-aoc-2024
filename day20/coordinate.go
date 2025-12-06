package day20

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) IsOutOfBounds(trackSize int) bool {
	return c.X < 0 || c.X >= trackSize || c.Y < 0 || c.Y >= trackSize
}
