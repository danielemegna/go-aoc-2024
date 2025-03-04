package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDenseDiskMap(t *testing.T) {
	var expected = DenseDiskMap{
		data: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
	}

	var actual = ParseDenseDiskMap("2333133121414131402")

	assert.Equal(t, expected, actual)
}
