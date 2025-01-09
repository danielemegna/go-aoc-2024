package day06

func ParseGuardMap(mapRows []string) GuardMap {
	return GuardMap{
		size: len(mapRows),
		guard: Guard{
			position:  Coordinate{x: 4, y: 6},
			direction: North,
		},
		obstacles: []Coordinate{},
	}
}

type GuardMap struct {
	size      int // assuming always-square maps here
	guard     Guard
	obstacles []Coordinate
}

type Guard struct {
	position  Coordinate
	direction Direction
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Coordinate struct {
	x int
	y int
}
