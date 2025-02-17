package day07

func CanBeTrue(equation Equation) bool {
	var firstOperand = equation.operands[0]
	var remainingOperands = equation.operands[1:]
	return CheckEquationTruthfulWith(remainingOperands, firstOperand, equation.total)
}

func CheckEquationTruthfulWith(remainingOperands []int, total int, targetTotal int) bool {
	if total > targetTotal {
		return false
	}

	if len(remainingOperands) == 0 {
		return total == targetTotal
	}

	var firstNumber = remainingOperands[0]
	remainingOperands = remainingOperands[1:]

	var newTotalWithSum = total + firstNumber
	var newTotalWithProduct = total * firstNumber

	return CheckEquationTruthfulWith(remainingOperands, newTotalWithSum, targetTotal) ||
		CheckEquationTruthfulWith(remainingOperands, newTotalWithProduct, targetTotal)
}
