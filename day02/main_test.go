package day02

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"7 6 4 2 1",
	"1 2 7 8 9",
	"9 7 6 2 1",
	"1 3 2 4 5",
	"8 6 4 4 1",
	"1 3 6 7 9",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 2, SafeReportsCount(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	var actual = SafeReportsCount(fileContent)
	assert.Equal(t, 9999, actual)
}

func simulateFileContent(inputLines []string) string {
	return strings.Join(inputLines, "\n") + "\n"
}

func readFileContent() string {
	var fileBytes, _ = os.ReadFile("input.txt")
	return string(fileBytes)
}