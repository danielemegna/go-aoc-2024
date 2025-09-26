package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunChronospatialComputerProgramWithSingleBstInstruction(t *testing.T) {
	var computer = NewChronospatialComputer(
		0, 0, 9, []int{2, 6},
	)

	computer.RunProgram()

	assert.Equal(t, 1, computer.RegisterValue('B'))
}

func TestRunChronospatialComputerProgramWithSingleBxlInstruction(t *testing.T) {
	var computer = NewChronospatialComputer(
		0, 29, 0, []int{1, 7},
	)

	computer.RunProgram()

	assert.Equal(t, 26, computer.RegisterValue('B'))
}

func TestRunChronospatialComputerProgramWithSingleBxcInstruction(t *testing.T) {
	var computer = NewChronospatialComputer(
		0, 2024, 43690, []int{4, 0},
	)

	computer.RunProgram()

	assert.Equal(t, 44354, computer.RegisterValue('B'))
}

func TestRunChronospatialComputerProgramWithThreeOutputInstructions(t *testing.T) {
	var computer = NewChronospatialComputer(
		10, 0, 0, []int{5, 0, 5, 1, 5, 4},
	)

	computer.RunProgram()

	assert.Equal(t, []int{0, 1, 2}, computer.GetOutput())
}

func TestRunChronospatialComputerProgramWithSomeDivisionOperations(t *testing.T) {
	var computer = NewChronospatialComputer(
		64, 0, 0, []int{0, 2, 6, 3, 7, 5},
	)

	computer.RunProgram()

	assert.Equal(t, 64/(2*2), computer.RegisterValue('A'))
	assert.Equal(t, 16/(2*2*2), computer.RegisterValue('B'))
	assert.Equal(t, 16/(2*2), computer.RegisterValue('C'))
}

func TestRunChronospatialComputerProgramWithMixedInstructions(t *testing.T) {
	var computer = NewChronospatialComputer(
		2024, 0, 0, []int{0, 1, 5, 4, 3, 0},
	)

	computer.RunProgram()

	assert.Equal(t, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, computer.GetOutput())
	assert.Equal(t, 0, computer.RegisterValue('A'))
}
