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
