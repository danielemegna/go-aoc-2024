package day21

import "fmt"

type DirectionalKeypad struct {
	position DirectionalKeypadButton
}

func (this DirectionalKeypad) MovesToReach(positionToReach DirectionalKeypadButton) []Move {
	switch this.position {
	case ACTIVATE:
		return movesToReachFromActivate(positionToReach)
	case UP:
		return movesToReachFromUp(positionToReach)
	case DOWN:
		return movesToReachFromDown(positionToReach)
	case LEFT:
		return movesToReachFromLeft(positionToReach)
	case RIGHT:
		return movesToReachFromRight(positionToReach)
	default:
		panic(fmt.Errorf("Unexpected keypad Position: %#v", this.position))
	}
}

func movesToReachFromUp(positionToReach DirectionalKeypadButton) []Move {
	switch positionToReach {
	case UP:
		return []Move{}
	case ACTIVATE:
		return []Move{RIGHT}
	default:
		return append([]Move{DOWN}, movesToReachFromDown(positionToReach)...)
	}
}

func movesToReachFromRight(positionToReach DirectionalKeypadButton) []Move {
	switch positionToReach {
	case RIGHT:
		return []Move{}
	case ACTIVATE:
		return []Move{UP}
	default:
		return append([]Move{LEFT}, movesToReachFromDown(positionToReach)...)
	}
}

func movesToReachFromActivate(positionToReach DirectionalKeypadButton) []Move {
	switch positionToReach {
	case ACTIVATE:
		return []Move{}
	case RIGHT:
		return []Move{DOWN}
	default:
		return append([]Move{LEFT}, movesToReachFromUp(positionToReach)...)
	}
}

func movesToReachFromDown(positionToReach DirectionalKeypadButton) []Move {
	switch positionToReach {
	case DOWN:
		return []Move{}
	case RIGHT:
		return []Move{RIGHT}
	case LEFT:
		return []Move{LEFT}
	default:
		return append([]Move{UP}, movesToReachFromUp(positionToReach)...)
	}
}

func movesToReachFromLeft(positionToReach DirectionalKeypadButton) []Move {
	switch positionToReach {
	case LEFT:
		return []Move{}
	default:
		return append([]Move{RIGHT}, movesToReachFromDown(positionToReach)...)
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
