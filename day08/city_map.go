package day08

import (
	"maps"
	"slices"
	"strings"
)

type CityMap struct {
	size          int
	antennaGroups map[Frequency]*AntennaGroup
}

type Frequency = rune

func ParseCityMap(rawMapContent string, withResonantHarmonics bool) CityMap {
	var lines = linesFrom(rawMapContent)
	var cityMap = CityMap{
		size:          len(lines), // assuming always square map here
		antennaGroups: make(map[Frequency]*AntennaGroup),
	}

	for y := 0; y < len(lines); y++ {
		var line = lines[y]
		for x := 0; x < len(line); x++ {
			var rawMapChar = rune(line[x])
			var coordinate = Coordinate{X: x, Y: y}
			cityMap.parseCityMapCoordinate(rawMapChar, coordinate, withResonantHarmonics)
		}
	}

	return cityMap
}

func (this CityMap) parseCityMapCoordinate(rawMapChar rune, coordinate Coordinate, withResonantHarmonics bool) {
	if rawMapChar == '.' {
		return
	}

	var frequency = Frequency(rawMapChar)
	var antennaGroup = this.getAntennaGroupFor(frequency)
	if withResonantHarmonics {
		antennaGroup.AddAntennaAtWithResonantHarmonics(coordinate)
	} else {
		antennaGroup.AddAntennaAt(coordinate)
	}
}

func (this CityMap) getAntennaGroupFor(frequency Frequency) *AntennaGroup {
	var antennaGroup, exists = this.antennaGroups[frequency]
	if exists {
		return antennaGroup
	}

	var newAntennaGroup = AntennaGroup{mapSize: this.size}
	this.antennaGroups[frequency] = &newAntennaGroup
	return &newAntennaGroup
}

func (this CityMap) Antennas() []Coordinate {
	var result = []Coordinate{}
	for _, antennaGroup := range this.antennaGroups {
		result = append(result, antennaGroup.locations...)
	}
	return result
}

func (this CityMap) AntinodesInMap() []Coordinate {
	var result = map[Coordinate]bool{} // we can avoid duplicate handling here if we turn the antinodes field in a set
	for _, antennaGroup := range this.antennaGroups {
		for _, antinode := range antennaGroup.antinodes {
			if this.IsOutOfBounds(antinode) { // we can avoid this check if we pass mapsize as field in groups
				continue
			}
			result[antinode] = true
		}
	}
	return slices.Collect(maps.Keys(result))
}

// delete me when out of bound check will be in AntennaGroup add functions
func (this CityMap) IsOutOfBounds(c Coordinate) bool {
	return c.IsOutOfBounds(this.size)
}

// maybe not needed after the refactoring
func (this CityMap) Size() int {
	return this.size
}

func linesFrom(s string) []string {
	var lines = strings.Split(s, "\n")
	return lines[:len(lines)-1]
}
