package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var patterns = InitAvailableTowelPatternsFrom("r, wr, b, g, bwu, rb, gb, br")

func TestSimplePossiblePatterns(t *testing.T) {
	assert.True(t, TowelPattern("r").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("br").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("bwu").IsDesignPossibleWith(patterns))
}

func TestSimpleImpossiblePatterns(t *testing.T) {
	assert.False(t, TowelPattern("w").IsDesignPossibleWith(patterns))
	assert.False(t, TowelPattern("rw").IsDesignPossibleWith(patterns))
	assert.False(t, TowelPattern("bru").IsDesignPossibleWith(patterns))
}

func TestSimplePossibleComposedPatterns(t *testing.T) {
	assert.True(t, TowelPattern("bg").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("rbgb").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("wrbr").IsDesignPossibleWith(patterns))
}

func TestProvidedPossiblePatterns(t *testing.T) {
	assert.True(t, TowelPattern("brwrr").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("bggr").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("gbbr").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("rrbgbr").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("bwurrg").IsDesignPossibleWith(patterns))
	assert.True(t, TowelPattern("brgr").IsDesignPossibleWith(patterns))
}

func TestProvidedImpossiblePatterns(t *testing.T) {
	assert.False(t, TowelPattern("ubwu").IsDesignPossibleWith(patterns))
	assert.False(t, TowelPattern("bbrgwb").IsDesignPossibleWith(patterns))
}
