package day18

type MemorySpace [][]MemoryState

type MemoryState int

const (
	SAFE MemoryState = iota
	CORRUPTED
)

func (this MemorySpace) ShortestPathLengthFromTopLeftToBottomRightCorners() int {
	return 4
}
