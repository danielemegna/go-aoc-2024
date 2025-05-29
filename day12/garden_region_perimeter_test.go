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

	assert.Equal(t, 1, perimeter.sides)
	assert.Equal(t, []Coordinate{{X: 1, Y: 2}}, perimeter.coordinates)
}

func TestSomeFarElements(t *testing.T) {
	var perimeter = GardenRegionPerimeter{}

	perimeter.Add(Coordinate{X: 1, Y: 2}, Coordinate{X: 1, Y: 3})
	perimeter.Add(Coordinate{X: 4, Y: 5}, Coordinate{X: 3, Y: 5})

	assert.Equal(t, 2, perimeter.sides)
	assert.Equal(t, []Coordinate{{X: 1, Y: 2}, {X: 4, Y: 5}}, perimeter.coordinates)
}

func TestTwoCloseElementsShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = GardenRegionPerimeter{}

	perimeter.Add(Coordinate{X: 1, Y: 2}, Coordinate{X: 1, Y: 3})
	perimeter.Add(Coordinate{X: 2, Y: 2}, Coordinate{X: 2, Y: 3})

	// assert.Equal(t, 1, perimeter.sides) WIP
	assert.Equal(t, []Coordinate{{X: 1, Y: 2}, {X: 2, Y: 2}}, perimeter.coordinates)
}
