package day04

import (
	"sort"
	"strings"
)

type CharactersMap [][]string

func (this CharactersMap) XMasOccurrencesAt(startingCoordinate Coordinate) int {
	if this.IsOutOfBounds(startingCoordinate) {
		return 0
	}

	if this.At(startingCoordinate) != "X" {
		return 0
	}

	var cursors = initStarCursorsWith(startingCoordinate)
	var neededChars = []string{"M", "A", "S"}

	for _, neededChar := range neededChars {

		var cursorsCount = len(cursors)
		if cursorsCount == 0 {
			return 0
		}

		for i := 0; i < cursorsCount; i++ {
			var cursor = cursors[0]
			cursors = cursors[1:]

			cursor.Increase()
			if !this.IsOutOfBounds(cursor.coordinate) && this.At(cursor.coordinate) == neededChar {
				cursors = append(cursors, cursor)
			}
		}

	}

	return len(cursors)
}

func (this CharactersMap) HasMasXAt(coordinate Coordinate) bool {
	var northOvest = coordinate.NorthOvest()
	var northEast = coordinate.NorthEast()
	var southOvest = coordinate.SouthOvest()
	var southEast = coordinate.SouthEast()

	if this.AnyOutOfBounds(coordinate, northOvest, northEast, southOvest, southEast) {
		return false
	}

	if this.At(coordinate) != "A" {
		return false
	}

	var first = strings.Join(Sorted(this.At(northOvest), this.At(southEast)), "")
	var second = strings.Join(Sorted(this.At(southOvest), this.At(northEast)), "")

	if first == "MS" && second == "MS" {
		return true
	}

	return false
}


func initStarCursorsWith(startingCoordinate Coordinate) []MapCursor {
	var horizontalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.East() },
	}
	var inverseHorizontalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.West() },
	}
	var verticalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.South() },
	}
	var inverseVerticalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.North() },
	}
	var northDiagonalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.NorthEast() },
	}
	var southDiagonalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.SouthEast() },
	}
	var northInverseDiagonalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.NorthOvest() },
	}
	var southInverseDiagonalMapCursor = MapCursor{
		coordinate: startingCoordinate,
		increaseFn: func(c Coordinate) Coordinate { return c.SouthOvest() },
	}
	return []MapCursor{
		horizontalMapCursor,
		inverseHorizontalMapCursor,
		verticalMapCursor,
		inverseVerticalMapCursor,
		northDiagonalMapCursor,
		southDiagonalMapCursor,
		northInverseDiagonalMapCursor,
		southInverseDiagonalMapCursor,
	}
}

func (this CharactersMap) IsOutOfBounds(c Coordinate) bool {
	// assuming always square shape maps
	return c.X < 0 || c.X >= len(this) || c.Y < 0 || c.Y >= len(this)
}

func (this CharactersMap) AnyOutOfBounds(coordinates ...Coordinate) bool {
	for _, c := range coordinates {
		if this.IsOutOfBounds(c) {
			return true
		}
	}

	return false
}

func (this CharactersMap) At(c Coordinate) string {
	return this[c.Y][c.X]
}

func Sorted(source ...string) []string {
	var clone = make([]string, len(source))
	copy(clone, source)
	sort.Strings(clone)
	return clone
}
