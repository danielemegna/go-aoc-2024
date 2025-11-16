package day15

import (
	"fmt"
	"strings"
)

type MapRowAppendFunction = func([]MapElement, MapElement) []MapElement

func ParseWarehouseMapAndMoves(inputFileContent string) (WarehouseMap, []Direction) {
	return parseWith(inputFileContent, singleWideMapRowAppend)
}

func ParseWarehouseMapAndMovesInDoubleWide(inputFileContent string) (WarehouseMap, []Direction) {
	return parseWith(inputFileContent, doubleWideMapRowAppend)
}

func parseWith(inputFileContent string, mapRowAppendFn MapRowAppendFunction) (WarehouseMap, []Direction) {
	var inputLines = linesFrom(inputFileContent)
	var warehouseMap = WarehouseMap{}
	var mapSize = len(inputLines[0]) // assuming always square non-empty map

	for lineIndex := 1; lineIndex < mapSize-1; lineIndex++ {
		var mapRow = []MapElement{}
		for charIndex := 1; charIndex < mapSize-1; charIndex++ {
			var char = rune(inputLines[lineIndex][charIndex])
			var mapElement = mapElementFrom(char)
			mapRow = mapRowAppendFn(mapRow, mapElement)
		}
		warehouseMap = append(warehouseMap, mapRow)
	}

	var moves = []Direction{}
	for lineIndex := mapSize + 1; lineIndex < len(inputLines); lineIndex++ {
		for charIndex := 0; charIndex < len(inputLines[lineIndex]); charIndex++ {
			var char = rune(inputLines[lineIndex][charIndex])
			moves = append(moves, directionFrom(char))
		}
	}

	return warehouseMap, moves
}

func directionFrom(char rune) Direction {
	switch char {
	case '>':
		return RIGHT
	case '<':
		return LEFT
	case 'v':
		return DOWN
	case '^':
		return UP
	}

	panic(fmt.Sprintf("Unexpected Direction value: %c", char))
}

func mapElementFrom(char rune) MapElement {
	switch char {
	case '.':
		return EMPTY
	case 'O':
		return BOX
	case '#':
		return WALL
	case '@':
		return ROBOT
	}

	panic(fmt.Sprintf("Unexpected MapElement value: %c", char))
}

func singleWideMapRowAppend(row []MapElement, e MapElement) []MapElement {
	return append(row, e)
}

func doubleWideMapRowAppend(row []MapElement, e MapElement) []MapElement {
	switch e {
	case BOX:
		return append(append(row, LBOX), RBOX)
	case EMPTY:
		return append(append(row, EMPTY), EMPTY)
	case WALL:
		return append(append(row, WALL), WALL)
	case ROBOT:
		return append(append(row, ROBOT), EMPTY)
	}

	panic(fmt.Sprintf("Unexpected MapElement: %#v", e))
}

func linesFrom(s string) []string {
	var rows = strings.Split(s, "\n")
	rows = rows[:len(rows)-1]
	return rows
}
