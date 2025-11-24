package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyAntennaGroup(t *testing.T) {
	var group = NewAntennaGroup(2)
	assert.Empty(t, group.locations)
	assert.Empty(t, group.antinodes)
}

func TestAddAntennaAt(t *testing.T) {

	t.Run("SingleAntennaGroup_NoAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAt(Coordinate{4, 3})

		assert.Equal(t, []Coordinate{{4, 3}}, group.locations)
		assert.Empty(t, group.antinodes)
	})

	t.Run("TwoVerticalAntennaGroup_TenSizeMap_TwoVerticalAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{4, 5})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{{4, 1}, {4, 7}}, group.antinodes)
	})

	t.Run("TwoHorizontalAntennaGroup_TenSizeMap_TwoHorizontalAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{5, 3})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{{3, 3}, {6, 3}}, group.antinodes)
	})

	t.Run("TwoHorizontalAntennaGroup_FourSizeMap_AllAntinodesAreOutOfBounds", func(t *testing.T) {
		var group = NewAntennaGroup(4)

		group.AddAntennaAt(Coordinate{1, 1})
		group.AddAntennaAt(Coordinate{3, 1})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{}, group.antinodes)
	})

	t.Run("TwoDiagonalAntennaGroup_TenSizeMap_TwoDiagonalAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{5, 5})

		assert.Len(t, group.locations, 2)
		assert.ElementsMatch(t, []Coordinate{{3, 1}, {6, 7}}, group.antinodes)
	})

	t.Run("ThreeAntennaGroupExample_TenSizeMap_OnlyFourAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAt(Coordinate{4, 3})
		group.AddAntennaAt(Coordinate{5, 5})
		group.AddAntennaAt(Coordinate{8, 4})

		assert.Len(t, group.locations, 3)
		assert.ElementsMatch(t, []Coordinate{
			{3, 1}, {6, 7},
			{0, 2}, {2, 6},
		}, group.antinodes)
	})

	t.Run("ThreeAntennaGroupProvidedExample_BigMap_SixAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(14)

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
		var group = NewAntennaGroup(10)

		group.AddAntennaAtWithResonantHarmonics(Coordinate{4, 3})

		assert.Equal(t, []Coordinate{{4, 3}}, group.locations)
		assert.Empty(t, group.antinodes)
	})

	t.Run("TwoVerticalAntennaGroup_TenSizeMap_FiveVerticalAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAtWithResonantHarmonics(Coordinate{4, 3})
		group.AddAntennaAtWithResonantHarmonics(Coordinate{4, 5})

		assert.ElementsMatch(t, []Coordinate{{4, 3}, {4, 5}}, group.locations)
		assert.ElementsMatch(t, []Coordinate{{4, 3}, {4, 5}, {4, 1}, {4, 7}, {4, 9}}, group.antinodes)
	})

	t.Run("ThreeAntennaGroupProvidedExample_TenSizeMap_NineAntinodes", func(t *testing.T) {
		var group = NewAntennaGroup(10)

		group.AddAntennaAtWithResonantHarmonics(Coordinate{0, 0})
		group.AddAntennaAtWithResonantHarmonics(Coordinate{1, 2})
		group.AddAntennaAtWithResonantHarmonics(Coordinate{3, 1})

		assert.Len(t, group.locations, 3)
		assert.ElementsMatch(t, []Coordinate{
			{0, 0}, {1, 2}, {3, 1},
			{5, 0}, {6, 2}, {9, 3},
			{2, 4}, {3, 6}, {4, 8},
		}, group.antinodes)
	})

}
