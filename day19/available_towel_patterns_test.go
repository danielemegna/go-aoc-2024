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
