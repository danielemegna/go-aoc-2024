package day15

import "fmt"

type MapElement int

const (
	EMPTY MapElement = iota
	ROBOT
	WALL
	BOX
	LBOX
	RBOX
)

func (e MapElement) OtherBigBoxPart() Direction {
	switch e {
	case LBOX:
		return RIGHT
	case RBOX:
		return LEFT
	default:
		panic(fmt.Sprintf("Unexpected call to OtherBigBoxPart with %#v", e))
	}
}

func (e MapElement) IsSmallBox() bool { return e == BOX }
func (e MapElement) IsBigBox() bool   { return e == LBOX || e == RBOX }
func (e MapElement) IsBox() bool      { return e.IsSmallBox() || e.IsBigBox() }
