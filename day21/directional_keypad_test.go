package day21

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMovesToReachAPositionOnDirectionalKeypad(t *testing.T) {
	var testCases = []struct {
		startingPosition DirectionalKeypadButton
		positionToReach  DirectionalKeypadButton
		expectedMoves    []Move
	}{
		{startingPosition: ACTIVATE, positionToReach: UP, expectedMoves: []Move{LEFT}},
		{startingPosition: ACTIVATE, positionToReach: RIGHT, expectedMoves: []Move{DOWN}},
		{startingPosition: ACTIVATE, positionToReach: DOWN, expectedMoves: []Move{LEFT, DOWN}},
		{startingPosition: ACTIVATE, positionToReach: LEFT, expectedMoves: []Move{LEFT, DOWN, LEFT}},
		{startingPosition: ACTIVATE, positionToReach: ACTIVATE, expectedMoves: []Move{}},
		{startingPosition: DOWN, positionToReach: LEFT, expectedMoves: []Move{LEFT}},
		{startingPosition: DOWN, positionToReach: UP, expectedMoves: []Move{UP}},
		{startingPosition: DOWN, positionToReach: ACTIVATE, expectedMoves: []Move{UP, RIGHT}},
		{startingPosition: DOWN, positionToReach: DOWN, expectedMoves: []Move{}},
		{startingPosition: LEFT, positionToReach: DOWN, expectedMoves: []Move{RIGHT}},
		{startingPosition: LEFT, positionToReach: UP, expectedMoves: []Move{RIGHT, UP}},
		{startingPosition: LEFT, positionToReach: ACTIVATE, expectedMoves: []Move{RIGHT, UP, RIGHT}},
		{startingPosition: LEFT, positionToReach: RIGHT, expectedMoves: []Move{RIGHT, RIGHT}},
		{startingPosition: RIGHT, positionToReach: ACTIVATE, expectedMoves: []Move{UP}},
		{startingPosition: RIGHT, positionToReach: UP, expectedMoves: []Move{LEFT, UP}},
		{startingPosition: RIGHT, positionToReach: LEFT, expectedMoves: []Move{LEFT, LEFT}},
		{startingPosition: RIGHT, positionToReach: RIGHT, expectedMoves: []Move{}},
		{startingPosition: UP, positionToReach: ACTIVATE, expectedMoves: []Move{RIGHT}},
		{startingPosition: UP, positionToReach: DOWN, expectedMoves: []Move{DOWN}},
		{startingPosition: UP, positionToReach: LEFT, expectedMoves: []Move{DOWN, LEFT}},
		{startingPosition: UP, positionToReach: RIGHT, expectedMoves: []Move{DOWN, RIGHT}},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			var keypad = DirectionalKeypad{position: testCase.startingPosition}
			var moves = keypad.MovesToReach(testCase.positionToReach)
			assert.Equal(t, testCase.expectedMoves, moves)
		})
	}

}

func TestComposeSequenceOfMoves(t *testing.T) {
	var sequenceToCompose = []DirectionalKeypadButton{
		LEFT, ACTIVATE, UP, ACTIVATE, UP, UP, RIGHT, ACTIVATE, DOWN, DOWN, DOWN, ACTIVATE,
	}
	var keypad = DirectionalKeypad{position: ACTIVATE}

	var moves = keypad.ComposeSequence(sequenceToCompose)

	assert.Len(t, moves, 28)
}

func TestComposeLongerSequenceOfMoves(t *testing.T) {
	var sequenceToCompose = []DirectionalKeypadButton{
		DOWN, LEFT, LEFT, ACTIVATE, RIGHT, RIGHT, UP, ACTIVATE, LEFT, ACTIVATE,
		RIGHT, ACTIVATE, DOWN, ACTIVATE, LEFT, UP, ACTIVATE, ACTIVATE, RIGHT, ACTIVATE,
		LEFT, DOWN, ACTIVATE, ACTIVATE, ACTIVATE, RIGHT, UP, ACTIVATE,
	}
	var keypad = DirectionalKeypad{position: ACTIVATE}

	var moves = keypad.ComposeSequence(sequenceToCompose)

	assert.Len(t, moves, 68)
}

func TestDirectionalKeypadShouldPanicsWithErrorOnUnexpectedPosition(t *testing.T) {
	assert.PanicsWithError(t, "Unexpected keypad Position: 42", func() {
		var keypad = DirectionalKeypad{position: 42}
		keypad.MovesToReach(UP)
	})
}
