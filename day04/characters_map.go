package day04

type CharactersMap [][]string

type Coordinate struct {
	X int
	Y int
}

func (this CharactersMap) XMasOccurrencesAt(startingCoordinate Coordinate) int {
	if this.IsOutOfBounds(startingCoordinate) {
		return 0
	}

	if this.At(startingCoordinate) != "X" {
		return 0
	}

	var neededChars = []string{"M", "A", "S"}
	var cursors = []MapCursor{
		&HorizontalMapCursor{startingCoordinate},
		&VerticalMapCursor{startingCoordinate},
		&DiagonalMapCursor{startingCoordinate},
	}

	for _, neededChar := range neededChars {

		var cursorsCount = len(cursors)
		if cursorsCount == 0 {
			return 0
		}

		for i := 0; i < cursorsCount; i++ {
			var cursor = cursors[0]
			cursors = cursors[1:]

			cursor.Increase()
			if !this.IsOutOfBounds(cursor.ToCoordinate()) && this.At(cursor.ToCoordinate()) == neededChar {
				cursors = append(cursors, cursor)
			}
		}

	}

	return len(cursors)
}

func (this CharactersMap) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X >= len(this) || c.Y < 0 || c.Y >= len(this)
}

func (this CharactersMap) At(c Coordinate) string {
	return this[c.Y][c.X]
}
