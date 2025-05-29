package day12

type GardenRegionPerimeter struct {
	coordinates []Coordinate
	sides       int
}

func (this *GardenRegionPerimeter) Add(outsideCoordinate Coordinate, insideCoordinate Coordinate) {
	this.coordinates = append(this.coordinates, outsideCoordinate)
	this.sides++
}
