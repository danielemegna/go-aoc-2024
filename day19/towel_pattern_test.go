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
}
