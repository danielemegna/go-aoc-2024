package day08

import (
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type AntennasMap struct {
	antennasByFrequency map[rune][]Coordinate
	width               int
	height              int
}

func NewAntennasMapFrom(input string) AntennasMap {
	var rows = strings.Split(input, "\n")
	if len(rows) > 0 && rows[len(rows)-1] == "" {
		rows = rows[:len(rows)-1]
	}

	var antennas = make(map[rune][]Coordinate)
	for y, row := range rows {
		for x, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], Coordinate{X: x, Y: y})
			}
		}
	}

	var width = 0
	if len(rows) > 0 {
		width = len(rows[0])
	}

	return AntennasMap{
		antennasByFrequency: antennas,
		width:               width,
		height:              len(rows),
	}
}

func (am AntennasMap) CountAntinodes() int {
	var antinodes = make(map[Coordinate]bool)

	for _, antennas := range am.antennasByFrequency {
		if len(antennas) < 2 {
			continue
		}

		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				var a = antennas[i]
				var b = antennas[j]

				// Antinode C = 2B - A
				var c1 = Coordinate{
					X: 2*b.X - a.X,
					Y: 2*b.Y - a.Y,
				}
				if am.isInBounds(c1) {
					antinodes[c1] = true
				}

				// Antinode C = 2A - B
				var c2 = Coordinate{
					X: 2*a.X - b.X,
					Y: 2*a.Y - b.Y,
				}
				if am.isInBounds(c2) {
					antinodes[c2] = true
				}
			}
		}
	}

	return len(antinodes)
}

func (am AntennasMap) isInBounds(c Coordinate) bool {
	return c.X >= 0 && c.X < am.width && c.Y >= 0 && c.Y < am.height
}

func SolveFirstPart(fileContent string) int {
	var am = NewAntennasMapFrom(fileContent)
	return am.CountAntinodes()
}
