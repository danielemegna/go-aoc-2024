package day18

type MemoryState int

const (
	SAFE MemoryState = iota
	CORRUPTED
)

type MemorySpace [][]MemoryState

func (this MemorySpace) ShortestPathLengthFromTopLeftToBottomRightCorners() int {
	return 4
}

func (this MemorySpace) StateAt(c Coordinate) MemoryState {
	return this[c.Y][c.X]
}

func (this MemorySpace) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.Y < 0 || c.Y >= len(this) || c.X >= len(this[c.Y])
}

func (this MemorySpace) IsVisitable(c Coordinate) bool {
	return !this.IsOutOfBounds(c) && this.StateAt(c) == SAFE
}
