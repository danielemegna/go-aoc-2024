package day11

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	assert.Equal(t, 55312, StonesCountAfterTwentyFiveBlinks("125 17"))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	assert.Equal(t, -1, StonesCountAfterTwentyFiveBlinks(fileContent))
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
