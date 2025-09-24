package day17

import "fmt"

type ChronospatialComputer struct {
	registerA          int
	registerB          int
	registerC          int
	instructions       []int
	instructionPointer int
	output             []int
}

func NewChronospatialComputer(
	registerA int,
	registerB int,
	registerC int,
	instructions []int,
) ChronospatialComputer {
	return ChronospatialComputer{
		registerA, registerB, registerC,
		instructions, 0, []int{},
	}
}

func (this *ChronospatialComputer) RunProgram() {
	var opcode = this.instructions[this.instructionPointer]
	var operand = this.instructions[this.instructionPointer+1]
	switch opcode {
	case 1:
		this.registerB = BitwiseXor(this.registerB, operand)
	case 2:
		this.registerB = this.comboOperandMod8(operand)
	}

	this.instructionPointer++
}

func (this ChronospatialComputer) RegisterValue(registerLabel rune) int {
	switch registerLabel {
	case 'A':
		return this.registerA
	case 'B':
		return this.registerB
	case 'C':
		return this.registerC
	}

	panic(fmt.Sprintf("Unexpected Register Value: %c", registerLabel))
}

func (this ChronospatialComputer) comboOperand(n int) int {
	if n < 4 {
		return n
	}

	switch n {
	case 4:
		return this.registerA
	case 5:
		return this.registerB
	case 6:
		return this.registerC
	}

	panic(fmt.Sprintf("Unexpected Combo Operand: %d", n))
}

func (this ChronospatialComputer) comboOperandMod8(n int) int {
	return this.comboOperand(n) % 8
}
