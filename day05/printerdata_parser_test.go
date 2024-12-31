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

	var expected = PageOrderingRules{
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

	assert.Equal(t, 21, len(actual))
	assert.Equal(t, PageOrderingRule{47, 53}, actual[0])
	assert.Equal(t, PageOrderingRule{97, 13}, actual[1])
	assert.Equal(t, PageOrderingRule{53, 29}, actual[9])
	assert.Equal(t, PageOrderingRule{53, 13}, actual[20])
}

func TestProvidedExamplePageUpdatesParsing(t *testing.T) {
	var _, actual = ParsePrinterData(PROVIDED_EXAMPLE_INPUT_LINES)

	assert.Equal(t, 6, len(actual))
	assert.Equal(t, PagesToProduceInTheUpdate{75, 47, 61, 53, 29}, actual[0])
	assert.Equal(t, PagesToProduceInTheUpdate{75, 29, 13}, actual[2])
	assert.Equal(t, PagesToProduceInTheUpdate{97, 13, 75, 29, 47}, actual[5])
}
