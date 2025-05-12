package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSinglePlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("A\n")

	var expected = GardenMap{
		'A': {area: 1, perimeter: 4},
	}
	assert.Equal(t, expected, actual)
}

func TestTwoDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("AB\n")

	var expected = GardenMap{
		'A': {area: 1, perimeter: 4},
		'B': {area: 1, perimeter: 4},
	}
	assert.Equal(t, expected, actual)
}

func TestFounrDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap(
		"AB\n" +
		"CD\n",
	)

	var expected = GardenMap{
		'A': {area: 1, perimeter: 4},
		'B': {area: 1, perimeter: 4},
		'C': {area: 1, perimeter: 4},
		'D': {area: 1, perimeter: 4},
	}
	assert.Equal(t, expected, actual)
}

func TestParseSimpleProvidedExampleGardenMap(t *testing.T) {
	t.Skip("WIP")
	var actual = ParseGardenMap(simulateFileContent(SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES))

	var expected = GardenMap{
		'A': {area: 4, perimeter: 10},
		'B': {area: 4, perimeter: 8},
		'C': {area: 4, perimeter: 10},
		'D': {area: 1, perimeter: 4},
		'E': {area: 3, perimeter: 8},
	}
	assert.Equal(t, expected, actual)
}
