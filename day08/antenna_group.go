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
		this.antinodes[toAdd] = true
		this.antinodes[alreadyPresent] = true

		var xDifference = toAdd.X - alreadyPresent.X
		var yDifference = toAdd.Y - alreadyPresent.Y

		var isLeftAntinodeInBounds = true
		var isRightAntinodeInBounds = true
		for i := 1; isLeftAntinodeInBounds || isRightAntinodeInBounds; i++ {
			if isLeftAntinodeInBounds {
				var nextLeftAntinode = Coordinate{
					alreadyPresent.X - (xDifference * i),
					alreadyPresent.Y - (yDifference * i),
				}
				if nextLeftAntinode.IsOutOfBounds(this.mapSize) {
					isLeftAntinodeInBounds = false
				} else {
					this.antinodes[nextLeftAntinode] = true
				}
			}

			if isRightAntinodeInBounds {
				var nextRightAntinode = Coordinate{
					toAdd.X + (xDifference * i),
					toAdd.Y + (yDifference * i),
				}
				if(nextRightAntinode.IsOutOfBounds(this.mapSize)) {
					isRightAntinodeInBounds = false
				} else {
					this.antinodes[nextRightAntinode] = true
				}
			}
		}
	}

	this.locations = append(this.locations, toAdd)
}

func (this AntennaGroup) GetAntennas() []Coordinate {
	return this.locations
}

func (this AntennaGroup) GetAntinodes() []Coordinate {
	return slices.Collect(maps.Keys(this.antinodes))
}
