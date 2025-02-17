package day07

import (
	"github.com/samber/lo"
	"strings"
)

func TotalCalibrationResultOfPossiblyTrueEquations(fileContent string) int {
	var inputLines = linesFrom(fileContent)
	var equations = lo.Map(inputLines, func(line string, _ int) Equation {
		return ParseEquation(line)
	})

	var trueEquations = lo.Filter(equations, func(equation Equation, _ int) bool {
		return CanBeTrue(equation)
	})

	return lo.SumBy(trueEquations, func(equation Equation) int {
		return equation.total
	})
}

func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
