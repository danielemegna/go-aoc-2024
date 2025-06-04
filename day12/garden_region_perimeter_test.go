package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSingleElement(t *testing.T) {
	var perimeter = GardenRegionPerimeter{}

	var outsideCoordinate = Coordinate{X: 1, Y: 2}
	var insideCoordinate = Coordinate{X: 1, Y: 3}
	perimeter.Add(outsideCoordinate, insideCoordinate)

	assert.Equal(t, 1, perimeter.NumberOfSides())
	assert.Equal(t, 1, perimeter.Length())
}

func TestSomeFarElements(t *testing.T) {
	var perimeter = GardenRegionPerimeter{}

	perimeter.Add(Coordinate{X: 1, Y: 2}, Coordinate{X: 1, Y: 3})
	perimeter.Add(Coordinate{X: 4, Y: 5}, Coordinate{X: 3, Y: 5})

	assert.Equal(t, 2, perimeter.NumberOfSides())
	assert.Equal(t, 2, perimeter.Length())
}

func TestTwoCloseElementsShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = GardenRegionPerimeter{}

	perimeter.Add(Coordinate{X: 1, Y: 2}, Coordinate{X: 1, Y: 3})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 2, Y: 2}, Coordinate{X: 2, Y: 3})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 0, Y: 2}, Coordinate{X: 0, Y: 3})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}
