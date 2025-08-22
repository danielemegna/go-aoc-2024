package day15

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (d Direction) IsHorizontal() bool { return d == RIGHT || d == LEFT }
func (d Direction) IsVertical() bool   { return d == UP || d == DOWN }
