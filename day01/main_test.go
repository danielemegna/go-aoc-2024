package day01

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 11, TotalDistanceBetweenListsElements(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	var actual = TotalDistanceBetweenListsElements(fileContent)
	assert.Equal(t, 1530215, actual)
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 31, SimilarityScoreFor(fileContent))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	var actual = SimilarityScoreFor(fileContent)
	assert.Equal(t, 26800609, actual)
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
