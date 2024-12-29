package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var SIMPLE_EXAMPLE_INPUT_LINES = []string{
	"47|53",
	"53|13",
	"",
	"75,47,61,53,29",
	"61,13,29",
}

func TestSimpleProvidedExamplePageOrderingRules(t *testing.T) {
	var actual, _ = ParsePrinterData(SIMPLE_EXAMPLE_INPUT_LINES)

	var expected = []PageOrderingRule{
		{before: 47, after: 53},
		{before: 53, after: 13},
	}
	assert.Equal(t, expected, actual)
}

func TestSimpleExamplePageUpdatesParsing(t *testing.T) {
	var _, actual = ParsePrinterData(SIMPLE_EXAMPLE_INPUT_LINES)

	var expected = []PagesToProduceInTheUpdate{
		{75, 47, 61, 53, 29},
		{61, 13, 29},
	}
	assert.Equal(t, expected, actual)
}

func TestProvidedExamplePageOrderingRules(t *testing.T) {
	var actual, _ = ParsePrinterData(PROVIDED_EXAMPLE_INPUT_LINES)
	var expected = []PageOrderingRule{}

	t.Skip("WIP")
	assert.Equal(t, expected, actual)
}

func TestProvidedExamplePageUpdatesParsing(t *testing.T) {
	var _, actual = ParsePrinterData(PROVIDED_EXAMPLE_INPUT_LINES)
	var expected = []PagesToProduceInTheUpdate{}

	t.Skip("WIP")
	assert.Equal(t, expected, actual)
}
