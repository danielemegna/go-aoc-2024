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
	for this.instructionPointer < len(this.instructions) {
		var opcode = this.instructions[this.instructionPointer]
		var operand = this.instructions[this.instructionPointer+1]
		switch opcode {
		case 0:
			this.registerA = AdvOperation(this.registerA, this.comboOperand(operand))
		case 1:
			this.registerB = BitwiseXor(this.registerB, operand)
		case 2:
			this.registerB = this.comboOperandMod8(operand)
		case 3:
			if(this.registerA > 0) {
				this.instructionPointer = operand
				continue
			}
		case 4:
			this.registerB = BitwiseXor(this.registerB, this.registerC)
		case 5:
			this.output = append(this.output, this.comboOperandMod8(operand))
		}

		this.instructionPointer += 2
	}
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

func (this ChronospatialComputer) GetOutput() []int {
	return this.output
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
