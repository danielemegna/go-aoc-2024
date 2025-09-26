package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseChronospatialComputerWithSimpleExample(t *testing.T) {
	var inputLines = []string{
		"Register A: 10",
		"Register B: 5",
		"Register C: 123",
		"",
		"Program: 2,6,4,1",
	}
	var computer = ParseChronospatialComputer(simulateFileContent(inputLines))

	assert.Equal(t, 10, computer.RegisterValue('A'))
	assert.Equal(t, 5, computer.RegisterValue('B'))
	assert.Equal(t, 123, computer.RegisterValue('C'))
	assert.Equal(t, []int{2, 6, 4, 1}, computer.instructions)
}

func TestParseChronospatialComputerWithProvidedExample(t *testing.T) {
	var computer = ParseChronospatialComputer(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 729, computer.RegisterValue('A'))
	assert.Equal(t, 0, computer.RegisterValue('B'))
	assert.Equal(t, 0, computer.RegisterValue('C'))
	assert.Equal(t, []int{0, 1, 5, 4, 3, 0}, computer.instructions)
}
