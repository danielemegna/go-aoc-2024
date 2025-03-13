package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefragSimpleMap(t *testing.T) {
	var diskMap = ExpandedDiskMap{data: []int{
		0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2,
	}}

	var actual = Defrag(diskMap)

	var expected = ExpandedDiskMap{data: []int{
		0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1,
	}}
	assert.Equal(t, expected, actual)
}

func TestDefragProvidedExampleMap(t *testing.T) {
	var diskMap = ExpandedDiskMap{data: []int{
		0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1,
		3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1,
		7, 7, 7, -1, 8, 8, 8, 8, 9, 9,
	}}

	var actual = Defrag(diskMap)

	var expected = ExpandedDiskMap{data: []int{
		0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}}
	assert.Equal(t, expected, actual)
}

func TestDefragSingleDigitMapDoesNothing(t *testing.T) {
	var actual = Defrag(ExpandedDiskMap{data: []int{0}})
	var expected = ExpandedDiskMap{data: []int{0}}
	assert.Equal(t, expected, actual)
}

func TestDefragMapWithoutEmptySpacesDoesNothing(t *testing.T) {
	var actual = Defrag(ExpandedDiskMap{data: []int{0, 0, 0, 1, 1, 2, 3, 3, 4, 4, 4}})
	var expected = ExpandedDiskMap{data: []int{0, 0, 0, 1, 1, 2, 3, 3, 4, 4, 4}}
	assert.Equal(t, expected, actual)
}

func TestDefragFileSimpleDenseMap(t *testing.T) {
	t.Skip("WIP")
	// 00...111.2.3.444 -> 0044411132......
	var diskMap = ExpandedDiskMap{data: []int{
		0, 0, -1, -1, -1, 1, 1, 1, -1, 2, -1, 3, -1, 4, 4, 4,
	}}

	var actual = DefragWholeFiles(diskMap)

	var expected = ExpandedDiskMap{data: []int{
		0, 0, 4, 4, 4, 1, 1, 1, 3, 2, -1, -1, -1, -1, -1, -1,
	}}
	assert.Equal(t, expected, actual)
}
