package day03

import "strconv"

type Instruction interface {
	GetFirstOperand() int
	GetSecondOperand() int
	GetTotal() int
}

type MulInstruction struct {
	multiplying int
	multiplier  int
}

func ParseInstruction(s string) Instruction {
	var multiplying, _ = strconv.Atoi(s[4:5])
	var multiplier, _ = strconv.Atoi(s[6:7])
	return MulInstruction{multiplying, multiplier}
}

func (this MulInstruction) GetFirstOperand() int  { return this.multiplying }
func (this MulInstruction) GetSecondOperand() int { return this.multiplier }
func (this MulInstruction) GetTotal() int {
	return this.multiplying * this.multiplier
}
