package day06

import (
	"fmt"
	"strings"
)

func ParseGuardMap(mapRows []string) GuardMap {
	var guard *Guard = nil
	var obstacles = []Coordinate{}

	for rowIndex, mapRow := range mapRows {
		if guard == nil {
			guard = findGuardIn(mapRow, rowIndex)
		}

	}

	if guard == nil {
		panic("Cannot find any guard in map!")
	}

	return GuardMap{
		size:      len(mapRows),
		guard:     *guard,
		obstacles: obstacles,
	}
}

func findGuardIn(mapRow string, rowIndex int) *Guard {
	var indexOfGuardInRow = strings.IndexAny(mapRow, "^><v")
	if indexOfGuardInRow == -1 {
		return nil
	}

	return &Guard{
		position:  Coordinate{x: indexOfGuardInRow, y: rowIndex},
		direction: guardDirecionFromChar(mapRow[indexOfGuardInRow]),
	}
}

func guardDirecionFromChar(value byte) Direction {
	switch value {
	case byte('^'):
		return North
	case byte('>'):
		return East
	case byte('<'):
		return West
	}

	panic(fmt.Sprintf("Cannot recognize guard direction from %c", value))
}

type GuardMap struct {
	size      int // assuming always-square maps here
	guard     Guard
	obstacles []Coordinate
}

type Guard struct {
	position  Coordinate
	direction Direction
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Coordinate struct {
	x int
	y int
}
