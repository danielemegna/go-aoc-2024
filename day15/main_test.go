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

var SECOND_PART_SMALLER_EXAMPLE_INPUT_LINES = []string{
	"#######",
	"#...#.#",
	"#.....#",
	"#..OO@#",
	"#..O..#",
	"#.....#",
	"#######",
	"",
	"<vv<<^^<<^^",
}

func TestSolveFirstPartWithSmallerProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SMALLER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 2028, BoxesGPSCoordinatesSumAfterAllRobotMoves(fileContent))
}

func TestSolveFirstPartWithLargerProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(LARGER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 10092, BoxesGPSCoordinatesSumAfterAllRobotMoves(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 1383666, BoxesGPSCoordinatesSumAfterAllRobotMoves(fileContent))
}

func TestSolveSecondPartWithSmallerProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SECOND_PART_SMALLER_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 618, BoxesGPSCoordinatesSumAfterAllRobotMovesInDoubleWide(fileContent))
}

func TestSolveSecondPartWithLargerProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(LARGER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 9021, BoxesGPSCoordinatesSumAfterAllRobotMovesInDoubleWide(fileContent))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	assert.Equal(t, -1, BoxesGPSCoordinatesSumAfterAllRobotMovesInDoubleWide(fileContent))
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
