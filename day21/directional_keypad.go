package day21

import (
	"fmt"
	"github.com/samber/lo"
)

type DirectionalKeypad struct {
	position DirectionalKeypadButton
}

// TODO: remove duplication with NumericKeypad#ComposeCode
func (this *DirectionalKeypad) ComposeSequence(sequenceToCompose []DirectionalKeypadButton) [][]Move {
	var firstCode = sequenceToCompose[0]
	var collectionOfMoves = this.MovesToReach(firstCode)
	collectionOfMoves = lo.Map(collectionOfMoves, func(moves []Move, _ int) []Move {
		return append(moves, ACTIVATE)
	})

	if len(sequenceToCompose) == 1 {
		return collectionOfMoves
	}

	this.position = firstCode
	var collectionOfRestsMoves = this.ComposeSequence(sequenceToCompose[1:])

	return lo.FlatMap(collectionOfMoves, func(moves []Move, _ int) [][]Move {
		return lo.Map(collectionOfRestsMoves, func(rest []Move, _ int) []Move {
			return append(moves, rest...)
		})
	})
}

func (this DirectionalKeypad) MovesToReach(positionToReach DirectionalKeypadButton) [][]Move {
	return movesToReachAPositionOnDirectionalKeypad(this.position, positionToReach)
}

func movesToReachAPositionOnDirectionalKeypad(
	currentPosition DirectionalKeypadButton,
	positionToReach DirectionalKeypadButton,
) [][]Move {
	if currentPosition == positionToReach {
		return [][]Move{{}}
	}

	switch currentPosition {

	case LEFT:
		return continueOnDirectionalKeypadWith(RIGHT, DOWN, positionToReach)

	case UP:
		switch positionToReach {
		case ACTIVATE:
			return [][]Move{{RIGHT}}
		case DOWN, LEFT:
			return continueOnDirectionalKeypadWith(DOWN, DOWN, positionToReach)
		case RIGHT:
			return append(
				continueOnDirectionalKeypadWith(DOWN, DOWN, positionToReach),
				continueOnDirectionalKeypadWith(RIGHT, ACTIVATE, positionToReach)...,
			)
		}

	case RIGHT:
		switch positionToReach {
		case ACTIVATE:
			return [][]Move{{UP}}
		case DOWN, LEFT:
			return continueOnDirectionalKeypadWith(LEFT, DOWN, positionToReach)
		case UP:
			return append(
				continueOnDirectionalKeypadWith(UP, ACTIVATE, positionToReach),
				continueOnDirectionalKeypadWith(LEFT, DOWN, positionToReach)...,
			)
		}

	case ACTIVATE:
		switch positionToReach {
		case UP:
			return [][]Move{{LEFT}}
		case RIGHT:
			return [][]Move{{DOWN}}
		case DOWN, LEFT:
			return append(
				continueOnDirectionalKeypadWith(DOWN, RIGHT, positionToReach),
				continueOnDirectionalKeypadWith(LEFT, UP, positionToReach)...,
			)
		}

	case DOWN:
		switch positionToReach {
		case RIGHT:
			return [][]Move{{RIGHT}}
		case LEFT:
			return [][]Move{{LEFT}}
		case UP:
			return [][]Move{{UP}}
		case ACTIVATE:
			return append(
				continueOnDirectionalKeypadWith(RIGHT, RIGHT, positionToReach),
				continueOnDirectionalKeypadWith(UP, UP, positionToReach)...,
			)
		}

	default:
		panic(fmt.Errorf("Unexpected keypad Position: %#v", currentPosition))

	}
	panic(fmt.Errorf(
		"Something went wrong from position %#v reaching %#v", currentPosition, positionToReach,
	))
}

func continueOnDirectionalKeypadWith(move Move, newPosition DirectionalKeypadButton, positionToReach DirectionalKeypadButton) [][]Move {
	return lo.Map(movesToReachAPositionOnDirectionalKeypad(newPosition, positionToReach), func(tail []Move, i int) []Move {
		return append([]Move{move}, tail...)
	})
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
