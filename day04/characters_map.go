package day04

import "github.com/samber/lo"

type CharactersMap [][]string

type Coordinate struct {
	X int
	Y int
}

func (this CharactersMap) XMasOccurrencesAt(c Coordinate) int {
	if this.IsOutOfBounds(c) {
		return 0
	}

	if this[c.Y][c.X] != "X" {
		return 0
	}

	var neededParts = []string{"M", "A", "S"}
	var horizontalCursor MapCursor = &HorizontalMapCursor{c}
	var verticalCursor MapCursor = &VerticalMapCursor{c}

	for _, neededPart := range neededParts {

		if horizontalCursor != nil {
			horizontalCursor.Increase()
			if this.IsOutOfBounds(horizontalCursor.ToCoordinate()) || this.At(horizontalCursor.ToCoordinate()) != neededPart {
				horizontalCursor = nil
			}
		}

		if verticalCursor != nil {
			verticalCursor.Increase()
			if this.IsOutOfBounds(verticalCursor.ToCoordinate()) || this.At(verticalCursor.ToCoordinate()) != neededPart {
				verticalCursor = nil
			}
		}

	}

	var cursors = []MapCursor{horizontalCursor, verticalCursor}
	return lo.CountBy(cursors, func(c MapCursor) bool {
		return c != nil
	})
}

func (this CharactersMap) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X >= len(this) || c.Y < 0 || c.Y >= len(this)
}

func (this CharactersMap) At(c Coordinate) string {
	return this[c.Y][c.X]
}
