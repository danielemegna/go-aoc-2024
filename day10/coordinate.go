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

type CoordinateSet map[Coordinate]bool

func (this CoordinateSet) Add(c Coordinate) {
	this[c] = true
}

func (this CoordinateSet) Merge(other CoordinateSet) {
	for c := range other {
		this[c] = true
	}
}

func (this CoordinateSet) ToSlice() []Coordinate {
	var result = []Coordinate{}
	for c := range this {
		result = append(result, c)
	}
	return result
}
