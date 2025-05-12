package day12

type GardenMap map[rune]GardenRegion

type GardenRegion struct {
	area int
	perimeter int
}
