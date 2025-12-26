package day21

import (
	"github.com/samber/lo"
)

type StackOfKeypads struct {
	totalNumberOfKeypads int
}

func (this StackOfKeypads) LengthOfShortestSequenceToPressOnFirstKeypadFor(codeToCompose []NumericKeypadButton) int {
	var numericKeypad = NumericKeypad{position: ACT}
	var candidates = numericKeypad.ComposeCode(codeToCompose)

	var shortestSequenceLength = lo.Min(lo.Map(candidates, func(moves []Move, _ int) int {
		return len(moves)
	}))

	var remainingKeypads = this.totalNumberOfKeypads - 2
	for remainingKeypads > 0 {
		candidates = filterByLength(candidates, shortestSequenceLength)

		candidates = lo.FlatMap(candidates, func(moves []Move, _ int) [][]Move {
			var directionalKeypad = DirectionalKeypad{position: ACTIVATE}
			return directionalKeypad.ComposeSequence(moves)
		})

		shortestSequenceLength = lo.Min(lo.Map(candidates, func(moves []Move, _ int) int {
			return len(moves)
		}))

		remainingKeypads--
	}

	return shortestSequenceLength
}

func filterByLength(candidates [][]Move, desiredLength int) [][]Move {
	return lo.Filter(candidates, func(moves []Move, _ int) bool {
		return len(moves) == desiredLength
	})
}
