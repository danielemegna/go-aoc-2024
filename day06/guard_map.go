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

func (this GuardMap) IsGuardOutOfBoundaries() bool {
	return true
}

type Guard struct {
	position         Coordinate
	direction        Direction
	visitedPositions []Coordinate
}

type Direction int

func (this Direction) NextClockwise() Direction {
	switch this {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		panic(fmt.Sprintf("Unexpected Direction: %#v", this))
	}
}

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
		var nextPosition = currentPosition.NextFor(this.guard.direction)
		if slices.Contains(this.obstacles, nextPosition) {
			break
		}

		currentPosition = nextPosition
		if currentPosition.x >= this.size || currentPosition.y < 0 {
			break
		}

		if !slices.Contains(visitedPositions, nextPosition) {
			visitedPositions = append(visitedPositions, nextPosition)
		}
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

func (this *GuardMap) TurnGuardClockwise() {
	this.guard.direction = this.guard.direction.NextClockwise()
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
