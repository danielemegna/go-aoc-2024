package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var EMPTY_CITY_MAP_EXAMPLE_INPUT_LINES = []string{
	"......",
	"......",
	"......",
	"......",
	"......",
	"......",
}

var SIMPLE_CITY_MAP_EXAMPLE_INPUT_LINES = []string{
	"........",
	"...A....",
	"..A.....",
	"........",
	"....B...",
	"........",
	"....B...",
	"........",
}

func TestParseAndInspectEmptyCityMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent(EMPTY_CITY_MAP_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 6, cityMap.Size())
	assert.Equal(t, []Coordinate{}, cityMap.Antennas())
	assert.Len(t, cityMap.AntinodesInMap(), 0)
}

func TestParseAndInspectSimpleCityMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent(SIMPLE_CITY_MAP_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 8, cityMap.Size())
	assert.ElementsMatch(t, []Coordinate{{3, 1}, {2, 2}, {4, 4}, {4, 6}}, cityMap.Antennas())
	assert.Len(t, cityMap.AntinodesInMap(), 4)
}

func TestParseAndInspectProvidedExampleMap(t *testing.T) {
	var cityMap = ParseCityMap(simulateFileContent(PROVIDED_EXAMPLE_INPUT_LINES))

	assert.Equal(t, 12, cityMap.Size())
	assert.ElementsMatch(t, []Coordinate{{8, 1}, {5, 2}, {7, 3}, {4, 4}, {6, 5}, {8, 8}, {9, 9}}, cityMap.Antennas())
	// TODO exclude out of bound antinodes
	// assert.Len(t, cityMap.AntinodesInMap(), 13)
}
