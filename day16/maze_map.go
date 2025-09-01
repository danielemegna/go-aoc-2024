package day16

type MazeMap [][]MapElement

type MapElement int

const (
	EMPTY MapElement = iota
	WALL
)

func (this MazeMap) GetWidth() int                     { return len(this[0]) }
func (this MazeMap) GetHeight() int                    { return len(this) }
func (this MazeMap) ElementAt(c Coordinate) MapElement { return this[c.Y][c.X] }

func (this MazeMap) isVisitable(c Coordinate) bool {
	return !this.IsOutOfBound(c) && this.ElementAt(c) != WALL
}

func (this MazeMap) IsOutOfBound(c Coordinate) bool {
	return c.X < 0 || c.Y < 0 || c.X >= this.GetWidth() || c.Y >= this.GetHeight()
}
