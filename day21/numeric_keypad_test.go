package day21

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMovesToReachAPositionOnNumericKeypad(t *testing.T) {
	var testCases = []struct {
		startingPosition NumericKeypadButton
		positionToReach  NumericKeypadButton
		expectedMoves    []Move
	}{
		{startingPosition: ACT, positionToReach: 0, expectedMoves: []Move{LEFT}},
		{startingPosition: ACT, positionToReach: 3, expectedMoves: []Move{UP}},
		{startingPosition: ACT, positionToReach: 2, expectedMoves: []Move{UP, LEFT}},
		{startingPosition: ACT, positionToReach: 8, expectedMoves: []Move{UP, UP, UP, LEFT}},
		{startingPosition: ACT, positionToReach: 7, expectedMoves: []Move{UP, UP, UP, LEFT, LEFT}},
		{startingPosition: ACT, positionToReach: 9, expectedMoves: []Move{UP, UP, UP}},
		{startingPosition: 2, positionToReach: 5, expectedMoves: []Move{UP}},
		{startingPosition: 2, positionToReach: 9, expectedMoves: []Move{UP, UP, RIGHT}},
		{startingPosition: 2, positionToReach: 7, expectedMoves: []Move{UP, UP, LEFT}},
		{startingPosition: 2, positionToReach: 1, expectedMoves: []Move{LEFT}},
		{startingPosition: 2, positionToReach: 4, expectedMoves: []Move{UP, LEFT}},
		{startingPosition: 2, positionToReach: ACT, expectedMoves: []Move{DOWN, RIGHT}},
		{startingPosition: 3, positionToReach: 1, expectedMoves: []Move{LEFT, LEFT}},
		{startingPosition: 3, positionToReach: 3, expectedMoves: []Move{}},
		{startingPosition: 3, positionToReach: 2, expectedMoves: []Move{LEFT}},
		{startingPosition: 3, positionToReach: 4, expectedMoves: []Move{UP, LEFT, LEFT}},
		{startingPosition: 4, positionToReach: 2, expectedMoves: []Move{RIGHT, DOWN}},
		{startingPosition: 5, positionToReach: 4, expectedMoves: []Move{LEFT}},
		{startingPosition: 5, positionToReach: 6, expectedMoves: []Move{RIGHT}},
		{startingPosition: 5, positionToReach: 8, expectedMoves: []Move{UP}},
		{startingPosition: 5, positionToReach: 7, expectedMoves: []Move{UP, LEFT}},
		{startingPosition: 5, positionToReach: 3, expectedMoves: []Move{DOWN, RIGHT}},
		{startingPosition: 5, positionToReach: 9, expectedMoves: []Move{UP, RIGHT}},
		{startingPosition: 5, positionToReach: ACT, expectedMoves: []Move{DOWN, DOWN, RIGHT}},
		{startingPosition: 6, positionToReach: 4, expectedMoves: []Move{LEFT, LEFT}},
		{startingPosition: 7, positionToReach: 1, expectedMoves: []Move{DOWN, DOWN}},
		{startingPosition: 7, positionToReach: 5, expectedMoves: []Move{DOWN, RIGHT}},
		{startingPosition: 7, positionToReach: 6, expectedMoves: []Move{DOWN, RIGHT, RIGHT}},
		{startingPosition: 7, positionToReach: ACT, expectedMoves: []Move{DOWN, RIGHT, DOWN, DOWN, RIGHT}},
		{startingPosition: 9, positionToReach: ACT, expectedMoves: []Move{DOWN, DOWN, DOWN}},
		{startingPosition: 9, positionToReach: 7, expectedMoves: []Move{LEFT, LEFT}},
		{startingPosition: 9, positionToReach: 2, expectedMoves: []Move{DOWN, DOWN, LEFT}},
		{startingPosition: 9, positionToReach: 1, expectedMoves: []Move{DOWN, DOWN, LEFT, LEFT}},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			var keypad = NumericKeypad{position: testCase.startingPosition}
			var collectionOfPossibleMoves = keypad.MovesToReach(testCase.positionToReach)
			assert.Contains(t, collectionOfPossibleMoves, testCase.expectedMoves)
			// assert every moves slice has same expected size
		})
	}

}

func TestComposeCodes(t *testing.T) {
	t.Skip("WIP")
	var codeToCompose = []NumericKeypadButton{0, 2, 9, ACT}
	var numericKeypad = NumericKeypad{position: ACT}

	var collectionOfPossibleMoves = numericKeypad.ComposeCode(codeToCompose)

	assert.Contains(t,
		collectionOfPossibleMoves,
		[]Move{LEFT, ACTIVATE, UP, ACTIVATE, UP, UP, RIGHT, ACTIVATE, DOWN, DOWN, DOWN, ACTIVATE},
	)
}

func TestNumericKeypadShouldPanicsWithErrorOnUnexpectedPosition(t *testing.T) {
	assert.PanicsWithError(t, "Unexpected keypad Position: 42", func() {
		var keypad = NumericKeypad{position: 42}
		keypad.MovesToReach(5)
	})
}
