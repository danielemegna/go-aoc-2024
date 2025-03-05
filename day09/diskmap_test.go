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
