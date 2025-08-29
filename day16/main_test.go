package day16

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var FIRST_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"###############",
	"#.......#....E#",
	"#.#.###.#.###.#",
	"#.....#.#...#.#",
	"#.###.#####.#.#",
	"#.#.#.......#.#",
	"#.#.#####.###.#",
	"#...........#.#",
	"###.#.#####.#.#",
	"#...#.....#.#.#",
	"#.#.#.###.#.#.#",
	"#.....#...#.#.#",
	"#.###.#.#.#.#.#",
	"#S..#.....#...#",
	"###############",
}

var SECOND_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"#################",
	"#...#...#...#..E#",
	"#.#.#.#.#.#.#.#.#",
	"#.#.#.#...#...#.#",
	"#.#.#.#.###.#.#.#",
	"#...#.#.#.....#.#",
	"#.#.#.#.#.#####.#",
	"#.#...#.#.#.....#",
	"#.#.#####.#.###.#",
	"#.#.#.......#...#",
	"#.#.###.#####.###",
	"#.#.#...#.....#.#",
	"#.#.#.#####.###.#",
	"#.#.#.........#.#",
	"#.#.#.#########.#",
	"#S#.............#",
	"#################",
}

func TestSolveFirstPartWithFirstProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(FIRST_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 7036, LowestReindeerCostFor(fileContent))
}

func TestSolveFirstPartWithSecondProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SECOND_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 11048, LowestReindeerCostFor(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 134588, LowestReindeerCostFor(fileContent))
}

func simulateFileContent(inputLines []string) string {
	return strings.Join(inputLines, "\n") + "\n"
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
