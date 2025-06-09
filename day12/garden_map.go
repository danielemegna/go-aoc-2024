package day12

type GardenMap []GardenRegion

type GardenRegion struct {
	plant     rune
	area      int
	perimeter GardenRegionPerimeter
}

func NewGardenRegion(plant rune) GardenRegion {
	return GardenRegion{
		plant: plant,
		area:  0,
		perimeter: GardenRegionPerimeter{},
	}
}
