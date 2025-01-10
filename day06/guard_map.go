package day06

import (
	"fmt"
	"strings"
)

type GuardMap struct {
	size      int // assuming always-square maps here
	guard     Guard
	obstacles []Coordinate
}

type Guard struct {
	position  Coordinate
	direction Direction
}

type Coordinate struct {
	x int
	y int
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func ParseGuardMap(mapRows []string) GuardMap {
	var mapSize = len(mapRows)
	var foundGuard *Guard = nil
	var foundObstacles = []Coordinate{}

	for y := 0; y < mapSize; y++ {
		for x := 0; x < mapSize; x++ {
			var char = rune(mapRows[y][x])

			if char == '#' {
				foundObstacles = append(foundObstacles, Coordinate{x, y})
				continue
			}

			if foundGuard != nil {
				continue
			}

			if strings.ContainsRune("^><v", char) {
				foundGuard = &Guard{
					position:  Coordinate{x, y},
					direction: guardDirecionFromChar(char),
				}
			}
		}
	}

	if foundGuard == nil {
		panic("Cannot find any guard in map!")
	}

	return GuardMap{
		size:      mapSize,
		guard:     *foundGuard,
		obstacles: foundObstacles,
	}
}

func guardDirecionFromChar(value rune) Direction {
	switch value {
	case '^':
		return North
	case '>':
		return East
	case '<':
		return West
	}

	panic(fmt.Sprintf("Cannot recognize guard direction from %c", value))
}
