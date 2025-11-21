package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAndInspectEmptyCityMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent([]string{
		"......",
		"......",
		"......",
		"......",
		"......",
		"......",
	}), false)

	assert.Equal(t, 6, cityMap.Size())
	assert.Equal(t, []Coordinate{}, cityMap.Antennas())
	assert.Len(t, cityMap.AntinodesInMap(), 0)
}

func TestParseAndInspectSimpleCityMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent([]string{
		"........",
		"...A....",
		"..A.....",
		"......B.",
		"........",
		"......B.",
		"........",
		"........",
	}), false)

	assert.Equal(t, 8, cityMap.Size())
	assert.ElementsMatch(t, []Coordinate{{3, 1}, {2, 2}, {6, 3}, {6, 5}}, cityMap.Antennas())
	assert.ElementsMatch(t, []Coordinate{{4, 0}, {1, 3}, {6, 1}, {6, 7}}, cityMap.AntinodesInMap())
}

func TestParseAndInspectProvidedExampleMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES), false)

	assert.Equal(t, 12, cityMap.Size())
	assert.ElementsMatch(t, []Coordinate{{8, 1}, {5, 2}, {7, 3}, {4, 4}, {6, 5}, {8, 8}, {9, 9}}, cityMap.Antennas())
	assert.Len(t, cityMap.AntinodesInMap(), 13+1)
}

func TestParseWithResonantHarmonicsAndInspectProvidedExampleMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES), true)

	assert.Equal(t, 12, cityMap.Size())
	assert.Len(t, cityMap.Antennas(), 7)
	assert.Len(t, cityMap.AntinodesInMap(), 27 + 7)
}
