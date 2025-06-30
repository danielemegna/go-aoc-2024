package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSingleElement(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	var insideCoordinate = Coordinate{X: 1, Y: 3}
	var outsideCoordinate = Coordinate{X: 1, Y: 2}
	perimeter.Add(insideCoordinate, outsideCoordinate)

	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestSomeFarElements(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.Add(Coordinate{X: 1, Y: 2}, Coordinate{X: 1, Y: 3})
	perimeter.Add(Coordinate{X: 4, Y: 5}, Coordinate{X: 3, Y: 5})

	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 2, perimeter.NumberOfSides())
}

func TestSomeDifferentSideCloseElements(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.Add(Coordinate{X: 1, Y: 1}, Coordinate{X: 1, Y: 2})
	perimeter.Add(Coordinate{X: 2, Y: 3}, Coordinate{X: 2, Y: 2})
	perimeter.Add(Coordinate{X: 3, Y: 1}, Coordinate{X: 3, Y: 2})

	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 3, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 6, Y: 7}, Coordinate{X: 6, Y: 6})
	perimeter.Add(Coordinate{X: 5, Y: 6}, Coordinate{X: 6, Y: 6})

	assert.Equal(t, 5, perimeter.Length())
	assert.Equal(t, 5, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 9, Y: 8}, Coordinate{X: 9, Y: 9})
	perimeter.Add(Coordinate{X: 10, Y: 10}, Coordinate{X: 10, Y: 9})

	assert.Equal(t, 7, perimeter.Length())
	assert.Equal(t, 7, perimeter.NumberOfSides())
}

func TestCloseSouthBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

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

func TestCloseNouthBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.Add(Coordinate{X: 1, Y: 2}, Coordinate{X: 1, Y: 1})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 2, Y: 2}, Coordinate{X: 2, Y: 1})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 0, Y: 2}, Coordinate{X: 0, Y: 1})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestCloseEastBorderShouldNotIncreaseSideCount(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()

	perimeter.Add(Coordinate{X: 2, Y: 3}, Coordinate{X: 3, Y: 3})
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 2, Y: 4}, Coordinate{X: 3, Y: 4})
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 2, Y: 2}, Coordinate{X: 3, Y: 2})
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())
}

func TestSinglePlant(t *testing.T) {
	var perimeter = NewGardenRegionPerimeter()
	var singlePlantCoordinate = Coordinate{X: 3, Y: 3}

	perimeter.Add(Coordinate{X: 3, Y: 2}, singlePlantCoordinate)
	assert.Equal(t, 1, perimeter.Length())
	assert.Equal(t, 1, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 3, Y: 4}, singlePlantCoordinate)
	assert.Equal(t, 2, perimeter.Length())
	assert.Equal(t, 2, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 2, Y: 3}, singlePlantCoordinate)
	assert.Equal(t, 3, perimeter.Length())
	assert.Equal(t, 3, perimeter.NumberOfSides())

	perimeter.Add(Coordinate{X: 4, Y: 3}, singlePlantCoordinate)
	assert.Equal(t, 4, perimeter.Length())
	assert.Equal(t, 4, perimeter.NumberOfSides())
}
