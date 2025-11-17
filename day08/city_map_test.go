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
	}))

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
	}))

	assert.Equal(t, 8, cityMap.Size())
	assert.ElementsMatch(t, []Coordinate{{3, 1}, {2, 2}, {6, 3}, {6, 5}}, cityMap.Antennas())
	assert.ElementsMatch(t, []Coordinate{{4, 0}, {1, 3}, {6, 1}, {6, 7}}, cityMap.AntinodesInMap())
}

func TestParseAndInspectProvidedExampleMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 12, cityMap.Size())
	assert.ElementsMatch(t, []Coordinate{{8, 1}, {5, 2}, {7, 3}, {4, 4}, {6, 5}, {8, 8}, {9, 9}}, cityMap.Antennas())
	// TODO exclude duplicated and out of bound antinodes
	// assert.Len(t, cityMap.AntinodesInMap(), 13+1)
}
