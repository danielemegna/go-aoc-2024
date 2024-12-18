package day03

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	assert.Equal(t, 2*4+5*5+11*8+8*5, SumOfInstructionsIn(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	var actual = SumOfInstructionsIn(fileContent)
	assert.Equal(t, 174561379, actual)
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	var fileContent = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	assert.Equal(t, 2*4+8*5, SumOfEnabledInstructionsIn(fileContent))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	var actual = SumOfEnabledInstructionsIn(fileContent)
	assert.Equal(t, 106921067, actual)
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
