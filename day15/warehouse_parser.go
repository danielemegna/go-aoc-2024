package day15

import (
	"strings"
)

func parseWarehouse(input string) (*Warehouse, string) {
	lines := strings.Split(input, "\n")
	var mapLines []string
	var moves string
	parsingMap := true

	for _, line := range lines {
		if line == "" {
			parsingMap = false
			continue
		}
		if parsingMap {
			mapLines = append(mapLines, line)
		} else {
			moves += line
		}
	}

	height := len(mapLines)
	width := 0
	if height > 0 {
		width = len(mapLines[0])
	}

	walls := make(map[Coordinate]bool)
	boxes := make(map[Coordinate]bool)
	var robot Coordinate

	for y, line := range mapLines {
		for x, char := range line {
			coord := Coordinate{X: x, Y: y}
			switch char {
			case '#':
				walls[coord] = true
			case 'O':
				boxes[coord] = true
			case '@':
				robot = coord
			}
		}
	}

	return &Warehouse{
		Width:  width,
		Height: height,
		Walls:  walls,
		Boxes:  boxes,
		Robot:  robot,
	}, moves
}
