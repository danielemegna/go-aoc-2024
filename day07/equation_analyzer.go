package day07

import "github.com/samber/lo"

func CanBeTrue(equation Equation) bool {
	var sum = lo.Sum(equation.operands)
	var product = lo.Reduce(equation.operands, func(product int, current int, index int) int {
		return product * current
	}, 1)
	return sum == equation.total || product == equation.total
}
