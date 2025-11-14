package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyAntennaGroup(t *testing.T) {
	var group = AntennaGroup{frequency: 'a'}
	assert.Equal(t, 'a', group.frequency)
	assert.Empty(t, group.locations)
	assert.Empty(t, group.antinodes)
}

func TestSingleAntennaGroup_NoAntinodes(t *testing.T) {
	var group = AntennaGroup{}

	group.AddAntennaAt(Coordinate{4, 3})

	assert.Equal(t, []Coordinate{{4, 3}}, group.locations)
	assert.Empty(t, group.antinodes)
}

func TestTwoVerticalAntennaGroup_TwoVerticalAntinodes(t *testing.T) {
	var group = AntennaGroup{}

	group.AddAntennaAt(Coordinate{4, 3})
	group.AddAntennaAt(Coordinate{4, 5})

	assert.Len(t, group.locations, 2)
	assert.ElementsMatch(t, []Coordinate{{4, 1}, {4, 7}}, group.antinodes)
}

func TestTwoHorizontalAntennaGroup_TwoHorizontalAntinodes(t *testing.T) {
	var group = AntennaGroup{}

	group.AddAntennaAt(Coordinate{4, 3})
	group.AddAntennaAt(Coordinate{5, 3})

	assert.Len(t, group.locations, 2)
	assert.ElementsMatch(t, []Coordinate{{3, 3}, {6, 3}}, group.antinodes)
}

func TestTwoDiagonalAntennaGroup_TwoDiagonalAntinodes(t *testing.T) {
	var group = AntennaGroup{}

	group.AddAntennaAt(Coordinate{4, 3})
	group.AddAntennaAt(Coordinate{5, 5})

	assert.Len(t, group.locations, 2)
	assert.ElementsMatch(t, []Coordinate{{3, 1}, {6, 7}}, group.antinodes)
}
