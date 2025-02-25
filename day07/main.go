package day07

import (
	"github.com/samber/lo"
	"strings"
)

func TotalCalibrationResultOfPossiblyTrueEquations(fileContent string) int {
	var equations = parseEquationsFrom(fileContent)

	var trueEquations = lo.Filter(equations, func(equation Equation, _ int) bool {
		return CanBeTrue(equation, SUM_AND_PRODUCT)
	})

	return sumTotalsOf(trueEquations)
}

func TotalCalibrationResultOfPossiblyTrueEquationsWithConcatenationOperator(fileContent string) int {
	var equations = parseEquationsFrom(fileContent)

	var trueEquations = lo.Filter(equations, func(equation Equation, _ int) bool {
		return CanBeTrue(equation, ALL_OPERATOR)
	})

	return sumTotalsOf(trueEquations)
}

func parseEquationsFrom(fileContent string) []Equation {
	var inputLines = linesFrom(fileContent)
	return lo.Map(inputLines, func(line string, _ int) Equation {
		return ParseEquation(line)
	})
}

func sumTotalsOf(equations []Equation) int {
	return lo.SumBy(equations, func(equation Equation) int {
		return equation.total
	})
}

func linesFrom(input string) []string {
	var lines = strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	return lines
}
