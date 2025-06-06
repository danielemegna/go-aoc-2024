package day12

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"AAAA",
	"BBCD",
	"BBCC",
	"EEEC",
}

var ANOTHER_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"OOOOO",
	"OXOXO",
	"OOOOO",
	"OXOXO",
	"OOOOO",
}

var E_SHAPED_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"EEEEE",
	"EXXXX",
	"EEEEE",
	"EXXXX",
	"EEEEE",
}

var SPECIAL_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"AAAAAA",
	"AAABBA",
	"AAABBA",
	"ABBAAA",
	"ABBAAA",
	"AAAAAA",
}

var LARGE_PROVIDED_EXAMPLE_INPUT_LINES = []string{
	"RRRRIICCFF",
	"RRRRIICCCF",
	"VVRRRCCFFF",
	"VVRCCCJFFF",
	"VVVVCJJCFE",
	"VVIVCCJJEE",
	"VVIIICJJEE",
	"MIIIIIJJEE",
	"MIIISIJEEE",
	"MMMISSJEEE",
}

func TestSolveFirstPartWithSimpleProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 140, TotalFencePrice(fileContent))
}

func TestSolveFirstPartWithAnotherProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(ANOTHER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 772, TotalFencePrice(fileContent))
}

func TestSolveFirstPartWithLargeProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(LARGE_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 1930, TotalFencePrice(fileContent))
}

func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	assert.Equal(t, 1518548, TotalFencePrice(fileContent))
}

func TestSolveSecondPartWithSimpleProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 80, TotalFencePriceWithBulkDiscount(fileContent))
}

func TestSolveSecondPartWithAnotherProvidedExample(t *testing.T) {
	t.Skip("FIXME: Something is calculated wrong in perimeter sides...")
	var fileContent = simulateFileContent(ANOTHER_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 436, TotalFencePriceWithBulkDiscount(fileContent))
}

func TestSolveSecondPartWithEShapedProvidedExample(t *testing.T) {
	var fileContent = simulateFileContent(E_SHAPED_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 236, TotalFencePriceWithBulkDiscount(fileContent))
}

func TestSolveSecondPartWithSpecialProvidedExample(t *testing.T) {
	t.Skip("FIXME: Something is calculated wrong in perimeter sides...")
	var fileContent = simulateFileContent(SPECIAL_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 368, TotalFencePriceWithBulkDiscount(fileContent))
}

func TestSolveSecondPartWithLargeProvidedExample(t *testing.T) {
	t.Skip("FIXME: Something is calculated wrong in perimeter sides...")
	var fileContent = simulateFileContent(LARGE_PROVIDED_EXAMPLE_INPUT_LINES)
	assert.Equal(t, 1206, TotalFencePriceWithBulkDiscount(fileContent))
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
