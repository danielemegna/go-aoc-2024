package day14

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"p=0,4 v=3,-3",
	"p=6,3 v=-1,-3",
	"p=10,3 v=-1,2",
	"p=2,0 v=2,-1",
	"p=0,0 v=1,3",
	"p=3,0 v=-2,-2",
	"p=7,6 v=-1,-3",
	"p=3,0 v=-1,-2",
	"p=9,3 v=2,3",
	"p=7,3 v=-1,2",
	"p=2,4 v=2,-3",
	"p=9,5 v=-3,-3",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	var spaceSize = SpaceSize{width: 11, height: 7}
	assert.Equal(t, 1*3*4*1, SafetyFactorAfter100Seconds(fileContent, spaceSize))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	var spaceSize = SpaceSize{width: 101, height: 103}
	assert.Equal(t, 222901875, SafetyFactorAfter100Seconds(fileContent, spaceSize))
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
