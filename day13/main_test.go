package day13

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"Button A: X+94, Y+34",
	"Button B: X+22, Y+67",
	"Prize: X=8400, Y=5400",
	"",
	"Button A: X+26, Y+66",
	"Button B: X+67, Y+21",
	"Prize: X=12748, Y=12176",
	"",
	"Button A: X+17, Y+86",
	"Button B: X+84, Y+37",
	"Prize: X=7870, Y=6450",
	"",
	"Button A: X+69, Y+23",
	"Button B: X+27, Y+71",
	"Prize: X=18641, Y=10279",
}

func TestSolveFirstPartWithProvidedExample(t *testing.T) {
	var fileContent = SimulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 280+200, TotalTokensNeededToWinAllThePrizes(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 26810, TotalTokensNeededToWinAllThePrizes(fileContent))
}

func TestSolveSecondPartWithProvidedExample(t *testing.T) {
	var fileContent = SimulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 875318608908, RealTotalTokensNeededToWinAllThePrizes(fileContent))
}

func TestSolveSecondPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 108713182988244, RealTotalTokensNeededToWinAllThePrizes(fileContent))
}

func SimulateFileContent(inputLines []string) string {
	return strings.Join(inputLines, "\n") + "\n"
}

func readFileContent() string {
	var fileBytes, err = os.ReadFile("input.txt")
	if err != nil {
		panic("Cannot find input.txt file to solve with file")
	}
	return string(fileBytes)
}
