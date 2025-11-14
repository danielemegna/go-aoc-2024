package day08

import "strings"

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
				antennaGroup = AntennaGroup{frequency: frequency}
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
	var result = []Coordinate{}
	for _, antennaGroup := range this.antennaGroups {
		// TODO exclude out of bound antinodes
		result = append(result, antennaGroup.antinodes...)
	}
	return result
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
