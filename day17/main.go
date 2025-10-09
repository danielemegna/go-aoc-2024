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
	var computer = ParseChronospatialComputer(inputContent)
	var expectedOutput = computer.GetInstructions()

	var registerAValue = 0
	for {
		computer.SetRegisterAValue(registerAValue)
		var isMatching = computer.RunProgramWithExpectedOutput(expectedOutput)
		if isMatching {
			return registerAValue
		}

		computer.ResetInstructionPointer()
		registerAValue++
	}
}

func LowestRegisterValueToPrintOutThe3BitConsumingLoopProgramItself(inputContent string) int {
	var computer = ParseChronospatialComputer(inputContent)
	var expectedOutput = computer.GetInstructions()
	var expectedOutputLength = len(expectedOutput)

	var registerAValue = 0
	var matchedOutputCount = 0

	for matchedOutputCount < expectedOutputLength {
		registerAValue = registerAValue << 3
		var expectedOutputTail = expectedOutput[expectedOutputLength-1-matchedOutputCount:]

		var valueAttempt = 0
		for {
			var registerAValueCandidate = registerAValue | valueAttempt
			computer.SetRegisterAValue(registerAValueCandidate)
			var isMatching = computer.RunProgramWithExpectedOutput(expectedOutputTail)
			if isMatching {
				registerAValue = registerAValueCandidate
				matchedOutputCount++
				break
			}

			computer.ResetInstructionPointer()
			valueAttempt++
		}

	}

	return registerAValue
}

func joinSliceOfIntToString(slice []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(slice), " "), ","), "[]")
}
