package day21

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposeCodeWithStackOfTwoKeypads(t *testing.T) {
	var stack = StackOfKeypads{totalNumberOfKeypads: 2}
	var codeToCompose = []NumericKeypadButton{0, 2, 9, ACT}

	var actual = stack.LengthOfShortestSequenceToPressOnFirstKeypadFor(codeToCompose)

	assert.Equal(t, 12, actual)
}

func TestComposeCodeWithStackOfThreeKeypads(t *testing.T) {
	var stack = StackOfKeypads{totalNumberOfKeypads: 3}
	var codeToCompose = []NumericKeypadButton{0, 2, 9, ACT}

	var actual = stack.LengthOfShortestSequenceToPressOnFirstKeypadFor(codeToCompose)

	assert.Equal(t, 28, actual)
}

func TestComposeCodeWithStackOfFourKeypads(t *testing.T) {
	var testCases = []struct {
		codeToCompose  []NumericKeypadButton
		expectedLength int
	}{
		{codeToCompose: []NumericKeypadButton{0, 2, 9, ACT}, expectedLength: 68},
		{codeToCompose: []NumericKeypadButton{9, 8, 0, ACT}, expectedLength: 60},
		{codeToCompose: []NumericKeypadButton{1, 7, 9, ACT}, expectedLength: 68},
		{codeToCompose: []NumericKeypadButton{4, 5, 6, ACT}, expectedLength: 64},
		{codeToCompose: []NumericKeypadButton{3, 7, 9, ACT}, expectedLength: 64},
	}

	var stack = StackOfKeypads{totalNumberOfKeypads: 4}
	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			var actual = stack.LengthOfShortestSequenceToPressOnFirstKeypadFor(testCase.codeToCompose)
			assert.Equal(t, testCase.expectedLength, actual)
		})
	}
}
