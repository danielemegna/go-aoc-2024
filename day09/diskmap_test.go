package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDenseDiskMap(t *testing.T) {
	var actual = ParseDenseDiskMap("2333133121414131402")
	var expected = DenseDiskMap{data: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}}
	assert.Equal(t, expected, actual)
}

func TestSingleFileDenseDiskMapToExpanded(t *testing.T) {
	var denseDiskMap = DenseDiskMap{data: []int{3}}

	var actual = denseDiskMap.ToExpandedDiskMap()

	var expected = ExpandedDiskMap{data: []int{0, 0, 0}}
	assert.Equal(t, expected, actual)
}

func TestSimpleDenseDiskMapToExpanded(t *testing.T) {
	var denseDiskMap = DenseDiskMap{data: []int{1, 2, 3, 4, 5}}

	var actual = denseDiskMap.ToExpandedDiskMap()

	var expected = ExpandedDiskMap{data: []int{
		0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2,
	}}
	assert.Equal(t, expected, actual)
}

func TestParseAndExpandDenseMap(t *testing.T) {
	// 233111113 -> 00...111.2.3.444
	var parsed = ParseDenseDiskMap("233111113")
	var expanded = parsed.ToExpandedDiskMap()

	var expectedParsed = DenseDiskMap{data: []int{2, 3, 3, 1, 1, 1, 1, 1, 3}}
	var expectedExpanded = ExpandedDiskMap{data: []int{
		0, 0, -1, -1, -1, 1, 1, 1, -1, 2, -1, 3, -1, 4, 4, 4,
	}}
	assert.Equal(t, expectedParsed, parsed)
	assert.Equal(t, expectedExpanded, expanded)
}

func TestProvidedExampleDenseDiskMapToExpanded(t *testing.T) {
	var denseDiskMap = DenseDiskMap{data: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}}

	var actual = denseDiskMap.ToExpandedDiskMap()

	var expected = ExpandedDiskMap{data: []int{
		0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1,
		3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1,
		7, 7, 7, -1, 8, 8, 8, 8, 9, 9,
	}}
	assert.Equal(t, expected, actual)
}

func TestSimpleExpanededDiskMapChecksum(t *testing.T) {
	var expandedDiskMap = ExpandedDiskMap{data: []int{
		0, 2, 2, 1, 1, 1, 2, 2, 2,
	}}

	var checksum = expandedDiskMap.Checksum()

	assert.Equal(t, 2+4+3+4+5+12+14+16, checksum)
}

func TestProvidedExampleExpanededDiskMapChecksum(t *testing.T) {
	var expandedDiskMap = ExpandedDiskMap{data: []int{
		0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7,
		3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}}

	var checksum = expandedDiskMap.Checksum()

	assert.Equal(t, 1928, checksum)
}

func TestAnotherProvidedExampleExpanededDiskMapChecksum(t *testing.T) {
	var expandedDiskMap = ExpandedDiskMap{data: []int{
		0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1,
		5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1,
	}}

	var checksum = expandedDiskMap.Checksum()

	assert.Equal(t, 2858, checksum)
}
