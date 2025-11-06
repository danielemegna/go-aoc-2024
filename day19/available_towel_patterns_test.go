package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var patterns = AvailableTowelPatternsFrom("r, wr, b, g, bwu, rb, gb, br")

func TestSimplePossiblePatterns(t *testing.T) {
	assert.True(t, patterns.IsDesignPossible(TowelPattern("r")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("br")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("bwu")))
}

func TestSimpleImpossiblePatterns(t *testing.T) {
	assert.False(t, patterns.IsDesignPossible(TowelPattern("w")))
	assert.False(t, patterns.IsDesignPossible(TowelPattern("rw")))
	assert.False(t, patterns.IsDesignPossible(TowelPattern("bru")))
	assert.False(t, patterns.IsDesignPossible(TowelPattern("pippo")))
}

func TestSimplePossibleComposedPatterns(t *testing.T) {
	assert.True(t, patterns.IsDesignPossible(TowelPattern("bg")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("rbgb")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("wrbr")))
}

func TestProvidedPossiblePatterns(t *testing.T) {
	assert.True(t, patterns.IsDesignPossible(TowelPattern("brwrr")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("bggr")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("gbbr")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("rrbgbr")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("bwurrg")))
	assert.True(t, patterns.IsDesignPossible(TowelPattern("brgr")))
}

func TestProvidedImpossiblePatterns(t *testing.T) {
	assert.False(t, patterns.IsDesignPossible(TowelPattern("ubwu")))
	assert.False(t, patterns.IsDesignPossible(TowelPattern("bbrgwb")))
}

func TestNoPossibleDesignsCombinations(t *testing.T) {
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("w")))
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("w")))
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("rw")))
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("bru")))
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("ubwu")))
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("bbrgwb")))
	assert.Equal(t, 0, patterns.PossibleDesignCombinationsFor(TowelPattern("pippo")))
}

func TestOnePossibleDesignsCombinations(t *testing.T) {
	assert.Equal(t, 1, patterns.PossibleDesignCombinationsFor(TowelPattern("bggr")))
	assert.Equal(t, 1, patterns.PossibleDesignCombinationsFor(TowelPattern("bwurrg")))
}

func TestTwoPossibleDesignsCombinations(t *testing.T) {
	t.Skip("WIP")
	assert.Equal(t, 2, patterns.PossibleDesignCombinationsFor(TowelPattern("brwrr")))
	assert.Equal(t, 2, patterns.PossibleDesignCombinationsFor(TowelPattern("brgr")))
}

func TestManyPossibleDesignsCombinations(t *testing.T) {
	t.Skip("WIP")
	assert.Equal(t, 4, patterns.PossibleDesignCombinationsFor(TowelPattern("gbbr")))
	assert.Equal(t, 6, patterns.PossibleDesignCombinationsFor(TowelPattern("rrbgbr")))
}
