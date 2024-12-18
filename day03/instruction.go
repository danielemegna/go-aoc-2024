package day03

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

func FindAndParseInstructions(s string) []Instruction {
	var r, _ = regexp.Compile(`mul\(\d+,\d+\)`)
	var matches = r.FindAllString(s, -1)

	return lo.Map(matches, func(m string, _ int) Instruction {
		return ParseInstruction(m)
	})
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
