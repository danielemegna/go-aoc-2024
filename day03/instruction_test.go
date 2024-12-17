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

func TestGetOperandsOfMultiplyInstruction(t *testing.T) {
	var instruction = MulInstruction{2, 3}
	assert.Equal(t, 2, instruction.GetFirstOperand())
	assert.Equal(t, 3, instruction.GetSecondOperand())
}

func TestParseMultiplyInstruction(t *testing.T) {
	var instruction = ParseInstruction("mul(2,3)")
	assert.Equal(t, 2, instruction.GetFirstOperand())
	assert.Equal(t, 3, instruction.GetSecondOperand())
	instruction = ParseInstruction("mul(5,2)")
	assert.Equal(t, 5, instruction.GetFirstOperand())
	assert.Equal(t, 2, instruction.GetSecondOperand())
}