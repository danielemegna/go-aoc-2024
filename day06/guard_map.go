package day06

import (
	"fmt"
	"slices"
	"strings"
)

type GuardMap struct {
	size      int // assuming always-square maps here
	guard     Guard
	obstacles []Coordinate
}

type Guard struct {
	position         Coordinate
	direction        Direction
	visitedPositions []Coordinate
}

type Coordinate struct {
	x int
	y int
}

func (this Coordinate) East() Coordinate  { return Coordinate{this.x + 1, this.y} }

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
				var guardPosition = Coordinate{x, y}
				foundGuard = &Guard{
					position:         guardPosition,
					direction:        guardDirecionFromChar(char),
					visitedPositions: []Coordinate{guardPosition},
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

func (this GuardMap) GuardWalk() GuardMap {
	var visitedPositions = this.guard.visitedPositions
	var currentPosition = this.guard.position

	for {
		var nextPosition = currentPosition.East()
		if slices.Contains(this.obstacles, nextPosition) {
			break
		}

		currentPosition = nextPosition
		if currentPosition.x >= this.size {
			break
		}
		visitedPositions = append(visitedPositions, nextPosition)
	}

	return GuardMap{
		size: this.size,
		guard: Guard{
			direction:        this.guard.direction,
			position:         currentPosition,
			visitedPositions: visitedPositions,
		},
		obstacles: this.obstacles,
	}
}

func guardDirecionFromChar(value rune) Direction {
	switch value {
	case '^':
		return North
	case 'v':
		return South
	case '>':
		return East
	case '<':
		return West
	}

	panic(fmt.Sprintf("Cannot recognize guard direction from %c", value))
}
