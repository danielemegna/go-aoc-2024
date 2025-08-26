package day16

import "fmt"

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (this Direction) Opposite() Direction {
	switch this {
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	case UP:
		return DOWN
	}

	panic(fmt.Sprintf("Unexpected Direction value: %v", this))
}
