package day21

import (
	"fmt"
	"github.com/samber/lo"
)

type NumericKeypad struct {
	position NumericKeypadButton
}

func (this *NumericKeypad) ComposeCode(codeToCompose []NumericKeypadButton) [][]Move {
	if len(codeToCompose) == 0 {
		return [][]Move{}
	}

	var firstCode = codeToCompose[0]
	var collectionOfMoves = this.MovesToReach(firstCode)
	this.position = firstCode
	var restOfCodeToCompose = codeToCompose[1:]
	var collectionOfRestsMoves = this.ComposeCode(restOfCodeToCompose)

	return lo.FlatMap(collectionOfMoves, func(moves []Move, _ int) [][]Move {
		return lo.Map(collectionOfRestsMoves, func(rest []Move, _ int) []Move {
			return append(append(moves, ACTIVATE), rest...)
		})
	})
}

func (this NumericKeypad) MovesToReach(positionToReach NumericKeypadButton) [][]Move {
	return movesToReachAPositionOnNumericKeypad(this.position, positionToReach)
}

func movesToReachAPositionOnNumericKeypad(
	currentPosition NumericKeypadButton,
	positionToReach NumericKeypadButton,
) [][]Move {
	if currentPosition == positionToReach {
		return [][]Move{{}}
	}

	switch currentPosition {
	case ACT:
		switch positionToReach {
		case 0:
			return [][]Move{{LEFT}}
		case 3, 6, 9:
			return continueWith(UP, 3, positionToReach)
		default:
			return append(
				continueWith(LEFT, 0, positionToReach),
				continueWith(UP, 3, positionToReach)...,
			)
		}
	case 0:
		switch positionToReach {
		case ACT:
			return [][]Move{{RIGHT}}
		case 3, 6, 9:
			return append(
				continueWith(RIGHT, ACT, positionToReach),
				continueWith(UP, 2, positionToReach)...,
			)
		}
		return continueWith(UP, 2, positionToReach)
	case 1:
		switch positionToReach {
		case 4, 7:
			return continueWith(UP, 4, positionToReach)
		case 2, 3, 0, ACT:
			return continueWith(RIGHT, 2, positionToReach)
		default:
			return append(
				continueWith(RIGHT, 2, positionToReach),
				continueWith(UP, 4, positionToReach)...,
			)
		}
	case 2:
		switch positionToReach {
		case 0:
			return [][]Move{{DOWN}}
		case 1:
			return [][]Move{{LEFT}}
		case 3:
			return [][]Move{{RIGHT}}
		case 5:
			return [][]Move{{UP}}
		case 8:
			return continueWith(UP, 5, positionToReach)
		case 4, 7:
			return append(
				continueWith(LEFT, 1, positionToReach),
				continueWith(UP, 5, positionToReach)...,
			)
		case 6, 9:
			return append(
				continueWith(RIGHT, 3, positionToReach),
				continueWith(UP, 5, positionToReach)...,
			)
		case ACT:
			return append(
				continueWith(DOWN, 0, positionToReach),
				continueWith(RIGHT, 3, positionToReach)...,
			)
		}
	case 3:
		switch positionToReach {
		case ACT:
			return [][]Move{{DOWN}}
		case 6, 9:
			return continueWith(UP, 6, positionToReach)
		case 1, 2:
			return continueWith(LEFT, 2, positionToReach)
		case 0:
			return append(
				continueWith(LEFT, 2, positionToReach),
				continueWith(DOWN, ACT, positionToReach)...,
			)
		default:
			return append(
				continueWith(LEFT, 2, positionToReach),
				continueWith(UP, 6, positionToReach)...,
			)
		}
	case 4:
		switch positionToReach {
		case 1:
			return [][]Move{{DOWN}}
		case 7:
			return [][]Move{{UP}}
		case 5, 6:
			return continueWith(RIGHT, 5, positionToReach)
		case 8, 9:
			return append(
				continueWith(UP, 7, positionToReach),
				continueWith(RIGHT, 5, positionToReach)...,
			)
		default:
			return append(
				continueWith(RIGHT, 5, positionToReach),
				continueWith(DOWN, 1, positionToReach)...,
			)
		}
	case 5:
		switch positionToReach {
		case 1:
			return [][]Move{{LEFT, DOWN}, {DOWN, LEFT}}
		case 2:
			return [][]Move{{DOWN}}
		case 4:
			return [][]Move{{LEFT}}
		case 6:
			return [][]Move{{RIGHT}}
		case 7:
			return [][]Move{{LEFT, UP}, {UP, LEFT}}
		case 8:
			return [][]Move{{UP}}
		case 9:
			return [][]Move{{RIGHT, UP}, {UP, RIGHT}}
		case 0:
			return continueWith(DOWN, 2, positionToReach)
		case ACT, 3:
			return append(
				continueWith(RIGHT, 6, positionToReach),
				continueWith(DOWN, 2, positionToReach)...,
			)
		}
	case 6:
		switch positionToReach {
		case 3, ACT:
			return continueWith(DOWN, 3, positionToReach)
		case 4, 5:
			return continueWith(LEFT, 5, positionToReach)
		case 9:
			return [][]Move{{UP}}
		case 7, 8:
			return append(
				continueWith(UP, 9, positionToReach),
				continueWith(LEFT, 5, positionToReach)...,
			)
		case 1, 2:
			return append(
				continueWith(DOWN, 3, positionToReach),
				continueWith(LEFT, 5, positionToReach)...,
			)
		}
	case 7:
		switch positionToReach {
		case 1, 4:
			return continueWith(DOWN, 4, positionToReach)
		case 8, 9:
			return continueWith(RIGHT, 8, positionToReach)
		default:
			return append(
				continueWith(DOWN, 4, positionToReach),
				continueWith(RIGHT, 8, positionToReach)...,
			)
		}
	case 8:
		switch positionToReach {
		case 7:
			return [][]Move{{LEFT}}
		case 9:
			return [][]Move{{RIGHT}}
		case 5, 2, 0:
			return continueWith(DOWN, 5, positionToReach)
		case 1, 4:
			return append(
				continueWith(LEFT, 7, positionToReach),
				continueWith(DOWN, 5, positionToReach)...,
			)
		case 6, 3, ACT:
			return append(
				continueWith(RIGHT, 9, positionToReach),
				continueWith(DOWN, 5, positionToReach)...,
			)
		}
	case 9:
		switch positionToReach {
		case 7, 8:
			return continueWith(LEFT, 8, positionToReach)
		case ACT, 3, 6:
			return continueWith(DOWN, 6, positionToReach)
		default:
			return append(
				continueWith(LEFT, 8, positionToReach),
				continueWith(DOWN, 6, positionToReach)...,
			)
		}
	default:
		panic(fmt.Errorf("Unexpected keypad Position: %#v", currentPosition))
	}
	panic(fmt.Errorf(
		"Something went wrong from position %#v reaching %#v", currentPosition, positionToReach,
	))
}

func continueWith(move Move, newPosition NumericKeypadButton, positionToReach NumericKeypadButton) [][]Move {
	return lo.Map(movesToReachAPositionOnNumericKeypad(newPosition, positionToReach), func(tail []Move, i int) []Move {
		return append([]Move{move}, tail...)
	})

}

type NumericKeypadButton = int

const ACT NumericKeypadButton = -1
