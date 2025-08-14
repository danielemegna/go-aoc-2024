package day15

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var SMALLER_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"########",
	"#..O.O.#",
	"##@.O..#",
	"#...O..#",
	"#.#.O..#",
	"#...O..#",
	"#......#",
	"########",
	"",
	"<^^>>>vv<v>>v<<",
}

var LARGER_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"##########",
	"#..O..O.O#",
	"#......O.#",
	"#.OO..O.O#",
	"#..O@..O.#",
	"#O#..O...#",
	"#O..O..O.#",
	"#.OO.O.OO#",
	"#....O...#",
	"##########",
	"",
	"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
	"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
	"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
	"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
	"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
	"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
	">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
	"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
	"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
	"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
}

func TestSolveFirstPartWithSmallerProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 2028, BoxesGPSCoordinatesSum(fileContent))
}

func TestSolveFirstPartWithLargerProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(LARGER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 10092, BoxesGPSCoordinatesSum(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 1383666, BoxesGPSCoordinatesSum(fileContent))
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
