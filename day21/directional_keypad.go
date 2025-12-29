package day21

import (
	"fmt"
	"math"
)

type DirectionalKeypad struct {
	position DirectionalKeypadButton
}

type Sequence = []Move

type SequencesTree struct {
	head  Sequence
	tails []SequencesTree
}

func (this *SequencesTree) AppendTail(tree SequencesTree) {
	this.tails = append(this.tails, tree)
}

func (this *SequencesTree) AddActivateToTails() {
	if len(this.tails) == 0 {
		this.head = append(this.head, ACTIVATE)
		return
	}

	for i := 0; i < len(this.tails); i++ {
		var tail = this.tails[i]
		tail.AddActivateToTails()
		this.tails[i] = tail
	}
}

func (this *SequencesTree) AddTailToTails(tailTree SequencesTree) {
	if len(this.tails) == 0 {
		this.tails = append(this.tails, tailTree)
		return
	}

	for i := 0; i < len(this.tails); i++ {
		var tail = this.tails[i]
		tail.AddTailToTails(tailTree)
		this.tails[i] = tail
	}
}

func (this SequencesTree) ShortestSequenceLength() int {
	if len(this.tails) == 0 {
		return len(this.head)
	}

	var shortest = math.MaxInt
	for _, tail := range this.tails {
		var length = len(this.head) + tail.ShortestSequenceLength()
		if length < shortest {
			shortest = length
		}
	}
	return shortest
}

func (this *DirectionalKeypad) ComposeSequence(sequenceToCompose Sequence) SequencesTree {
	var firstCode = sequenceToCompose[0]
	var sequenceTree = this.MovesToReach(firstCode)
	sequenceTree.AddActivateToTails()

	if len(sequenceToCompose) == 1 {
		return sequenceTree
	}

	this.position = firstCode
	var tailTree = this.ComposeSequence(sequenceToCompose[1:])

	sequenceTree.AddTailToTails(tailTree)
	return sequenceTree
}

func (this DirectionalKeypad) MovesToReach(positionToReach DirectionalKeypadButton) SequencesTree {
	return movesToReachAPositionOnDirectionalKeypad(this.position, positionToReach)
}

func movesToReachAPositionOnDirectionalKeypad(
	currentPosition DirectionalKeypadButton,
	positionToReach DirectionalKeypadButton,
) SequencesTree {
	if currentPosition == positionToReach {
		return SequencesTree{head: Sequence{}}
	}

	switch currentPosition {

	case LEFT:
		return continueOnDirectionalKeypadWith(RIGHT, DOWN, positionToReach)

	case UP:
		switch positionToReach {
		case ACTIVATE:
			return SequencesTree{head: Sequence{RIGHT}}
		case DOWN, LEFT:
			return continueOnDirectionalKeypadWith(DOWN, DOWN, positionToReach)
		case RIGHT:
			return SequencesTree{
				head: Sequence{},
				tails: []SequencesTree{
					continueOnDirectionalKeypadWith(DOWN, DOWN, positionToReach),
					continueOnDirectionalKeypadWith(RIGHT, ACTIVATE, positionToReach),
				},
			}
		}

	case RIGHT:
		switch positionToReach {
		case ACTIVATE:
			return SequencesTree{head: Sequence{UP}}
		case DOWN, LEFT:
			return continueOnDirectionalKeypadWith(LEFT, DOWN, positionToReach)
		case UP:
			return SequencesTree{
				head: Sequence{},
				tails: []SequencesTree{
					continueOnDirectionalKeypadWith(UP, ACTIVATE, positionToReach),
					continueOnDirectionalKeypadWith(LEFT, DOWN, positionToReach),
				},
			}
		}

	case ACTIVATE:
		switch positionToReach {
		case UP:
			return SequencesTree{head: Sequence{LEFT}}
		case RIGHT:
			return SequencesTree{head: Sequence{DOWN}}
		case DOWN, LEFT:
			return SequencesTree{
				head: Sequence{},
				tails: []SequencesTree{
					continueOnDirectionalKeypadWith(DOWN, RIGHT, positionToReach),
					continueOnDirectionalKeypadWith(LEFT, UP, positionToReach),
				},
			}
		}

	case DOWN:
		switch positionToReach {
		case RIGHT:
			return SequencesTree{head: Sequence{RIGHT}}
		case LEFT:
			return SequencesTree{head: Sequence{LEFT}}
		case UP:
			return SequencesTree{head: Sequence{UP}}
		case ACTIVATE:
			return SequencesTree{
				head: Sequence{},
				tails: []SequencesTree{
					continueOnDirectionalKeypadWith(RIGHT, RIGHT, positionToReach),
					continueOnDirectionalKeypadWith(UP, UP, positionToReach),
				},
			}
		}

	default:
		panic(fmt.Errorf("Unexpected keypad Position: %#v", currentPosition))

	}
	panic(fmt.Errorf(
		"Something went wrong from position %#v reaching %#v", currentPosition, positionToReach,
	))
}

func continueOnDirectionalKeypadWith(move Move, newPosition DirectionalKeypadButton, positionToReach DirectionalKeypadButton) SequencesTree {
	var tailTree = movesToReachAPositionOnDirectionalKeypad(newPosition, positionToReach)
	return SequencesTree{
		head:  append(Sequence{move}, tailTree.head...),
		tails: tailTree.tails,
	}
}

type Move int

const (
	ACTIVATE Move = iota
	LEFT
	RIGHT
	DOWN
	UP
)

type DirectionalKeypadButton = Move
