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
