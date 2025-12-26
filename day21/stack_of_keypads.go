package day21

import (
	"github.com/samber/lo"
)

type StackOfKeypads struct {
	totalNumberOfKeypads int
}

func (this StackOfKeypads) LengthOfShortestSequenceToPressOnFirstKeypadFor(codeToCompose []NumericKeypadButton) int {
	var numericKeypad = NumericKeypad{position: ACT}
	var possibleMovesToComposeCode = numericKeypad.ComposeCode(codeToCompose)

	var shortestSequenceLength = lo.Min(lo.Map(possibleMovesToComposeCode, func(moves []Move, _ int) int {
		return len(moves)
	}))

	if this.totalNumberOfKeypads == 2 {
		return shortestSequenceLength
	}
	var onlyShortSequences = lo.Filter(possibleMovesToComposeCode, func(moves []Move, _ int) bool {
		return len(moves) == shortestSequenceLength
	})

	var possibleMovesToComposeSequence = lo.FlatMap(onlyShortSequences, func(moves []Move, _ int) [][]Move {
		var directionalKeypad = DirectionalKeypad{position: ACTIVATE}
		return directionalKeypad.ComposeSequence(moves)
	})

	shortestSequenceLength = lo.Min(lo.Map(possibleMovesToComposeSequence, func(moves []Move, _ int) int {
		return len(moves)
	}))

	if this.totalNumberOfKeypads == 3 {
		return shortestSequenceLength
	}

	onlyShortSequences = lo.Filter(possibleMovesToComposeSequence, func(moves []Move, _ int) bool {
		return len(moves) == shortestSequenceLength
	})

	possibleMovesToComposeSequence = lo.FlatMap(onlyShortSequences, func(moves []Move, _ int) [][]Move {
		var directionalKeypad = DirectionalKeypad{position: ACTIVATE}
		return directionalKeypad.ComposeSequence(moves)
	})

	return lo.Min(lo.Map(possibleMovesToComposeSequence, func(moves []Move, _ int) int {
		return len(moves)
	}))
}
