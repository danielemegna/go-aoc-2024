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

func TestParseMultiplyInstructionWithSingleDigitsOperand(t *testing.T) {
	var instruction = ParseInstruction("mul(2,3)")
	assert.Equal(t, MulInstruction{2, 3}, instruction)
	instruction = ParseInstruction("mul(5,2)")
	assert.Equal(t, MulInstruction{5, 2}, instruction)
}

func TestParseMultiplyInstructionWithMultipleDigitsOperand(t *testing.T) {
	var instruction = ParseInstruction("mul(11,4)")
	assert.Equal(t, MulInstruction{11, 4}, instruction)
	instruction = ParseInstruction("mul(2,24)")
	assert.Equal(t, MulInstruction{2, 24}, instruction)
}

func TestParseMultiplyInstructionWithGarbageAround(t *testing.T) {
	var instruction = ParseInstruction("%&xmul(14,8)!@^do")
	assert.Equal(t, MulInstruction{14, 8}, instruction)
}

func TestFindAndParseInstructionsOfEmptyString(t *testing.T) {
	var actual = FindAndParseInstructions("")
	assert.Equal(t, []Instruction{}, actual)
}

func TestFindAndParseInstructionsOfGarbageString(t *testing.T) {
	var actual = FindAndParseInstructions("4]then(mu")
	assert.Equal(t, []Instruction{}, actual)
}

func TestFindAndParseInstructionsSingleMulInstruction(t *testing.T) {
	var actual = FindAndParseInstructions("mul(11,8)")
	var expected = []Instruction{
		MulInstruction{11, 8},
	}
	assert.Equal(t, expected, actual)
}
