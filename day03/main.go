package day03

import "github.com/samber/lo"

func SumOfInstructionsIn(fileContent string) int {
	var instructions = FindAndParseInstructions(fileContent)
	return lo.SumBy(instructions, func(i Instruction) int {
		return i.GetTotal()
	})
}
