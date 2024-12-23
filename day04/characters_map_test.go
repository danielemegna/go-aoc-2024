package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyMapHasNoXMasOccurrences(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 1, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 1, Y: 1}))
}

func TestSingleXMasOccurrenceAtZeroZeroHorizontal(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", "M", "A", "S"},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 1, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 1}))
}

func TestUncompleteXMasHorizontalOccurrence(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{".", ".", ".", "."},
		{"X", "M", "A", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 1}))
}

func TestNoXMasOccurrencesOutOfMapBounds(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", "M", "A", "S"},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
	}

	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 99, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 99}))
}

func TestNoXMasOccurrencesInSmallMap(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", "M", "A"},
		{".", ".", "."},
		{".", ".", "."},
	}

	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
}
