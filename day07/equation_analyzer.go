package day07

import "strconv"

func CanBeTrue(equation Equation, operatorsMode OperatorsMode) bool {
	var firstOperand = equation.operands[0]
	var remainingOperands = equation.operands[1:]
	return CheckEquationTruthfulWith(remainingOperands, firstOperand, equation.total, operatorsMode)
}

func CheckEquationTruthfulWith(remainingOperands []int, total int, targetTotal int, operatorsMode OperatorsMode) bool {
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

	if operatorsMode == SUM_AND_PRODUCT {
		return CheckEquationTruthfulWith(remainingOperands, newTotalWithSum, targetTotal, operatorsMode) ||
			CheckEquationTruthfulWith(remainingOperands, newTotalWithProduct, targetTotal, operatorsMode)
	}

	// ALL_OPERATORS
	var newTotalWithConcatenation, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(firstNumber))
	return CheckEquationTruthfulWith(remainingOperands, newTotalWithSum, targetTotal, operatorsMode) ||
		CheckEquationTruthfulWith(remainingOperands, newTotalWithProduct, targetTotal, operatorsMode) ||
		CheckEquationTruthfulWith(remainingOperands, newTotalWithConcatenation, targetTotal, operatorsMode)
}

type OperatorsMode int

const (
	SUM_AND_PRODUCT OperatorsMode = iota
	ALL_OPERATOR
)
