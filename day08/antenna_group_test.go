package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyAntennaGroup(t *testing.T) {
	var group = AntennaGroup{}
	assert.Empty(t, group.locations)
	assert.Empty(t, group.antinodes)
}

func TestAddAntennaAt(t *testing.T) {

	t.Run("SingleAntennaGroup_NoAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAt(Coordinate{4, 3})

		assert.Equal(t, []Coordinate{{4, 3}}, group.locations)
		assert.Empty(t, group.antinodes)
	})

	t.Run("TwoVerticalAntennaGroup_TwoVerticalAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{4, 5})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{{4, 1}, {4, 7}}, group.antinodes)
	})

	t.Run("TwoHorizontalAntennaGroup_TwoHorizontalAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{5, 3})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{{3, 3}, {6, 3}}, group.antinodes)
	})

	t.Run("TwoDiagonalAntennaGroup_TwoDiagonalAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{5, 5})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{{3, 1}, {6, 7}}, group.antinodes)
	})

	t.Run("ThreeAntennaGroup_SixAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{5, 5})
		group.AddAntennaAt(Coordinate{8, 4})

		assert.Len(t, group.locations, 3)
		assert.ElementsMatch(t, []Coordinate{
			{3, 1}, {6, 7},
			{0, 2}, {12, 5},
			{2, 6}, {11, 3},
		}, group.antinodes)
	})

}

func TestAddAntennaWithResonantHarmonics(t *testing.T) {

	t.Run("SingleAntennaGroup_NoAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAtWithResonantHarmonics(Coordinate{4, 3}, 10)

		assert.Equal(t, []Coordinate{{4, 3}}, group.locations)
		assert.Empty(t, group.antinodes)
	})

	t.Run("TwoVerticalAntennaGroup_TenSizeMap_FiveVerticalAntinodes", func(t *testing.T) {
		var group = AntennaGroup{}

		group.AddAntennaAtWithResonantHarmonics(Coordinate{4, 3}, 10)
		group.AddAntennaAtWithResonantHarmonics(Coordinate{4, 5}, 10)

		assert.ElementsMatch(t, []Coordinate{{4, 3}, {4, 5}}, group.locations)
		assert.ElementsMatch(t, []Coordinate{{4, 3}, {4, 5}, {4, 1}, {4, 7}, {4, 9}}, group.antinodes)
	})

}
