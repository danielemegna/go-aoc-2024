package day07

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 292 + 190 + 3267, TotalCalibrationResultOfPossiblyTrueEquations(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 6231007345478, TotalCalibrationResultOfPossiblyTrueEquations(fileContent))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	t.Skip("WIP")
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	var actual = TotalCalibrationResultOfPossiblyTrueEquationsWithConcatenationOperator(fileContent)
	assert.Equal(t, (292 + 190 + 3267 + 156 + 7290 + 192), actual)
}

func TestSolveSecondPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	var actual = TotalCalibrationResultOfPossiblyTrueEquationsWithConcatenationOperator(fileContent)
	assert.Equal(t, -1, actual)
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
