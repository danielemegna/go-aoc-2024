package day03

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 2*4+5*5+11*8+8*5, SumOfOperationsIn(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	t.Skip("WIP")
	var fileContent = readFileContent()
	var actual = SumOfOperationsIn(fileContent)
	assert.Equal(t, 999, actual)
}

func simulateFileContent(inputLines []string) string {
	return strings.Join(inputLines, "\n") + "\n"
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
