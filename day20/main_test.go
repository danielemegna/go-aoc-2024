package day20

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"###############",
	"#...#...#.....#",
	"#.#.#.#.#.###.#",
	"#S#...#.#.#...#",
	"#######.#.#.###",
	"#######.#.#...#",
	"#######.#.###.#",
	"###..E#...#...#",
	"###.#######.###",
	"#...###...#...#",
	"#.#####.#.###.#",
	"#.#...#.#.#...#",
	"#.#.#.#.#.#.###",
	"#...#...#...###",
	"###############",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 2 + 3 + 1 + 1 + 1 + 1 + 1, CountCheatsSavingAtLeast(10, fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 1355, CountCheatsSavingAtLeast(100, fileContent))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 285, CountLongCheatsSavingAtLeast(50, fileContent))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 1007335, CountLongCheatsSavingAtLeast(100, fileContent))
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
