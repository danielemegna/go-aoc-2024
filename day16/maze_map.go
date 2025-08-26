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
