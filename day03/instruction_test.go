package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTotalOfMultiplyInstruction(t *testing.T) {
	var instruction = MulInstruction{2, 3}
	assert.Equal(t, 2*3, instruction.GetTotal())
	instruction = MulInstruction{3, 5}
	assert.Equal(t, 3*5, instruction.GetTotal())
}

func TestParseMultiplyInstruction(t *testing.T) {
	t.Skip("WIP")
	var instruction = ParseInstruction("mul(2,3)")
	assert.Equal(t, 2, instruction.GetFirstOperand())
	assert.Equal(t, 3, instruction.GetSecondOperand())
}
