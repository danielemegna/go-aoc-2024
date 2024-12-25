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
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: -1, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: -1}))
}

func TestNoXMasOccurrencesInSmallMap(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", "M", "A"},
		{".", ".", "."},
		{".", ".", "."},
	}

	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
}

func TestSingleXMasOccurrenceAtZeroZeroVertical(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", ".", ".", "."},
		{"M", ".", ".", "."},
		{"A", ".", ".", "."},
		{"S", ".", ".", "."},
	}

	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 1, Y: 0}))
	assert.Equal(t, 0, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 1}))
}

func TestTwoXMasOccurrencesAtZeroZero(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", "M", "A", "S"},
		{"M", ".", ".", "."},
		{"A", ".", ".", "."},
		{"S", ".", ".", "."},
	}

	assert.Equal(t, 2, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
}

func TestSingleXMasOccurrenceAtZeroZeroDiagonal(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", ".", ".", "."},
		{".", "M", ".", "."},
		{".", ".", "A", "."},
		{".", ".", ".", "S"},
	}

	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
}

func TestSingleInverseXMasOccurrence(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{".", ".", ".", "."},
		{"S", "A", "M", "X"},
	}

	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 3, Y: 3}))
}

func TestThreeInverseXMasOccurrencesFromSameCoordinate(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"S", ".", ".", "S"},
		{".", "A", ".", "A"},
		{".", ".", "M", "M"},
		{"S", "A", "M", "X"},
	}

	assert.Equal(t, 3, m.XMasOccurrencesAt(Coordinate{X: 3, Y: 3}))
}

func TestBothNorthAndSouthDiagonalOccurrences(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", ".", ".", "S"},
		{".", "M", "A", "."},
		{".", "M", "A", "."},
		{"X", ".", ".", "S"},
	}

	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 3}))
}

func TestBothNorthAndSouthInverseDiagonalOccurrences(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"S", ".", ".", "X"},
		{".", "A", "M", "."},
		{".", "A", "M", "."},
		{"S", ".", ".", "X"},
	}

	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 3, Y: 0}))
	assert.Equal(t, 1, m.XMasOccurrencesAt(Coordinate{X: 3, Y: 3}))
}

func TestSomeXMasOccurrences(t *testing.T) {
	var m CharactersMap = CharactersMap{
		{"X", "M", "A", "S", ".", "."},
		{"M", "M", "S", ".", ".", "S"},
		{"A", ".", "A", "A", ".", "A"},
		{"S", ".", ".", "S", "M", "M"},
		{".", ".", "S", "A", "M", "X"},
		{".", ".", ".", ".", ".", "."},
	}

	assert.Equal(t, 3, m.XMasOccurrencesAt(Coordinate{X: 0, Y: 0}))
	assert.Equal(t, 3, m.XMasOccurrencesAt(Coordinate{X: 5, Y: 4}))
}
