package day08

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 14, AntinodesInMap(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	// sorry but i am tired of 2x2 two dimensional problems
	// let me skip to day 09 and keep this for the future
	// ... a future where i'll like to solve 2D map problems
	t.Skip("WIP")
	var fileContent = readFileContent()
	assert.Equal(t, 5409, AntinodesInMap(fileContent))
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
