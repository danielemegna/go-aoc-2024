package day03

import (
	"strconv"
	"strings"
)

func FindAndParseInstructions(s string) []Instruction {
	if strings.HasPrefix(s, "mul") {
		return []Instruction{ParseInstruction(s)}
	}
	return []Instruction{}
}

func ParseInstruction(s string) Instruction {
	var firstOperandPart, secondOperandPart, _ = strings.Cut(s, ",")
	var _, firstOperandString, _ = strings.Cut(firstOperandPart, "(")
	var secondOperandString, _, _ = strings.Cut(secondOperandPart, ")")
	var multiplying, _ = strconv.Atoi(firstOperandString)
	var multiplier, _ = strconv.Atoi(secondOperandString)
	return MulInstruction{multiplying, multiplier}
}

type Instruction interface {
	GetTotal() int
}

type MulInstruction struct {
	multiplying int
	multiplier  int
}

func (this MulInstruction) GetTotal() int {
	return this.multiplying * this.multiplier
}
