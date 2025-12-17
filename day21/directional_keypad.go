package day21

import "fmt"

type DirectionalKeypad struct {
	position DirectionalKeypadButton
}

func (this DirectionalKeypad) MovesToReach(positionToReach DirectionalKeypadButton) []Move {
	return movesToReachFromPosition(this.position, positionToReach)
}

func movesToReachFromPosition(currentPosition DirectionalKeypadButton, positionToReach DirectionalKeypadButton) []Move {
	if currentPosition == positionToReach {
		return []Move{}
	}

	switch currentPosition {

	case LEFT:
		return append([]Move{RIGHT}, movesToReachFromPosition(DOWN, positionToReach)...)

	case UP:
		if positionToReach == ACTIVATE {
			return []Move{RIGHT}
		}
		return append([]Move{DOWN}, movesToReachFromPosition(DOWN, positionToReach)...)

	case RIGHT:
		if positionToReach == ACTIVATE {
			return []Move{UP}
		}
		return append([]Move{LEFT}, movesToReachFromPosition(DOWN, positionToReach)...)

	case ACTIVATE:
		if positionToReach == RIGHT {
			return []Move{DOWN}
		}
		return append([]Move{LEFT}, movesToReachFromPosition(UP, positionToReach)...)

	case DOWN:
		switch positionToReach {
		case RIGHT:
			return []Move{RIGHT}
		case LEFT:
			return []Move{LEFT}
		default:
			return append([]Move{UP}, movesToReachFromPosition(UP, positionToReach)...)
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
