package day21

import "fmt"

type DirectionalKeypad struct {
	position DirectionalKeypadButton
}

func (this *DirectionalKeypad) ComposeSequence(sequenceToCompose []DirectionalKeypadButton) []Move {
	if len(sequenceToCompose) == 0 {
		return []Move{}
	}
	var firstCode = sequenceToCompose[0]
	var moves = this.MovesToReach(firstCode)
	this.position = firstCode
	return append(append(moves, ACTIVATE), this.ComposeSequence(sequenceToCompose[1:])...)
}

func (this DirectionalKeypad) MovesToReach(positionToReach DirectionalKeypadButton) []Move {
	return movesToReachAPositionOnDirectionalKeypad(this.position, positionToReach)
}

func movesToReachAPositionOnDirectionalKeypad(
	currentPosition DirectionalKeypadButton,
	positionToReach DirectionalKeypadButton,
) []Move {
	if currentPosition == positionToReach {
		return []Move{}
	}

	switch currentPosition {

	case LEFT:
		return append([]Move{RIGHT}, movesToReachAPositionOnDirectionalKeypad(DOWN, positionToReach)...)

	case UP:
		if positionToReach == ACTIVATE {
			return []Move{RIGHT}
		}
		return append([]Move{DOWN}, movesToReachAPositionOnDirectionalKeypad(DOWN, positionToReach)...)

	case RIGHT:
		if positionToReach == ACTIVATE {
			return []Move{UP}
		}
		return append([]Move{LEFT}, movesToReachAPositionOnDirectionalKeypad(DOWN, positionToReach)...)

	case ACTIVATE:
		if positionToReach == RIGHT {
			return []Move{DOWN}
		}
		return append([]Move{LEFT}, movesToReachAPositionOnDirectionalKeypad(UP, positionToReach)...)

	case DOWN:
		switch positionToReach {
		case RIGHT:
			return []Move{RIGHT}
		case LEFT:
			return []Move{LEFT}
		default:
			return append([]Move{UP}, movesToReachAPositionOnDirectionalKeypad(UP, positionToReach)...)
		}

	default:
		panic(fmt.Errorf("Unexpected keypad Position: %#v", currentPosition))
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
