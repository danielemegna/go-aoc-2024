package day04

type CharactersMap [][]string

type Coordinate struct {
	X int
	Y int
}

func (this CharactersMap) XMasOccurrencesAt(c Coordinate) int {
	if this[c.Y][c.X] != "X" {
		return 0
	}

	return 1
}
