package day18

import (
	"regexp"
	"strconv"
	"strings"
)

type MemoryState int

const (
	SAFE MemoryState = iota
	CORRUPTED
)

type MemorySpace [][]MemoryState

func BuildMemorySpaceFrom(fileContent string, memorySpaceSize int) MemorySpace {
	var memorySpace = initSafeMemorySpace(memorySpaceSize)
	var lines = linesFrom(fileContent)
	for _, line := range lines {
		var r, _ = regexp.Compile(`(\d+),(\d+)`)
		var matches = r.FindStringSubmatch(line)
		var X, _ = strconv.Atoi(matches[1])
		var Y, _ = strconv.Atoi(matches[2])
		memorySpace[Y][X] = CORRUPTED
	}

	return memorySpace
}

func initSafeMemorySpace(size int) MemorySpace {
	var memorySpace = make(MemorySpace, size)
	for y := 0; y < size; y++ {
		memorySpace[y] = make([]MemoryState, size)
	}
	return memorySpace
}

func (this MemorySpace) StateAt(c Coordinate) MemoryState {
	return this[c.Y][c.X]
}

func (this MemorySpace) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.Y < 0 || c.Y >= this.Height() || c.X >= this.Width()
}

func (this MemorySpace) IsVisitable(c Coordinate) bool {
	return !this.IsOutOfBounds(c) && this.StateAt(c) == SAFE
}

func (this MemorySpace) Width() int {
	return len(this[0])
}

func (this MemorySpace) Height() int {
	return len(this)
}

func (this MemorySpace) BottomRightCoordinate() Coordinate {
	return Coordinate{this.Width() - 1, this.Height() - 1}
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
