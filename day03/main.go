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
	var remaining = fileContent
	var total = 0

	for true {
		var beforeDisable, afterDisable, _ = strings.Cut(remaining, "don't()")
		total += SumOfInstructionsIn(beforeDisable)

		var _, afterEnable, foundEnable = strings.Cut(afterDisable, "do()")
		if !foundEnable {
			break
		}

		remaining = afterEnable
	}
	return total
}
