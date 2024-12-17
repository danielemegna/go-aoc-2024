package day03

import (
	"strconv"
	"strings"
)

func ParseInstruction(s string) Instruction {
	var commaSplitParts = strings.Split(s, ",")
	var firstOperandPart = strings.Split(commaSplitParts[0], "(")
	var secondOperandParts = strings.Split(commaSplitParts[1], ")")
	var multiplying, _ = strconv.Atoi(firstOperandPart[len(firstOperandPart)-1])
	var multiplier, _ = strconv.Atoi(secondOperandParts[0])
	return MulInstruction{multiplying, multiplier}
}

type Instruction interface {
	GetFirstOperand() int
	GetSecondOperand() int
	GetTotal() int
}

type MulInstruction struct {
	multiplying int
	multiplier  int
}

func (this MulInstruction) GetFirstOperand() int  { return this.multiplying }
func (this MulInstruction) GetSecondOperand() int { return this.multiplier }
func (this MulInstruction) GetTotal() int {
	return this.multiplying * this.multiplier
}
