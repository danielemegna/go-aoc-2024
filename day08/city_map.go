package day08

import (
	"maps"
	"slices"
	"strings"
)

type CityMap struct {
	size          int
	antennaGroups map[Frequency]AntennaGroup
}

type Frequency = rune

func ParseCityMap(rawMapContent string) CityMap {
	var lines = linesFrom(rawMapContent)
	var cityMap = CityMap{
		size:          len(lines), // assuming always square map here
		antennaGroups: make(map[Frequency]AntennaGroup),
	}

	for y := 0; y < len(lines); y++ {
		var line = lines[y]
		for x := 0; x < len(line); x++ {
			var frequency = rune(line[x])
			if frequency == '.' {
				continue
			}

			var antennaGroup, groupExists = cityMap.antennaGroups[frequency]
			if !groupExists {
				antennaGroup = AntennaGroup{}
			}

			antennaGroup.AddAntennaAt(Coordinate{X: x, Y: y})
			cityMap.antennaGroups[frequency] = antennaGroup
		}
	}

	return cityMap
}

func (this CityMap) Size() int {
	return this.size
}

func (this CityMap) Antennas() []Coordinate {
	var result = []Coordinate{}
	for _, antennaGroup := range this.antennaGroups {
		result = append(result, antennaGroup.locations...)
	}
	return result
}

func (this CityMap) AntinodesInMap() []Coordinate {
	var result = map[Coordinate]bool{}
	for _, antennaGroup := range this.antennaGroups {
		for _, antinode := range antennaGroup.antinodes {
			if this.IsOutOfBounds(antinode) {
				continue
			}
			result[antinode] = true
		}
	}
	return slices.Collect(maps.Keys(result))
}

func (this CityMap) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X >= this.size || c.Y < 0 || c.Y >= this.size
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
