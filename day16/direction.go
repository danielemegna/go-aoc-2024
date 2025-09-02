package day16

import "fmt"

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (this Direction) Clockwise() Direction {
	switch this {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}

	panic(fmt.Sprintf("Unexpected Direction value: %v", this))
}

func (this Direction) CounterClockwise() Direction {
	switch this {
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	}

	panic(fmt.Sprintf("Unexpected Direction value: %v", this))
}

func (this Direction) Opposite() Direction {
	switch this {
	case UP:
		return DOWN
	case RIGHT:
		return LEFT
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	}

	panic(fmt.Sprintf("Unexpected Direction value: %v", this))
}
