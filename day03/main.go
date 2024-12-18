package day03

import (
	"github.com/samber/lo"
	"strings"
)

func SumOfInstructionsIn(fileContent string) int {
	var instructions = FindAndParseInstructions(fileContent)
	return lo.SumBy(instructions, func(i Instruction) int {
		return i.GetTotal()
	})
}

func SumOfEnabledInstructionsIn(fileContent string) int {
	return SumOfInstructionsIn(keepOnlyEnabledPartsIn(fileContent))
}

func keepOnlyEnabledPartsIn(fileContent string) string {
	var enabledParts = ""

	var remaining = fileContent
	for true {
		var beforeDisable, afterDisable, foundDisableInstruction = strings.Cut(remaining, "don't()")
		enabledParts += beforeDisable

		if !foundDisableInstruction {
			break
		}

		var _, afterEnable, foundEnableInstruction = strings.Cut(afterDisable, "do()")
		if !foundEnableInstruction {
			break
		}

		remaining = afterEnable
	}
	return enabledParts
}
