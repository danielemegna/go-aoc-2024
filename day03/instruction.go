package day03

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
	return MulInstruction{}
}

func (this MulInstruction) GetFirstOperand() int  { return -1 }
func (this MulInstruction) GetSecondOperand() int { return -1 }
func (this MulInstruction) GetTotal() int {
	return this.multiplying * this.multiplier
}
