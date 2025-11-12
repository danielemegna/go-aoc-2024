package day18

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"5,4",
	"4,2",
	"4,5",
	"3,0",
	"2,1",
	"6,3",
	"2,4",
	"1,5",
	"0,6",
	"3,3",
	"2,6",
	"5,1",
	"1,2",
	"5,5",
	"2,5",
	"6,5",
	"1,4",
	"0,4",
	"6,4",
	"1,1",
	"6,1",
	"1,0",
	"0,5",
	"1,6",
	"2,0",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 22, ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent, 7, 12))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 264, ShortestPathLengthFromTopLeftToBottomRightCorners(fileContent, 71, 1024))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, Coordinate{6, 1}, FirstByteMakesBottomRightCornerUnreachable(fileContent, 7))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, Coordinate{41, 26}, FirstByteMakesBottomRightCornerUnreachable(fileContent, 71))
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
