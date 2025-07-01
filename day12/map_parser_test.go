package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSinglePlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("A\n")

	assert.Equal(t, 1, len(actual))

	assert.Equal(t, 'A', actual[0].plant)
	assert.Equal(t, 1, actual[0].area)
	assert.Equal(t, 4, actual[0].perimeter.Length())
	assert.Equal(t, 4, actual[0].perimeter.NumberOfSides())
}

func TestTwoDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap("AB\n")

	assert.Equal(t, 2, len(actual))

	assert.Equal(t, 'A', actual[0].plant)
	assert.Equal(t, 1, actual[0].area)
	assert.Equal(t, 4, actual[0].perimeter.Length())
	assert.Equal(t, 4, actual[0].perimeter.NumberOfSides())

	assert.Equal(t, 'B', actual[1].plant)
	assert.Equal(t, 1, actual[1].area)
	assert.Equal(t, 4, actual[1].perimeter.Length())
	assert.Equal(t, 4, actual[1].perimeter.NumberOfSides())
}

func TestFourDifferentPlantGardenMap(t *testing.T) {
	var actual = ParseGardenMap(
		"AB\n" +
		"CD\n",
	)

	assert.Equal(t, 4, len(actual))
}

func Test2x2RegionOfSamePlant(t *testing.T) {
	var actual = ParseGardenMap(
		"AA\n" +
		"AA\n",
	)

	assert.Equal(t, 'A', actual[0].plant)
	assert.Equal(t, 4, actual[0].area)
	assert.Equal(t, 8, actual[0].perimeter.Length())
	assert.Equal(t, 4, actual[0].perimeter.NumberOfSides())
}

func TestParseSimpleProvidedExampleGardenMap(t *testing.T) {
	var actual = ParseGardenMap(simulateFileContent(SIMPLE_PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 5, len(actual))

	assert.Equal(t, 'A', actual[0].plant)
	assert.Equal(t, 4, actual[0].area)
	assert.Equal(t, 10, actual[0].perimeter.Length())
	assert.Equal(t, 4, actual[0].perimeter.NumberOfSides())

	assert.Equal(t, 'B', actual[1].plant)
	assert.Equal(t, 4, actual[1].area)
	assert.Equal(t, 8, actual[1].perimeter.Length())
	assert.Equal(t, 4, actual[1].perimeter.NumberOfSides())

	assert.Equal(t, 'C', actual[2].plant)
	assert.Equal(t, 4, actual[2].area)
	assert.Equal(t, 10, actual[2].perimeter.Length())
	assert.Equal(t, 8, actual[2].perimeter.NumberOfSides())

	assert.Equal(t, 'D', actual[3].plant)
	assert.Equal(t, 1, actual[3].area)
	assert.Equal(t, 4, actual[3].perimeter.Length())
	assert.Equal(t, 4, actual[3].perimeter.NumberOfSides())

	assert.Equal(t, 'E', actual[4].plant)
	assert.Equal(t, 3, actual[4].area)
	assert.Equal(t, 8, actual[4].perimeter.Length())
	assert.Equal(t, 4, actual[4].perimeter.NumberOfSides())
}

func TestParseAnotherProvidedExampleGardenMap(t *testing.T) {
	var actual = ParseGardenMap(simulateFileContent(ANOTHER_PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 1+4, len(actual))

	assert.Equal(t, 'O', actual[0].plant)
	assert.Equal(t, 21, actual[0].area)
	assert.Equal(t, 36, actual[0].perimeter.Length())
	assert.Equal(t, 20, actual[0].perimeter.NumberOfSides())

	assert.Equal(t, 'X', actual[1].plant)
	assert.Equal(t, 1, actual[1].area)
	assert.Equal(t, 4, actual[1].perimeter.Length())
	assert.Equal(t, 4, actual[1].perimeter.NumberOfSides())

	assert.Equal(t, 'X', actual[2].plant)
	assert.Equal(t, 1, actual[2].area)
	assert.Equal(t, 4, actual[2].perimeter.Length())
	assert.Equal(t, 4, actual[2].perimeter.NumberOfSides())

	assert.Equal(t, 'X', actual[3].plant)
	assert.Equal(t, 1, actual[3].area)
	assert.Equal(t, 4, actual[3].perimeter.Length())
	assert.Equal(t, 4, actual[3].perimeter.NumberOfSides())

	assert.Equal(t, 'X', actual[4].plant)
	assert.Equal(t, 1, actual[4].area)
	assert.Equal(t, 4, actual[4].perimeter.Length())
	assert.Equal(t, 4, actual[4].perimeter.NumberOfSides())
}

func TestParseLargeProvidedExampleGardenMap(t *testing.T) {
	var actual = ParseGardenMap(simulateFileContent(LARGE_PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 11, len(actual))

	assert.Equal(t, 'R', actual[0].plant)
	assert.Equal(t, 12, actual[0].area)
	assert.Equal(t, 18, actual[0].perimeter.Length())
	assert.Equal(t, 10, actual[0].perimeter.NumberOfSides())

	assert.Equal(t, 'V', actual[4].plant)
	assert.Equal(t, 13, actual[4].area)
	assert.Equal(t, 20, actual[4].perimeter.Length())
	assert.Equal(t, 10, actual[4].perimeter.NumberOfSides())

	assert.Equal(t, 'I', actual[8].plant)
	assert.Equal(t, 14, actual[8].area)
	assert.Equal(t, 22, actual[8].perimeter.Length())
	assert.Equal(t, 16, actual[8].perimeter.NumberOfSides())
}
