package day15

import "fmt"

func ParseWarehouseMapAndMoves(inputLines []string) (WarehouseMap, []Direction) {
	var warehouseMap = WarehouseMap{}
	var mapSize = len(inputLines[0]) // assuming always square non-empty map

	for lineIndex := 1; lineIndex < mapSize-1; lineIndex++ {
		var mapRow = []MapElement{}
		for charIndex := 1; charIndex < mapSize-1; charIndex++ {
			var char = rune(inputLines[lineIndex][charIndex])
			mapRow = append(mapRow, mapElementFrom(char))
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
