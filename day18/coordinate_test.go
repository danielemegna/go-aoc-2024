package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCoordinate(t *testing.T) {
	assert.Equal(t, Coordinate{18,2}, ParseCoordinateFrom("18,2"))
	assert.Equal(t, Coordinate{10,0}, ParseCoordinateFrom("10,0"))
	assert.Equal(t, Coordinate{54,3}, ParseCoordinateFrom("54,3"))
}
