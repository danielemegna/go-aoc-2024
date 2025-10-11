package day10

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

func (c Coordinate) IsOutOfBoundsOf(rawMap TopographicMap) bool {
	return c.X < 0 || c.Y < 0 || c.Y >= len(rawMap) || c.X >= len(rawMap[c.Y])
}
