package day07

import "github.com/samber/lo"

func CanBeTrue(equation Equation) bool {
	var sum = lo.Sum(equation.operands)
	return sum == equation.total
}
