package day07

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Equation struct {
	operands []int
	total    int
}

func ParseEquation(s string) Equation {
	var regex, _ = regexp.Compile(`(\d+): (.+)`)
	var matches = regex.FindStringSubmatch(s)
	var total, _ = strconv.Atoi(matches[1])
	var operandsAsStrings = strings.Split(matches[2], " ")
	var operands = lo.Map(operandsAsStrings, func(operandAsString string, _ int) int {
		var operand, _ = strconv.Atoi(operandAsString)
		return operand
	})
	return Equation{operands: operands, total: total}
}
