package day17

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"Register A: 729",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,1,5,4,3,0",
}

var SECOND_PART_EXAMPLE_INPUT_LINES = []string{
	"Register A: 2024",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,3,5,4,3,0",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", ChronospatialComputerOutputFor(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, "7,5,4,3,4,5,3,4,6", ChronospatialComputerOutputFor(fileContent))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SECOND_PART_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 117440, LowestRegisterValueToPrintOutTheProgramItself(fileContent))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	assert.Equal(t, -1, LowestRegisterValueToPrintOutTheProgramItself(fileContent))
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
