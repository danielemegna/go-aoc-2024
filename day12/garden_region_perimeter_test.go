package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSingleElement(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	var coordinate = Coordinate{X: 1, Y: 3}
	perimeter.AddBorder(Border{coordinate: coordinate, direction: NORTH})

	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestSomeFarElements(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 1, Y: 2}, direction: SOUTH})
	perimeter.AddBorder(Border{coordinate: Coordinate{X: 4, Y: 5}, direction: WEST})

	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 2, perimeter.NumberOfSides())
}

func TestSomeDifferentSideCloseElements(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 1, Y: 1}, direction: SOUTH})
	perimeter.AddBorder(Border{coordinate: Coordinate{X: 2, Y: 3}, direction: NORTH})
	perimeter.AddBorder(Border{coordinate: Coordinate{X: 3, Y: 1}, direction: SOUTH})

	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 3, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 6, Y: 7}, direction: NORTH})
	perimeter.AddBorder(Border{coordinate: Coordinate{X: 5, Y: 6}, direction: EAST})

	assert.Equal(t, 5, perimeter.Length())
	assert.Equal(t, 5, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 9, Y: 8}, direction: SOUTH})
	perimeter.AddBorder(Border{coordinate: Coordinate{X: 10, Y: 10}, direction: NORTH})

	assert.Equal(t, 7, perimeter.Length())
	assert.Equal(t, 7, perimeter.NumberOfSides())
}

func TestCloseSouthBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 1, Y: 2}, direction: SOUTH})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 2, Y: 2}, direction: SOUTH})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 0, Y: 2}, direction: SOUTH})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestCloseNouthBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 1, Y: 2}, direction: NORTH})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 2, Y: 2}, direction: NORTH})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 0, Y: 2}, direction: NORTH})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestCloseEastBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 2, Y: 3}, direction: EAST})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 2, Y: 4}, direction: EAST})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: Coordinate{X: 2, Y: 2}, direction: EAST})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestSinglePlant(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()
	var singlePlantCoordinate = Coordinate{X: 3, Y: 3}

	perimeter.AddBorder(Border{coordinate: singlePlantCoordinate, direction: NORTH})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: singlePlantCoordinate, direction: SOUTH})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 2, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: singlePlantCoordinate, direction: WEST})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 3, perimeter.NumberOfSides())

	perimeter.AddBorder(Border{coordinate: singlePlantCoordinate, direction: EAST})
	assert.Equal(t, 4, perimeter.Length())
	assert.Equal(t, 4, perimeter.NumberOfSides())
}
