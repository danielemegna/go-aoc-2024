package day06

import (
	"fmt"
	"strings"
)

func ParseGuardMap(mapRows []string) GuardMap {
	return GuardMap{
		size:      len(mapRows),
		guard:     findGuardIn(mapRows),
		obstacles: []Coordinate{},
	}
}

func findGuardIn(mapRows []string) Guard {
	for rowIndex, row := range mapRows {
		var indexOfGuardInRow = strings.IndexAny(row, "^><v")
		if indexOfGuardInRow == -1 {
			continue
		}

		return Guard{
			position:  Coordinate{x: indexOfGuardInRow, y: rowIndex},
			direction: guardDirecionFromChar(row[indexOfGuardInRow]),
		}
	}

	panic("Cannot find any guard")
}

func guardDirecionFromChar(value byte) Direction {
	switch value {
	case byte('^'):
		return North
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
