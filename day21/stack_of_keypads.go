package day21

import (
	"github.com/samber/lo"
	"math"
)

type StackOfKeypads struct {
	totalNumberOfKeypads int
}

func (this StackOfKeypads) LengthOfShortestSequenceToPressOnFirstKeypadFor(codeToCompose []NumericKeypadButton) int {
	var numericKeypad = NumericKeypad{position: ACT}

	var numericMoves = numericKeypad.ComposeCode(codeToCompose)
	var shortestSequenceLength = lo.Min(lo.Map(numericMoves, func(moves []Move, _ int) int {
		return len(moves)
	}))

	var sequencesTree = SequencesTree{
		head: Sequence{},
		tails: lo.Map(numericMoves, func(sequence Sequence, _ int) SequencesTree {
			return SequencesTree{head: sequence}
		}),
	}

	var remainingKeypads = this.totalNumberOfKeypads - 2
	for remainingKeypads > 0 {
		shortestSequenceLength = math.MaxInt

		var newSequenceTree SequencesTree
		for _, tailTree := range sequencesTree.tails {
			var innerTree, newLength = fooo(Sequence{}, tailTree, shortestSequenceLength)
			if newLength < shortestSequenceLength {
				shortestSequenceLength = newLength
				newSequenceTree = innerTree
			} else if newLength == shortestSequenceLength {
				newSequenceTree.AppendTail(innerTree)
			}
		}

		sequencesTree = newSequenceTree
		remainingKeypads--
	}

	return shortestSequenceLength
}

func fooo(partialSequence Sequence, tail SequencesTree, shortestSequenceSoFar int) (SequencesTree, int) {
	if len(tail.tails) == 0 {
		var sequenceToCompose = append(partialSequence, tail.head...)
		if len(sequenceToCompose) == 0 {
			return SequencesTree{head: Sequence{}}, shortestSequenceSoFar
		}
		var directionalKeypad = DirectionalKeypad{position: ACTIVATE}
		var innerSequencesTree = directionalKeypad.ComposeSequence(sequenceToCompose)
		var length = innerSequencesTree.ShortestSequenceLength()
		if length <= shortestSequenceSoFar {
			return innerSequencesTree, length
		}
		return SequencesTree{head: Sequence{}}, shortestSequenceSoFar
	}

	var newSequenceTree = SequencesTree{head: Sequence{}}
	for _, tailTree := range tail.tails {
		var newPartial = append(partialSequence, tail.head...)
		var innerTree, newLength = fooo(newPartial, tailTree, shortestSequenceSoFar)
		if newLength == shortestSequenceSoFar {
			newSequenceTree.AppendTail(innerTree)
			continue
		}

		if newLength > shortestSequenceSoFar {
			panic("This should never happen")
		}

		newSequenceTree = innerTree
		shortestSequenceSoFar = newLength
	}

	return newSequenceTree, shortestSequenceSoFar
}

func filterByLength(candidates [][]Move, desiredLength int) [][]Move {
	return lo.Filter(candidates, func(moves []Move, _ int) bool {
		return len(moves) == desiredLength
	})
}
