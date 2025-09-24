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
