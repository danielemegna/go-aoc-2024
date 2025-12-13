package day20

import (
	"math"
	"strings"
)

type RacetrackMap struct {
	values [][]MapValue
	start  RacetrackElement
}

type RacetrackElement struct {
	Coordinate        Coordinate
	DistanceFromStart int
	Next              *RacetrackElement
}

type MapValue = int

const START MapValue = 0
const WALL MapValue = -1
const TRACK MapValue = math.MaxInt

func ParseRacetrack(rawMapString string) RacetrackMap {
	var mapValues, racetrackStartCoordinate = mapValuesParsing(rawMapString)
	var racetrackStartElement = RacetrackElement{Coordinate: racetrackStartCoordinate, DistanceFromStart: 0, Next: nil}

	createRacetrackElementsChain(&racetrackStartElement, nil, mapValues)
	replaceTrackValuesWithPicosecondsFromStart(&racetrackStartElement, mapValues)

	return RacetrackMap{
		values: mapValues,
		start:  racetrackStartElement,
	}
}

func (this RacetrackMap) RacetrackStart() *RacetrackElement {
	return &this.start
}

func (this RacetrackMap) MapSize() int {
	return len(this.values) // assuming always square maps
}

func (this RacetrackMap) RacetrackLength() int {
	var length = 0
	var element = &this.start
	for element.Next != nil {
		element = element.Next
		length++
	}
	return length
}

func mapValuesParsing(rawMapString string) ([][]MapValue, Coordinate) {
	var inputLines = linesFrom(rawMapString)
	var rawMap = make([][]MapValue, len(inputLines))

	var racetrackStartCoordinate Coordinate
	for y := 0; y < len(inputLines); y++ {
		rawMap[y] = make([]MapValue, len(inputLines[y]))
		for x := 0; x < len(inputLines[y]); x++ {
			switch inputLines[y][x] {
			case '#':
				rawMap[y][x] = WALL
			case 'S':
				rawMap[y][x] = START
				racetrackStartCoordinate = Coordinate{x, y}
			default:
				rawMap[y][x] = TRACK
			}
		}
	}
	return rawMap, racetrackStartCoordinate
}

func createRacetrackElementsChain(current *RacetrackElement, previous *RacetrackElement, mapValues [][]MapValue) {
	var nextCoordinate = findNextRacetrackCoordinate(current.Coordinate, previous, mapValues)
	if nextCoordinate == nil {
		return
	}

	var newRacetrackElement = RacetrackElement{
		Coordinate: *nextCoordinate,
		DistanceFromStart: current.DistanceFromStart+1,
		Next: nil,
	}
	current.Next = &newRacetrackElement
	createRacetrackElementsChain(&newRacetrackElement, current, mapValues)
}

func findNextRacetrackCoordinate(current Coordinate, previous *RacetrackElement, mapValues [][]MapValue) *Coordinate {
	for _, close := range current.CloseCoordinates() {
		if mapValues[close.Y][close.X] == WALL {
			continue
		}
		if previous != nil && close == previous.Coordinate {
			continue
		}

		return &close
	}

	return nil
}

func replaceTrackValuesWithPicosecondsFromStart(racetrackStart *RacetrackElement, mapValues [][]MapValue) {
	var picosecondsFromStart = 0
	var currentElement = racetrackStart
	for currentElement.Next != nil {
		currentElement = currentElement.Next
		mapValues[currentElement.Coordinate.Y][currentElement.Coordinate.X] = picosecondsFromStart + 1
		picosecondsFromStart++
	}
}

func linesFrom(s string) []string {
	var trimmed = strings.TrimSpace(s)
	return strings.Split(trimmed, "\n")
}
