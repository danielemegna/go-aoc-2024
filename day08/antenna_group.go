package day08

import (
	"maps"
	"slices"
)

type AntennaGroup struct {
	locations []Coordinate
	antinodes map[Coordinate]bool
	mapSize   int
}

func NewAntennaGroup(mapSize int) AntennaGroup {
	return AntennaGroup{
		locations: []Coordinate{},
		antinodes: map[Coordinate]bool{},
		mapSize:   mapSize,
	}
}

func (this *AntennaGroup) AddAntennaAt(toAdd Coordinate) {
	for _, alreadyPresent := range this.locations {
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y
		var firstAntinode = Coordinate{
			alreadyPresent.X - xDifference,
			alreadyPresent.Y - yDifference,
		}
		var secondAntinode = Coordinate{
			toAdd.X + xDifference,
			toAdd.Y + yDifference,
		}
		if !firstAntinode.IsOutOfBounds(this.mapSize) {
			this.antinodes[firstAntinode] = true
		}
		if !secondAntinode.IsOutOfBounds(this.mapSize) {
			this.antinodes[secondAntinode] = true
		}
	}

	this.locations = append(this.locations, toAdd)
}

func (this *AntennaGroup) AddAntennaAtWithResonantHarmonics(toAdd Coordinate) {
	for _, alreadyPresent := range this.locations {
		this.antinodes[alreadyPresent] = true
		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y

		var continueLoop = true
		for i := 1; continueLoop; i++ {
			var firstAntinode = Coordinate{
				alreadyPresent.X - (xDifference * i),
				alreadyPresent.Y - (yDifference * i),
			}
			var secondAntinode = Coordinate{
				toAdd.X + (xDifference * i),
				toAdd.Y + (yDifference * i),
			}

			continueLoop = false
			if !firstAntinode.IsOutOfBounds(this.mapSize) {
				this.antinodes[firstAntinode] = true
				continueLoop = true
			}
			if !secondAntinode.IsOutOfBounds(this.mapSize) {
				this.antinodes[secondAntinode] = true
				continueLoop = true
			}
		}
	}

	this.locations = append(this.locations, toAdd)
	if len(this.locations) > 1 {
		this.antinodes[toAdd] = true
	}
}

func (this AntennaGroup) GetAntennas() []Coordinate {
	return this.locations
}

func (this AntennaGroup) GetAntinodes() []Coordinate {
	return slices.Collect(maps.Keys(this.antinodes))
}
