package day08

import (
	"testing"
)

func TestNewAntennasMapFrom(t *testing.T) {
	var input = `
..A..
.B...
..A..
`
	input = input[1:] // remove leading newline
	var am = NewAntennasMapFrom(input)

	if am.width != 5 {
		t.Errorf("Expected width 5, got %d", am.width)
	}
	if am.height != 3 {
		t.Errorf("Expected height 3, got %d", am.height)
	}

	var antennasA = am.antennasByFrequency['A']
	if len(antennasA) != 2 {
		t.Fatalf("Expected 2 antennas for A, got %d", len(antennasA))
	}
	if antennasA[0] != (Coordinate{X: 2, Y: 0}) {
		t.Errorf("Expected first A at (2, 0), got %v", antennasA[0])
	}
	if antennasA[1] != (Coordinate{X: 2, Y: 2}) {
		t.Errorf("Expected second A at (2, 2), got %v", antennasA[1])
	}

	var antennasB = am.antennasByFrequency['B']
	if len(antennasB) != 1 {
		t.Fatalf("Expected 1 antenna for B, got %d", len(antennasB))
	}
	if antennasB[0] != (Coordinate{X: 1, Y: 1}) {
		t.Errorf("Expected B at (1, 1), got %v", antennasB[0])
	}
}

func TestCountAntinodes(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Simple case",
			input: `
...#......
..........
....a.....
..........
.....a....
..........
......#...
`,
			expected: 2,
		},
		{
			name: "Provided example",
			input: `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`,
			expected: 14,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var input = tc.input[1:] // remove leading newline
			var am = NewAntennasMapFrom(input)
			var result = am.CountAntinodes()
			if result != tc.expected {
				t.Errorf("Expected %d antinodes, got %d", tc.expected, result)
			}
		})
	}
}
