package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSingleElement(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	var coordinate = Coordinate{X: 1, Y: 3}
	perimeter.AddBorder(coordinate, NORTH)

	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestSomeFarElements(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Coordinate{X: 1, Y: 2}, SOUTH)
	perimeter.AddBorder(Coordinate{X: 4, Y: 5}, WEST)

	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 2, perimeter.NumberOfSides())
}

func TestSomeDifferentSideCloseElements(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Coordinate{X: 1, Y: 1}, SOUTH)
	perimeter.AddBorder(Coordinate{X: 2, Y: 3}, NORTH)
	perimeter.AddBorder(Coordinate{X: 3, Y: 1}, SOUTH)

	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 3, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 6, Y: 7}, NORTH)
	perimeter.AddBorder(Coordinate{X: 5, Y: 6}, EAST)

	assert.Equal(t, 5, perimeter.Length())
	assert.Equal(t, 5, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 9, Y: 8}, SOUTH)
	perimeter.AddBorder(Coordinate{X: 10, Y: 10}, NORTH)

	assert.Equal(t, 7, perimeter.Length())
	assert.Equal(t, 7, perimeter.NumberOfSides())
}

func TestCloseSouthBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Coordinate{X: 1, Y: 2}, SOUTH)
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 2, Y: 2}, SOUTH)
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 0, Y: 2}, SOUTH)
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestCloseNouthBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Coordinate{X: 1, Y: 2}, NORTH)
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 2, Y: 2}, NORTH)
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 0, Y: 2}, NORTH)
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestCloseEastBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Coordinate{X: 2, Y: 3}, EAST)
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 2, Y: 4}, EAST)
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Coordinate{X: 2, Y: 2}, EAST)
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestSinglePlant(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()
	var singlePlantCoordinate = Coordinate{X: 3, Y: 3}

	perimeter.AddBorder(singlePlantCoordinate, NORTH)
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(singlePlantCoordinate, SOUTH)
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 2, perimeter.NumberOfSides())

	perimeter.AddBorder(singlePlantCoordinate, WEST)
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 3, perimeter.NumberOfSides())

	perimeter.AddBorder(singlePlantCoordinate, EAST)
	assert.Equal(t, 4, perimeter.Length())
	assert.Equal(t, 4, perimeter.NumberOfSides())
}
