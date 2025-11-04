package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPatternAvailable(t *testing.T) {
	var patterns = InitAvailableTowelPatternsFrom("r, wr, b, g, bwu, rb, gb, br")

	assert.True(t, patterns.IsAvailable("r"))
	assert.True(t, patterns.IsAvailable("wr"))
	assert.True(t, patterns.IsAvailable("bwu"))
	assert.False(t, patterns.IsAvailable("w"))
	assert.False(t, patterns.IsAvailable("u"))
	assert.False(t, patterns.IsAvailable("gr"))
	assert.False(t, patterns.IsAvailable("gnr"))
}

func TestMaxPatternLength(t *testing.T) {
	var patterns = InitAvailableTowelPatternsFrom("r, wr, b, g, bwu, rb, br")

	assert.Equal(t, 1, patterns.MaxPatternLengthFor('g'))
	assert.Equal(t, 2, patterns.MaxPatternLengthFor('w'))
	assert.Equal(t, 2, patterns.MaxPatternLengthFor('r'))
	assert.Equal(t, 3, patterns.MaxPatternLengthFor('b'))
	assert.Equal(t, 0, patterns.MaxPatternLengthFor('u'))
}
