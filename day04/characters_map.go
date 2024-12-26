package day04

import "github.com/samber/lo"

type CharactersMap [][]string

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
		&InverseHorizontalMapCursor{startingCoordinate},
		&VerticalMapCursor{startingCoordinate},
		&InverseVerticalMapCursor{startingCoordinate},
		&NorthDiagonalMapCursor{startingCoordinate},
		&SouthDiagonalMapCursor{startingCoordinate},
		&NorthInverseDiagonalMapCursor{startingCoordinate},
		&SouthInverseDiagonalMapCursor{startingCoordinate},
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

func (this CharactersMap) MasXAt(coordinate Coordinate) bool {
	var northOvest = coordinate.NorthOvest()
	var northEast = coordinate.NorthEast()
	var southOvest = coordinate.SouthOvest()
	var southEast = coordinate.SouthEast()

	var outOfBounds = lo.ContainsBy(
		[]Coordinate{coordinate, northOvest, northEast, southOvest, southEast},
		func(c Coordinate) bool { return this.IsOutOfBounds(c) },
	)

	if outOfBounds {
		return false
	}

	if this.At(coordinate) != "A" {
		return false
	}

	var first = this.At(northOvest) + this.At(southEast)
	var second = this.At(southOvest) + this.At(northEast)

	if first == "MS" && second == "MS" {
		return true
	}

	return false
}

func (this CharactersMap) IsOutOfBounds(c Coordinate) bool {
	// assuming always square shape maps
	return c.X < 0 || c.X >= len(this) || c.Y < 0 || c.Y >= len(this)
}

func (this CharactersMap) At(c Coordinate) string {
	return this[c.Y][c.X]
}
