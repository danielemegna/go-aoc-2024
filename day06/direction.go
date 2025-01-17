package day06

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

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
