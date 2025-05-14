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
	assert.Equal(t, expected, actual)
}

func TestTwoDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("AB\n")

	var expected = GardenMap{
		{plant: 'A', area: 1, perimeter: 4},
		{plant: 'B', area: 1, perimeter: 4},
	}
	assert.Equal(t, expected, actual)
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
	assert.Equal(t, expected, actual)
}

func Test2x2RegionOfSamePlant(t *testing.T) {
	t.Skip("refactoring")
	var actual = ParseGardenMap(
		"AA\n" +
		"AA\n",
	)

	var expected = GardenMap{
		{plant: 'A', area: 4, perimeter: 8},
	}
	assert.Equal(t, expected, actual)
}

func TestParseSimpleProvidedExampleGardenMap(t *testing.T) {
	t.Skip("refactoring")
	var actual = ParseGardenMap(simulateFileContent(SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES))

	var expected = GardenMap{
		{plant: 'A', area: 4, perimeter: 10},
		{plant: 'B', area: 4, perimeter: 8},
		{plant: 'C', area: 4, perimeter: 10},
		{plant: 'D', area: 1, perimeter: 4},
		{plant: 'E', area: 3, perimeter: 8},
	}
	assert.Equal(t, expected, actual)
}
