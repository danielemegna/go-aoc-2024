package day12

type GardenMap []GardenRegion

type GardenRegion struct {
	plant     rune
	area      int
	perimeter GardenRegionPerimeter
}

func NewGardenRegion(plant rune) GardenRegion {
	return GardenRegion{
		plant:     plant,
		area:      0,
		perimeter: NewGardenRegionPerimeter(),
	}
}

func (this GardenRegion) AddPerimeterPiece(coordinate Coordinate, closeCoordinate Coordinate) {
	var border = NewBorderFor(coordinate, closeCoordinate)
	this.perimeter.AddBorder(border)
}
