package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSinglePlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("A\n")

	var expected = GardenMap{
		{plant: 'A', area: 1, perimeter: 4},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestTwoDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("AB\n")

	var expected = GardenMap{
		{plant: 'A', area: 1, perimeter: 4},
		{plant: 'B', area: 1, perimeter: 4},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestFourDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap(
		"AB\n" +
		"CD\n",
	)

	var expected = GardenMap{
		{plant: 'A', area: 1, perimeter: 4},
		{plant: 'B', area: 1, perimeter: 4},
		{plant: 'C', area: 1, perimeter: 4},
		{plant: 'D', area: 1, perimeter: 4},
	}
	assert.ElementsMatch(t, expected, actual)
}

func Test2x2RegionOfSamePlant(t *testing.T) {
	var actual = ParseGardenMap(
		"AA\n" +
		"AA\n",
	)

	var expected = GardenMap{
		{plant: 'A', area: 4, perimeter: 8},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestParseSimpleProvidedExampleGardenMap(t *testing.T) {
	var actual = ParseGardenMap(simulateFileContent(SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES))

	var expected = GardenMap{
		{plant: 'A', area: 4, perimeter: 10},
		{plant: 'B', area: 4, perimeter: 8},
		{plant: 'C', area: 4, perimeter: 10},
		{plant: 'D', area: 1, perimeter: 4},
		{plant: 'E', area: 3, perimeter: 8},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestParseAnotherProvidedExampleGardenMap(t *testing.T) {
	var actual = ParseGardenMap(simulateFileContent(ANOTHER_PROVIDED_EXAMPLE_INPUT_LINES))

	var expected = GardenMap{
		{plant: 'X', area: 1, perimeter: 4},
		{plant: 'X', area: 1, perimeter: 4},
		{plant: 'X', area: 1, perimeter: 4},
		{plant: 'O', area: 21, perimeter: 36},
		{plant: 'X', area: 1, perimeter: 4},
	}
	assert.ElementsMatch(t, expected, actual)
}

func TestParseLargeProvidedExampleGardenMap(t *testing.T) {
	var actual = ParseGardenMap(simulateFileContent(LARGE_PROVIDED_EXAMPLE_INPUT_LINES))

	var expected = GardenMap{
		{plant: 'R', area: 12, perimeter: 18},
		{plant: 'I', area: 4, perimeter: 8},
		{plant: 'C', area: 14, perimeter: 28},
		{plant: 'F', area: 10, perimeter: 18},
		{plant: 'V', area: 13, perimeter: 20},
		{plant: 'J', area: 11, perimeter: 20},
		{plant: 'C', area: 1, perimeter: 4},
		{plant: 'E', area: 13, perimeter: 18},
		{plant: 'I', area: 14, perimeter: 22},
		{plant: 'M', area: 5, perimeter: 12},
		{plant: 'S', area: 3, perimeter: 8},
	}
	assert.ElementsMatch(t, expected, actual)
}
