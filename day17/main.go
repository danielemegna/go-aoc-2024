package day17

import (
	"fmt"
	"strings"
)

func ChronospatialComputerOutputFor(inputContent string) string {
	var computer = ParseChronospatialComputer(inputContent)
	computer.RunProgram()
	return joinSliceOfIntToString(computer.GetOutput())
}

func LowestRegisterValueToPrintOutTheProgramItself(inputContent string) int {
	var registerAValue = 1
	var computer = ParseChronospatialComputer(inputContent)
	var expectedOutput = computer.GetInstructions()
	for {
		computer.SetRegisterAValue(registerAValue)
		var isMatching = computer.RunProgramWithExpectedOutput(expectedOutput)
		if(isMatching) {
			return registerAValue
		}

		computer.ResetInstructionPointer()
		registerAValue++
	}
}

func joinSliceOfIntToString(slice []int) string {
    return strings.Trim(strings.Join(strings.Split(fmt.Sprint(slice), " "), ","), "[]")
}
