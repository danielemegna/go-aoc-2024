package day21

import "fmt"

type NumericKeypad struct {
	position NumericKeypadButton
}

func (this NumericKeypad) MovesToReach(positionToReach NumericKeypadButton) []Move {
	return movesToReachAPositionOnNumericKeypad(this.position, positionToReach)
}

func movesToReachAPositionOnNumericKeypad(
	currentPosition NumericKeypadButton,
	positionToReach NumericKeypadButton,
) []Move {
	if currentPosition == positionToReach {
		return []Move{}
	}

	switch currentPosition {
	case ACT:
		if positionToReach == 0 {
			return []Move{LEFT}
		}
		return append([]Move{UP}, movesToReachAPositionOnNumericKeypad(3, positionToReach)...)
	case 0:
		switch positionToReach {
		case ACT:
			return []Move{RIGHT}
		default:
			return append([]Move{UP}, movesToReachAPositionOnNumericKeypad(2, positionToReach)...)
		}
	case 2:
		switch positionToReach {
		case 1:
			return []Move{LEFT}
		case 3:
			return []Move{RIGHT}
		default:
			if positionToReach < 1 {
				return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(0, positionToReach)...)
			}
			return append([]Move{UP}, movesToReachAPositionOnNumericKeypad(5, positionToReach)...)
		}
	case 3:
		if positionToReach > 3 {
			return append([]Move{UP}, movesToReachAPositionOnNumericKeypad(6, positionToReach)...)
		}
		if positionToReach < 1 {
			return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(ACT, positionToReach)...)
		}
		return append([]Move{LEFT}, movesToReachAPositionOnNumericKeypad(2, positionToReach)...)
	case 4:
		switch positionToReach {
		case 1:
			return []Move{DOWN}
		case 7:
			return []Move{UP}
		default:
			return append([]Move{RIGHT}, movesToReachAPositionOnNumericKeypad(5, positionToReach)...)
		}
	case 5:
		switch positionToReach {
		case 4:
			return []Move{LEFT}
		case 6:
			return []Move{RIGHT}
		default:
			if positionToReach > 5 {
				return append([]Move{UP}, movesToReachAPositionOnNumericKeypad(8, positionToReach)...)
			}
			return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(2, positionToReach)...)
		}
	case 6:
		switch positionToReach {
		case 5:
			return []Move{LEFT}
		default:
			if positionToReach > 6 {
				return append([]Move{UP}, movesToReachAPositionOnNumericKeypad(9, positionToReach)...)
			}
			return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(3, positionToReach)...)
		}
	case 7:
		if positionToReach > 7 {
			return append([]Move{RIGHT}, movesToReachAPositionOnNumericKeypad(8, positionToReach)...)
		}
		return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(4, positionToReach)...)
	case 8:
		switch positionToReach {
		case 7:
			return []Move{LEFT}
		case 9:
			return []Move{RIGHT}
		default:
			return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(5, positionToReach)...)
		}
	case 9:
		if positionToReach > 6 {
			return append([]Move{LEFT}, movesToReachAPositionOnNumericKeypad(8, positionToReach)...)
		}
		return append([]Move{DOWN}, movesToReachAPositionOnNumericKeypad(6, positionToReach)...)

	default:
		panic(fmt.Errorf("Unexpected keypad Position: %#v", currentPosition))
	}
}

type NumericKeypadButton = int

const ACT NumericKeypadButton = -1
