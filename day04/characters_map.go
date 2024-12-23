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
	var horizontalCursor = c.X + 1
	var verticalCursor = c.Y + 1

	for _, neededPart := range neededParts {
		if this.IsOutOfBounds(horizontalCursor) {
			horizontalCursor = -1
		} else {
			if this[c.Y][horizontalCursor] != neededPart {
				horizontalCursor = -1
			}
		}
		if this.IsOutOfBounds(verticalCursor) {
			verticalCursor = -1
		} else {
			if this[verticalCursor][c.X] != neededPart {
				verticalCursor = -1
			}
		}

		horizontalCursor++
		verticalCursor++

	}

	if horizontalCursor > 0 || verticalCursor > 0 {
		return 1
	}

	return 0
}

func (this CharactersMap) IsOutOfBounds(i int) bool {
	return i < 0 || i >= len(this)
}
