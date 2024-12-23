package day04

type CharactersMap [][]string

type Coordinate struct {
	X int
	Y int
}

func (this CharactersMap) XMasOccurrencesAt(c Coordinate) int {
	if this.IsOutOfBounds(c.Y) || this.IsOutOfBounds(c.X) {
		return 0
	}

	if this[c.Y][c.X] != "X" {
		return 0
	}

	var neededParts = []string{"M", "A", "S"}
	for partIndex, neededPart := range neededParts {
		var checkingX = c.X + partIndex + 1
		if this.IsOutOfBounds(checkingX) {
			return 0
		}
		if this[c.Y][checkingX] != neededPart {
			return 0
		}
	}

	return 1
}

func (this CharactersMap) IsOutOfBounds(i int) bool {
	return i < 0 || i >= len(this)
}
