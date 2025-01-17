package day06

import "fmt"

type Coordinate struct {
	x int
	y int
}

func (this Coordinate) NextFor(direction Direction) Coordinate {
	switch direction {
	case North:
		return this.North()
	case South:
		return this.South()
	case East:
		return this.East()
	case West:
		return this.West()
	default:
		panic(fmt.Sprintf("Unexpected Direction: %#v", direction))
	}
}

func (this Coordinate) North() Coordinate { return Coordinate{this.x, this.y - 1} }
func (this Coordinate) South() Coordinate { return Coordinate{this.x, this.y + 1} }
func (this Coordinate) East() Coordinate  { return Coordinate{this.x + 1, this.y} }
func (this Coordinate) West() Coordinate  { return Coordinate{this.x - 1, this.y} }
