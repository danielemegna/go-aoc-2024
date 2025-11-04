package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPatternAvailable(t *testing.T) {
	var patterns = InitAvailableTowelPatternsFrom("r, wr, b, g, bwu, rb, gb, br")

	assert.True(t, patterns.IsStringPatternAvailable("r"))
	assert.True(t, patterns.IsStringPatternAvailable("wr"))
	assert.True(t, patterns.IsStringPatternAvailable("bwu"))
	assert.False(t, patterns.IsStringPatternAvailable("w"))
	assert.False(t, patterns.IsStringPatternAvailable("u"))
	assert.False(t, patterns.IsStringPatternAvailable("gr"))
	assert.False(t, patterns.IsStringPatternAvailable("gnr"))
}

func TestMaxPatternLength(t *testing.T) {
	var patterns = InitAvailableTowelPatternsFrom("r, wr, b, g, bwu, rb, br")

	assert.Equal(t, 1, patterns.MaxPatternLengthFor('g'))
	assert.Equal(t, 2, patterns.MaxPatternLengthFor('w'))
	assert.Equal(t, 2, patterns.MaxPatternLengthFor('r'))
	assert.Equal(t, 3, patterns.MaxPatternLengthFor('b'))
	assert.Equal(t, 0, patterns.MaxPatternLengthFor('u'))
}
