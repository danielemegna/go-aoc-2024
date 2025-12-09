package day20

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

func (c Coordinate) ManhattanDistanceTo(other Coordinate) int {
	dx := c.X - other.X
	if dx < 0 {
		dx = -dx
	}
	dy := c.Y - other.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}
